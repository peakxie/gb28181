package trans

import (
	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/transport"
	"github.com/peakxie/gb28181/pkg/sip/utils"
)

/*
                               |INVITE from TU
             Timer A fires     |INVITE sent
             Reset A,          V                      Timer B fires
             INVITE sent +-----------+                or Transport Err.
               +---------|           |---------------+inform TU
               |         |  Calling  |               |
               +-------->|           |-------------->|
                         +-----------+ 2xx           |
                            |  |       2xx to TU     |
                            |  |1xx                  |
    300-699 +---------------+  |1xx to TU            |
   ACK sent |                  |                     |
resp. to TU |  1xx             V                     |
            |  1xx to TU  -----------+               |
            |  +---------|           |               |
            |  |         |Proceeding |-------------->|
            |  +-------->|           | 2xx           |
            |            +-----------+ 2xx to TU     |
            |       300-699    |                     |
            |       ACK sent,  |                     |
            |       resp. to TU|                     |
            |                  |                     |      NOTE:
            |  300-699         V                     |
            |  ACK sent  +-----------+Transport Err. |  transitions
            |  +---------|           |Inform TU      |  labeled with
            |  |         | Completed |-------------->|  the event
            |  +-------->|           |               |  over the action
            |            +-----------+               |  to take
            |              ^   |                     |
            |              |   | Timer D fires       |
            +--------------+   | -                   |
                               |                     |
                               V                     |
                         +-----------+               |
                         |           |               |
                         | Terminated|<--------------+
                         |           |
                         +-----------+

                 Figure 5: INVITE client trans
*/

type Ictcber interface {
	Recv1xx(*Ict, *Event) error
	Recv2xx(*Ict, *Event) error
	Recv3456xx(*Ict, *Event) error
}

type Ict struct {
	id           string
	state        State
	cb           Ictcber
	origRequest  *sip.Msg
	lastResponse *sip.Msg
	ack          *sip.Msg
}

func NewIct(id string, cb Ictcber) *Ict {
	t := &Ict{
		id:    id,
		state: ICT_PRE_CALLING,
		cb:    cb,
	}
	return t
}

func (t *Ict) IsTerminated() bool {
	return t.state == ICT_TERMINATED
}

/*
	ICT_PRE_CALLING: {
		SND_REQINVITE: sendInvite,
	},
	ICT_CALLING: {
		TIMEOUT_A:         timeoutA,
		TIMEOUT_B:         timeoutB,
		RCV_STATUS_1XX:    recv1xx,
		RCV_STATUS_2XX:    recv2xx,
		RCV_STATUS_3456XX: recv3456xx,
	},
	ICT_PROCEEDING: {
		RCV_STATUS_1XX:    recv1xx,
		RCV_STATUS_2XX:    recv2xx,
		RCV_STATUS_3456XX: recv3456xx,
	},
	ICT_COMPLETED: {
		RCV_STATUS_3456XX: recv3456xx,
		TIMEOUT_D:         timeoutD,
	},

*/
func (t *Ict) Do(e *Event) error {
	switch t.state {
	case ICT_PRE_CALLING:
		switch e.Type {
		case SND_REQINVITE:
			t.sendInvite(e)
		}
	case ICT_CALLING:
		switch e.Type {
		case TIMEOUT_A:
			t.timeoutA(e)
		case TIMEOUT_B:
			t.timeoutB(e)
		case RCV_STATUS_1XX:
			t.recv1xx(e)
		case RCV_STATUS_2XX:
			t.recv2xx(e)
		case RCV_STATUS_3456XX:
			t.recv3456xx(e)
		}

	case ICT_PROCEEDING:
		switch e.Type {
		case SND_STATUS_1XX:
			t.recv1xx(e)
		case RCV_STATUS_2XX:
			t.recv2xx(e)
		case RCV_STATUS_3456XX:
			t.recv3456xx(e)
		}
	case ICT_COMPLETED:
		switch e.Type {
		case RCV_STATUS_3456XX:
			t.recv3456xx(e)
		case TIMEOUT_D:
			t.timeoutD(e)
		}
	}
	return nil
}

func (t *Ict) sendInvite(e *Event) error {
	t.origRequest = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), e.Msg)
	if err != nil {
		t.state = ICT_TERMINATED
		return err
	}

	//启动定时器A
	//启动定时器B

	return nil
}

func (t *Ict) recv1xx(e *Event) error {
	t.lastResponse = e.Msg
	t.cb.Recv1xx(t, e)

	t.state = ICT_PROCEEDING
	return nil
}

func createAck(t *Ict, m *sip.Msg) *sip.Msg {
	return m
}

//注意GB28181收到2xx也要求Ack
func (t *Ict) recv2xx(e *Event) error {
	t.lastResponse = e.Msg
	t.cb.Recv2xx(t, e)

	if t.ack == nil {
		t.ack = createAck(t, e.Msg)
	}
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.ack)
	if err != nil {
		t.state = ICT_TERMINATED
		return err
	}
	//创建TimeoutD
	t.state = ICT_COMPLETED
	return nil
}

func (t *Ict) recv3456xx(e *Event) error {
	t.lastResponse = e.Msg
	t.cb.Recv3456xx(t, e)

	if t.ack == nil {
		t.ack = createAck(t, e.Msg)
	}
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.ack)
	if err != nil {
		t.state = ICT_TERMINATED
		return err
	}

	//创建TimeoutD
	t.state = ICT_COMPLETED

	return nil
}
func (t *Ict) timeoutA(e *Event) error {
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.origRequest)
	if err != nil {
		t.state = ICT_TERMINATED
		return err
	}
	//不稳定网络
	//t.timerA.Reset(t.timerA.timeout * 2)

	return nil
}

func (t *Ict) timeoutB(e *Event) error {
	t.state = ICT_TERMINATED
	return nil
}

func (t *Ict) timeoutD(e *Event) error {
	t.state = ICT_TERMINATED
	return nil
}
