package trans

import (
	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/transport"
	"github.com/peakxie/gb28181/pkg/sip/utils"
)

/*
                         |Request received
                         |pass to TU
                         V
                   +-----------+
                   |           |
                   | Trying    |-------------+
                   |           |             |
                   +-----------+             |200-699 from TU
                         |                   |send response
                         |1xx from TU        |
                         |send response      |
                         |                   |
      Request            V      1xx from TU  |
      send response+-----------+send response|
          +--------|           |--------+    |
          |        | Proceeding|        |    |
          +------->|           |<-------+    |
   +<--------------|           |             |
   |Trnsprt Err    +-----------+             |
   |Inform TU            |                   |
   |                     |                   |
   |                     |200-699 from TU    |
   |                     |send response      |
   |  Request            V                   |
   |  send response+-----------+             |
   |      +--------|           |             |
   |      |        | Completed |<------------+
   |      +------->|           |
   +<--------------|           |
   |Trnsprt Err    +-----------+
   |Inform TU            |
   |                     |Timer J fires
   |                     |-
   |                     |
   |                     V
   |               +-----------+
   |               |           |
   +-------------->| Terminated|
                   |           |
                   +-----------+

       Figure 8: non-INVITE server trans

*/

type Nistcber interface {
	RecvRequest(*Nist, *Event) error
}

type Nist struct {
	id           string
	state        State
	cb           Nistcber
	origRequest  *sip.Msg
	lastResponse *sip.Msg
}

func NewNist(id string, cb Nistcber) *Nist {
	t := &Nist{
		id:    id,
		state: NIST_PRE_TRYING,
		cb:    cb,
	}
	return t
}

func (t *Nist) IsTerminated() bool {
	return t.state == NIST_TERMINATED
}

/*
   NIST_PRE_TRYING: {
       RCV_REQUEST: recvRequest,
   },
   NIST_TRYING: {
       SND_STATUS_1XX:    send1xx,
       SND_STATUS_2XX:    send23456xx,
       SND_STATUS_3456XX: send23456xx,
   },
   NIST_PROCEEDING: {
       RCV_REQUEST:       recvRequest,
       SND_STATUS_1XX:    send1xx,
       SND_STATUS_2XX:    send23456xx,
       SND_STATUS_3456XX: send23456xx,
   },
   NIST_COMPLETED: {
       RCV_REQUEST: recvRequest,
       TIMEOUT_J:   timeoutJ,
   },
*/
func (t *Nist) Do(e *Event) error {
	switch t.state {
	case NIST_PRE_TRYING:
		switch e.Type {
		case RCV_REQUEST:
			t.recvRequest(e)
		}
	case NIST_TRYING:
		switch e.Type {
		case SND_STATUS_1XX:
			t.send1xx(e)
		case SND_STATUS_2XX:
			t.send23456xx(e)
		case SND_STATUS_3456XX:
			t.send23456xx(e)
		}
	case NIST_PROCEEDING:
		switch e.Type {
		case RCV_REQUEST:
			t.recvRequest(e)
		case SND_STATUS_1XX:
			t.send1xx(e)
		case SND_STATUS_2XX:
			t.send23456xx(e)
		case SND_STATUS_3456XX:
			t.send23456xx(e)
		}
	case NIST_COMPLETED:
		switch e.Type {
		case RCV_REQUEST:
			t.recvRequest(e)
		case TIMEOUT_J:
			t.timeoutJ(e)
		}
	}
	return nil
}

func (t *Nist) recvRequest(e *Event) error {
	if t.state == NIST_PRE_TRYING {
		t.origRequest = e.Msg
		t.state = NIST_TRYING
		t.cb.RecvRequest(t, e)
	} else {
		if t.lastResponse != nil {
			err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
			if err != nil {
				t.state = NIST_TERMINATED
				return err
			}
		}
	}
	return nil
}

func (t *Nist) send1xx(e *Event) error {
	t.lastResponse = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = NIST_TERMINATED
		return err
	}

	t.state = NIST_PROCEEDING
	return nil
}

func (t *Nist) send23456xx(e *Event) error {
	t.lastResponse = e.Msg
	err := transport.SendSip(utils.GetDeviceId(t.origRequest.From.Uri.User), t.lastResponse)
	if err != nil {
		t.state = NIST_TERMINATED
		return err
	}

	if t.origRequest.Via.Transport == "UDP" {
		/*			time.AfterFunc(T1*64, func() {
					t.event <- &*Event{
						evt: TIMEOUT_J,
						tid: t.id,
					}
				})*/
	}

	t.state = NIST_COMPLETED
	return nil
}
func (t *Nist) timeoutJ(e *Event) error {
	t.state = NIST_TERMINATED
	return nil
}
