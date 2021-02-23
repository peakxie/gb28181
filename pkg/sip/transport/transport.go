package transport

import (
	"bytes"
	"errors"

	"github.com/jart/gosip/sip"
)

var Send func(deviceID string, b []byte) error = defaultSend

func SetSendFunc(handler func(deviceID string, b []byte) error) {
	Send = handler
}

func defaultSend(deviceID string, b []byte) error {
	return nil
}

func SendSip(deviceID string, msg *sip.Msg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}
	var buf bytes.Buffer
	msg.Append(&buf)
	return Send(deviceID, buf.Bytes())
}
