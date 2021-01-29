package udp

import (
	"context"
	"fmt"
	"net"
	"strconv"
)

const (
	ContentLength = "Content-Length:"
)

type Server struct {
	listen   *net.UDPConn
	stopCtx  context.Context
	stopFunc context.CancelFunc
	handler  func(*net.UDPAddr, []byte) bool //数据处理函数
}

func defaultHandler(addr *net.UDPAddr, b []byte) bool { return true }

func NewServer() *Server {
	s := &Server{
		handler: defaultHandler,
	}
	s.stopCtx, s.stopFunc = context.WithCancel(context.Background())
	return s
}

func (s *Server) SetHandler(handler func(*net.UDPAddr, []byte) bool) *Server {
	if handler != nil {
		s.handler = handler
	}
	return s
}

func (s *Server) ListenAndServe(host string, port int) error {
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return err
	}
	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	defer func() {
		s.stopFunc()
		_ = l.Close()
	}()

	s.listen = l

	return s.ReadFromLoop()
}

func (s *Server) ReadFromLoop() error {

	buf := make([]byte, 4096)

	for {

		select {
		case <-s.stopCtx.Done():
			return s.stopCtx.Err()
		default:
		}

		n, remoteAddr, err := s.listen.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("read eof: %v\n", err)
			continue
		}

		_ = s.handler(remoteAddr, buf[:n])
	}
}

func (s *Server) WriteTo(b []byte, address string) (err error) {
	select {
	case <-s.stopCtx.Done():
		return s.stopCtx.Err()
	default:
	}
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}
	_, err = s.listen.WriteToUDP(b, addr)
	return err
}
