package tu

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/1lann/go-sip/sipnet"
	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/trans"
	"github.com/peakxie/gb28181/pkg/sip/utils"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
)

const DigestLeader = "Digest "

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

func md5Hex(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

func (c *Nistcb) recvRegister(t *trans.Nist, e *trans.Event) error {
	rsp401 := true

	authArgs := make(sipnet.HeaderArgs)
	defer func() {
		if rsp401 {
			//TODO
			Hostname := ""
			authArgs.Set("realm", Hostname)
			authArgs.Set("nonce", randstr.Hex(16))
			authArgs.Set("algorithm", "MD5")
			wwwAuthenticate := "Digest " + authArgs.CommaString()

			rsp := NewResponse(e.Msg, sip.StatusUnauthorized)
			rsp.WWWAuthenticate = wwwAuthenticate
			//err := transport.SendSip((utils.GetDeviceId(t.origRequest.From.Uri.User)), rsp)
			//if err != nil {
			//	logrus.Errorf("SendSip error:%v", err)
			//		return
			//	}
			return
		}

	}()

	//首次请求
	if e.Msg.Authorization == "" {
		return nil
	}

	//验证认证信息
	dlen := len(DigestLeader)
	if len(e.Msg.Authorization) <= dlen || e.Msg.Authorization[:dlen] != DigestLeader {
		logrus.Errorf("Authorization: %s", e.Msg.Authorization)
		return errors.New("Authorization header error")
	}
	//err := authArgs.ParsePairs(e.Msg.Authorization[dlen:])
	//if err != nil {
	//	logrus.Errorf("Parse Authorization err:%v", err)
	//	return err
	//}
	//TODO
	passwd := ""
	// https://tools.ietf.org/html/rfc2069#page-3
	ha1 := md5Hex(utils.GetDeviceId(e.Msg.From.Uri.User) + ":" + authArgs.Get("realm") + ":" + passwd)
	ha2 := md5Hex(e.Msg.Method + ":" + authArgs.Get("uri"))
	response := md5Hex(ha1 + ":" + authArgs.Get("nonce") + ":" + ha2)

	if response != authArgs.Get("response") {
		logrus.Errorf("auth fail, %s != req: %s", response, authArgs.Get("response"))
		return errors.New("auth fail")
	}

	rsp401 = false
	rsp := NewResponse(e.Msg, sip.StatusOK)
	rsp.Expires = e.Msg.Expires
	//err := transport.SendSip((utils.GetDeviceId(t.origRequest.From.Uri.User)), rsp)
	//if err != nil {
	//	logrus.Errorf("SendSip error:%v", err)
	//	return
	//}

	//注销
	if e.Msg.Expires == 0 {
		//TODO
	}

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
