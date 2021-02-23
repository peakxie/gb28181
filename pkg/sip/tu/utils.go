package tu

import (
	"github.com/jart/gosip/sip"
)

func NewResponse(msg *sip.Msg, status int) *sip.Msg {
	return &sip.Msg{
		Status:     status,
		Phrase:     sip.Phrase(status),
		Via:        msg.Via,
		From:       msg.From,
		To:         msg.To,
		CallID:     msg.CallID,
		CSeq:       msg.CSeq,
		CSeqMethod: msg.CSeqMethod,
	}
}
