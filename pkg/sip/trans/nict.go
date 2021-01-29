package trans

import (
	"time"

	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/transport"
	"github.com/peakxie/gb28181/pkg/sip/utils"
)

/*
非invite事务的状态机：

                                   |Request from TU
                                   |send request
               Timer E             V
               send request  +-----------+
                   +---------|           |-------------------+
                   |         |  Trying   |  Timer F          |
                   +-------->|           |  or Transport Err.|
                             +-----------+  inform TU        |
                200-699         |  |                         |
                resp. to TU     |  |1xx                      |
                +---------------+  |resp. to TU              |
                |                  |                         |
                |   Timer E        V       Timer F           |
                |   send req +-----------+ or Transport Err. |
                |  +---------|           | inform TU         |
                |  |         |Proceeding |------------------>|
                |  +-------->|           |-----+             |
                |            +-----------+     |1xx          |
                |              |      ^        |resp to TU   |
                | 200-699      |      +--------+             |
                | resp. to TU  |                             |
                |              |                             |
                |              V                             |
                |            +-----------+                   |
                |            |           |                   |
                |            | Completed |                   |
                |            |           |                   |
                |            +-----------+                   |
                |              ^   |                         |
                |              |   | Timer K                 |
                +--------------+   | -                       |
                                   |                         |
                                   V                         |
             NOTE:           +-----------+                   |
                             |           |                   |
         transitions         | Terminated|<------------------+
         labeled with        |           |
         the event           +-----------+
         over the action
         to take

                 Figure 6: non-INVITE client trans
*/

type Nictcber interface {
	Recv1xx(*Nict, *Event) error
	Recv23456xx(*Nict, *Event) error
}

type Nict struct {
	id           string
	state        State
	cb           Nictcber
	origRequest  *sip.Msg
	lastResponse *sip.Msg
	timerE       *time.Timer
}

func NewNict(id string, cb Nictcber) *Nict {
	t := &Nict{
		id:    id,
		state: NICT_PRE_TRYING,
		cb:    cb,
	}
	return t
}

func (t *Nict) IsTerminated() bool {
	return t.state == NICT_TERMINATED
}

/*
	NICT_PRE_TRYING: {
		SND_REQUEST: sendRequest,
	},
	NICT_TRYING: {
		TIMEOUT_F:         timeoutF,
		TIMEOUT_E:         timeoutE,
		RCV_STATUS_1XX:    recv1xx,
		RCV_STATUS_2XX:    recv23456xx,
		RCV_STATUS_3456XX: recv23456xx,
	},
	NICT_PROCEEDING: {
		TIMEOUT_F:         timeoutF,
		TIMEOUT_E:         timeoutE,
		RCV_STATUS_1XX:    recv1xx,
		RCV_STATUS_2XX:    recv23456xx,
		RCV_STATUS_3456XX: recv23456xx,
	},
	NICT_COMPLETED: {
		TIMEOUT_K: timeoutK,
	},

*/
func (t *Nict) Do(e *Event) error {
	switch t.state {
	case NICT_PRE_TRYING:
		switch e.Type {
		case SND_REQUEST:
			t.sendRequest(e)
		}
	case NICT_TRYING, NICT_PROCEEDING:
		switch e.Type {
		case TIMEOUT_F:
			t.timeoutF(e)
		case TIMEOUT_E:
			t.timeoutE(e)
		case RCV_STATUS_1XX:
			t.recv1xx(e)
		case RCV_STATUS_2XX:
			t.recv23456xx(e)
		case RCV_STATUS_3456XX:
			t.recv23456xx(e)
		}
	case NICT_COMPLETED:
		switch e.Type {
		case TIMEOUT_K:
			t.timeoutK(e)
		}
	}
	return nil
}

func (t *Nict) sendRequest(e *Event) error {
	t.origRequest = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.origRequest)
	if err != nil {
		t.state = NICT_TERMINATED
		return err
	}

	//启动定时器E
	if t.origRequest.Via.Transport == "UDP" {
		/*			time.AfterFunc(T1*64, func() {
					t.event <- &*Event{
						evt: TIMEOUT_K,
						tid: t.id,
					}
				})*/
	}

	//启动定时器F?在何时被启动

	t.state = NICT_TRYING
	return nil
}

func (t *Nict) recv1xx(e *Event) error {
	t.lastResponse = e.Msg
	t.cb.Recv1xx(t, e)

	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = NICT_TERMINATED
		return err
	}

	//重置定时器E
	t.state = NICT_PROCEEDING
	return nil
}

func (t *Nict) recv23456xx(e *Event) error {
	t.lastResponse = e.Msg
	t.cb.Recv23456xx(t, e)

	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = NICT_TERMINATED
		return err
	}
	//启动定时器K
	if t.origRequest.Via.Transport == "UDP" {
		/*			time.AfterFunc(T1*64, func() {
					t.event <- &*Event{
						evt: TIMEOUT_K,
						tid: t.id,
					}
				})*/
	}

	t.state = NICT_COMPLETED
	return nil
}
func (t *Nict) timeoutK(e *Event) error {
	t.state = NICT_TERMINATED
	return nil
}

func (t *Nict) timeoutF(e *Event) error {
	t.state = NICT_TERMINATED
	return nil
}

func (t *Nict) timeoutE(e *Event) error {
	if t.state == NICT_TRYING {
		//reset timer
		t.timerE.Reset(T1 * 2)
	} else {
		//in PROCEEDING STATE, TIMER is always T2
		t.timerE.Reset(T2)
	}

	//resend origin request
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.origRequest)
	if err != nil {
		t.state = NICT_TERMINATED
		return err
	}
	return nil
}
