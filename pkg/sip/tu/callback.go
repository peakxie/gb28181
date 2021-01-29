package tu

import (
	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/trans"
)

// nist cb
type Nistcb struct {
}

func NewNistcb() *Nistcb {
	return &Nistcb{}
}

func (c *Nistcb) RecvRequest(t *trans.Nist, e *trans.Event) error {
	switch e.Msg.Method {
	case sip.MethodRegister:
		return c.recvRegister(t, e)
	case sip.MethodMessage:
		return c.recvMessage(t, e)
	}
	return nil
}

func (c *Nistcb) recvRegister(t *trans.Nist, e *trans.Event) error {
	//首次请求
	if e.Msg.Authorization == "" {

	}

	//验证认证信息

	return nil
}

func (c *Nistcb) recvMessage(t *trans.Nist, e *trans.Event) error {
	return nil
}

// ist cb
type Istcb struct {
}

func NewIstcb() *Istcb {
	return &Istcb{}
}

func (c *Istcb) RecvInvite(t *trans.Ist, e *trans.Event) error {
	return nil
}
func (c *Istcb) RecvAck(t *trans.Ist, e *trans.Event) error {
	return nil
}

// nict cb
type Nictcb struct {
}

func NewNictcb() *Nictcb {
	return &Nictcb{}
}

func (c *Nictcb) Recv1xx(t *trans.Nict, e *trans.Event) error {
	return nil
}
func (c *Nictcb) Recv23456xx(t *trans.Nict, e *trans.Event) error {
	return nil
}

// ict cb
type Ictcb struct {
}

func NewIctcb() *Ictcb {
	return &Ictcb{}
}
func (c *Ictcb) Recv1xx(t *trans.Ict, e *trans.Event) error {
	return nil
}
func (c *Ictcb) Recv2xx(t *trans.Ict, e *trans.Event) error {
	return nil
}
func (c *Ictcb) Recv3456xx(t *trans.Ict, e *trans.Event) error {
	return nil
}
