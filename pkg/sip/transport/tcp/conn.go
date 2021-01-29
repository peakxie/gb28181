package tcp

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync/atomic"
	"time"
)

type Conn struct {
	cid        uint64
	conn       net.Conn
	srv        *Server
	closed     int32
	lastActive int64
	stopCtx    context.Context
}

func (c *Conn) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		c.conn.Close()
	}
}

func (c *Conn) IsClosed() bool {
	return atomic.LoadInt32(&c.closed) != 0
}

func (c *Conn) GetCid() uint64 {
	return c.cid
}

func (c *Conn) CompleteRead(spliter func(b []byte) (int, bool), handler func(*Conn, []byte) bool) error {

	buf := make([]byte, 4096)
	hasRead := 0

	for {
		if c.IsClosed() {
			return nil
		}

		select {
		case <-c.stopCtx.Done():
			return c.stopCtx.Err()
		default:
		}

		now := time.Now()

		_ = c.conn.SetReadDeadline(now.Add(100 * time.Millisecond))

		n, err := c.conn.Read(buf[hasRead:])
		if 0 < n {
			hasRead += n
			c.lastActive = time.Now().Unix()
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("read eof: %v\n", err)
			}
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() || netErr.Temporary() {
					continue
				}
			}
			return err
		}

		pkgLen, skip := spliter(buf[:hasRead])
		if skip {
			buf = buf[pkgLen:]
			hasRead -= pkgLen
			continue
		}
		if pkgLen <= hasRead {
			if !handler(c, buf[:pkgLen]) {
				return nil
			}
			buf = buf[pkgLen:]
			hasRead -= pkgLen
		} else if pkgLen == 0 {
		} else if len(buf) < pkgLen {
			newBuf := make([]byte, pkgLen)
			copy(newBuf, buf[:hasRead])
			buf = newBuf
		}
	}
}

func (c *Conn) CompleteWrite(b []byte) (err error) {
	pos := 0
	end := len(b)

	for {
		if c.IsClosed() {
			return fmt.Errorf("net closed")
		}

		select {
		case <-c.stopCtx.Done():
			return c.stopCtx.Err()
		default:
		}

		if end <= pos {
			return nil
		}

		now := time.Now()

		_ = c.conn.SetWriteDeadline(now.Add(100 * time.Millisecond))

		n, err := c.conn.Write(b[pos:end])
		if 0 < n {
			pos += n
			c.lastActive = time.Now().Unix()
		}
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() || netErr.Temporary() {
					continue
				}
			}
			return err
		}
	}
}
