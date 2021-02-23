package base

import (
	"sync"
)

type Params struct {
	M sync.Map
}

/*
// Get returns an entry in O(n) time.
func (p *Params) Get(name string) *Params {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI parameters in insertion order.
func (p *Params) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}
*/
