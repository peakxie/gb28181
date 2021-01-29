package trans

import (
	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/transport"
	"github.com/peakxie/gb28181/pkg/sip/utils"
)

/*
                      |INVITE
                      |pass INV to TU
   INVITE             V send 100 if TU won't in 200ms
   send response+-----------+
       +--------|           |--------+101-199 from TU
       |        | Proceeding|        |send response
       +------->|           |<-------+
                |           |          Transport Err.
                |           |          Inform TU
                |           |--------------->+
                +-----------+                |
   300-699 from TU |     |2xx from TU        |
   send response   |     |send response      |
                   |     +------------------>+
                   |                         |
   INVITE          V          Timer G fires  |
   send response+-----------+ send response  |
       +--------|           |--------+       |
       |        | Completed |        |       |
       +------->|           |<-------+       |
                +-----------+                |
                   |     |                   |
               ACK |     |                   |
               -   |     +------------------>+
                   |        Timer H fires    |
                   V        or Transport Err.|
                +-----------+  Inform TU     |
                |           |                |
                | Confirmed |                |
                |           |                |
                +-----------+                |
                      |                      |
                      |Timer I fires         |
                      |-                     |
                      |                      |
                      V                      |
                +-----------+                |
                |           |                |
                | Terminated|<---------------+
                |           |
                +-----------+

     Figure 7: INVITE server transaction
*/

type Istcber interface {
	RecvInvite(*Ist, *Event) error
	RecvAck(*Ist, *Event) error
}

type Ist struct {
	id           string
	state        State
	cb           Istcber
	origRequest  *sip.Msg
	lastResponse *sip.Msg
	ack          *sip.Msg
}

func NewIst(id string, cb Istcber) *Ist {
	t := &Ist{
		id:    id,
		state: IST_PRE_PROCEEDING,
		cb:    cb,
	}
	return t
}

func (t *Ist) IsTerminated() bool {
	return t.state == IST_TERMINATED
}

/*
	IST_PRE_PROCEEDING: {
		RCV_REQINVITE: recvInvite,
	},
	IST_PROCEEDING: {
		RCV_REQINVITE:     recvInvite,
		SND_STATUS_1XX:    send1xx,
		SND_STATUS_2XX:    send2xx,
		SND_STATUS_3456XX: send3456xx,
	},
	IST_COMPLETED: {
		RCV_REQINVITE: recvInvite,
		RCV_REQACK:    recvAck,
		TIMEOUT_G:     timeoutG,
		TIMEOUT_H:     timeoutH,
		TIMEOUT_I:     timeoutI,
	},

*/
func (t *Ist) Do(e *Event) error {
	switch t.state {
	case IST_PRE_PROCEEDING:
		switch e.Type {
		case RCV_REQINVITE:
			t.recvInvite(e)
		}
	case IST_PROCEEDING:
		switch e.Type {
		case RCV_REQINVITE:
			t.recvInvite(e)
		case SND_STATUS_1XX:
			t.send1xx(e)
		case SND_STATUS_2XX:
			t.send2xx(e)
		case SND_STATUS_3456XX:
			t.send3456xx(e)
		}
	case IST_COMPLETED:
		switch e.Type {
		case RCV_REQINVITE:
			t.recvInvite(e)
		case RCV_REQACK:
			t.recvAck(e)
		case TIMEOUT_G:
			t.timeoutG(e)
		case TIMEOUT_H:
			t.timeoutH(e)
		case TIMEOUT_I:
			t.timeoutI(e)
		}
	}
	return nil
}

func (t *Ist) recvInvite(e *Event) error {
	if t.state == IST_PRE_PROCEEDING {
		t.origRequest = e.Msg
		switch e.Msg.Method {
		case sip.MethodInvite:
			t.cb.RecvInvite(t, e)
		}
		t.state = IST_PROCEEDING
	} else {
		if t.lastResponse != nil {
			err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
			if err != nil {
				t.state = IST_TERMINATED
				return err
			}
		}
	}
	return nil
}

func (t *Ist) send1xx(e *Event) error {
	t.lastResponse = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = IST_TERMINATED
		return err
	}
	return nil
}

func (t *Ist) send2xx(e *Event) error {
	t.lastResponse = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = IST_TERMINATED
		return err
	}
	t.state = IST_TERMINATED
	return nil
}

func (t *Ist) send3456xx(e *Event) error {
	t.lastResponse = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = IST_TERMINATED
		return err
	}

	//创建TimeoutG

	//创建TimeoutH

	if t.origRequest.Via.Transport == "UDP" {
		/*			time.AfterFunc(T1*64, func() {
					t.event <- &*Event{
						evt: TIMEOUT_J,
						tid: t.id,
					}
				})*/
	}

	t.state = IST_COMPLETED
	return nil
}
func (t *Ist) recvAck(e *Event) error {
	t.ack = e.Msg
	switch e.Msg.Method {
	case sip.MethodAck:
		t.cb.RecvAck(t, e)
	}
	t.state = IST_CONFIRMED
	return nil
}

func (t *Ist) timeoutG(e *Event) error {
	//创建TimeoutG 时间*2  大于T2时保持T2

	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = IST_TERMINATED
		return err
	}
	return nil
}

func (t *Ist) timeoutH(e *Event) error {
	t.state = IST_TERMINATED
	return nil
}

func (t *Ist) timeoutI(e *Event) error {
	t.state = IST_TERMINATED
	return nil
}
