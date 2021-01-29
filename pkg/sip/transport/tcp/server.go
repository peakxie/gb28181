package tcp

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	HeaderEnd     = "\r\n\r\n"
	LineEnd       = "\r\n"
	ContentLength = "Content-Length:"
)

type Server struct {
	stopCtx  context.Context
	stopFunc context.CancelFunc
	spliter  func([]byte) (int, bool) //tcp分包函数
	handler  func(*Conn, []byte) bool //数据处理函数
	connPool sync.Map                 //cid:net.Conn
	gcid     uint64
}

func defaultSpliter(b []byte) (int, bool) {
	if b == nil {
		return 0, false
	}

	//包不够
	ind := strings.Index(string(b), HeaderEnd)
	if ind == -1 {
		return 0, false
	}
	indEnd := ind + len(HeaderEnd)
	//没有Content-Length:头，不是有效的包，跳过
	indCL := strings.Index(string(b[:ind]), ContentLength)
	if indCL == -1 {
		return indEnd, true
	}
	indCLEnd := indCL + len(ContentLength)

	//未找到Content-Length:头的行结尾，理论上不可能
	indLine := strings.Index(string(b[indCL:indEnd]), LineEnd)
	if indLine == -1 {
		return indEnd, true
	}

	//Content-Length值为空，理论上不可能
	n, err := strconv.ParseInt(strings.Trim(string(b[indCLEnd:indLine]), " "), 10, 64)
	if err != nil {
		return indEnd, true
	}

	return indEnd + int(n), false
}
func defaultHandler(*Conn, []byte) bool { return true }

func NewServer() *Server {
	s := &Server{
		spliter: defaultSpliter,
		handler: defaultHandler,
		gcid:    1,
	}
	s.stopCtx, s.stopFunc = context.WithCancel(context.Background())
	return s
}

func (s *Server) SetSpliter(spliter func([]byte) (int, bool)) *Server {
	if spliter != nil {
		s.spliter = spliter
	}
	return s
}

func (s *Server) SetHandler(handler func(*Conn, []byte) bool) *Server {
	if handler != nil {
		s.handler = handler
	}
	return s
}

func (s *Server) NewConn(conn net.Conn) *Conn {
	cid := atomic.AddUint64(&s.gcid, 1)
	c := &Conn{
		cid:        cid,
		conn:       conn,
		srv:        s,
		closed:     0,
		lastActive: time.Now().Unix(),
		stopCtx:    s.stopCtx,
	}
	s.connPool.Store(c.cid, c)
	return c
}

func (s *Server) GetConn(cid uint64) *Conn {
	if v, ok := s.connPool.Load(cid); ok {
		return v.(*Conn)
	}
	return nil
}

func (s *Server) ListenAndServe(host string, port int) error {
	l, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return err
	}

	defer s.stopFunc()

	var tempDelay time.Duration
	for {
		select {
		case <-s.stopCtx.Done():
			_ = l.Close()
			return nil
		default:
		}

		conn, err := l.Accept()
		if err != nil {
			//照搬http出错处理
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				fmt.Printf("http: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}

		c := s.NewConn(conn)

		go func() {
			_ = c.CompleteRead(s.spliter, s.handler)
			s.connPool.Delete(c.cid)
			c.Close()
		}()
	}

	return nil
}
