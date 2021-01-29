package app

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/jart/gosip/sip"
	"github.com/peakxie/gb28181/pkg/sip/trans"
	"github.com/peakxie/gb28181/pkg/sip/transport"
	"github.com/peakxie/gb28181/pkg/sip/transport/tcp"
	"github.com/peakxie/gb28181/pkg/sip/transport/udp"
	"github.com/peakxie/gb28181/pkg/sip/tu"
	"github.com/peakxie/gb28181/pkg/sip/utils"
)

type Server struct {
	registers sync.Map //DeviceId: *UserConn
	tcp       *tcp.Server
	udp       *udp.Server
	tu        *tu.TU
}

type UserConn struct {
	tp      bool   //tcp，否则udp
	cid     uint64 //tcp时用
	address string //udp时用
}

func NewServer() *Server {
	s := &Server{
		tcp: tcp.NewServer(),
		udp: udp.NewServer(),
		tu:  tu.NewTU(),
	}

	s.tcp.SetHandler(s.tcpHandler)
	s.udp.SetHandler(s.udpHandler)
	transport.SetSendFunc(s.sendHandler)
	return s
}

func (s *Server) Run(host string, port int) {
	go func() {
		fmt.Errorf("tcp server:%v", s.tcp.ListenAndServe(host, port))
	}()

	go func() {
		fmt.Errorf("udp server:%v", s.udp.ListenAndServe(host, port))
	}()

	go s.tu.Run()

	select {}
}

func (s *Server) tcpHandler(conn *tcp.Conn, b []byte) bool {
	msg, err := sip.ParseMsg(b)
	if err != nil {
		return true
	}
	ev := &trans.Event{
		Msg:  msg,
		Tid:  utils.GetTid(msg.Method, msg.Via.Param.Get("branch").Value),
		Type: trans.SetIncomingEventType(msg),
	}

	if !msg.IsResponse() && msg.Method == sip.MethodRegister {
		deviceId := utils.GetDeviceId(msg.From.Uri.User)
		if deviceId == "" {
			return true
		}
		userconn := &UserConn{
			tp:  true,
			cid: conn.GetCid(),
		}

		s.registers.Store(deviceId, userconn)
	}

	s.tu.SetEvent(ev)
	return true
}

func (s *Server) udpHandler(addr *net.UDPAddr, b []byte) bool {
	msg, err := sip.ParseMsg(b)
	if err != nil {
		return true
	}

	ev := &trans.Event{
		Msg:  msg,
		Tid:  utils.GetTid(msg.Method, msg.Via.Param.Get("branch").Value),
		Type: trans.SetIncomingEventType(msg),
	}

	if !msg.IsResponse() && msg.Method == sip.MethodRegister {
		deviceId := utils.GetDeviceId(msg.From.Uri.User)
		if deviceId == "" {
			return true
		}
		userconn := &UserConn{
			tp:      false,
			address: addr.String(),
		}

		s.registers.Store(deviceId, userconn)
	}

	s.tu.SetEvent(ev)
	return true
}

func (s *Server) sendHandler(deviceID string, b []byte) error {
	v, ok := s.registers.Load(deviceID)
	if !ok {
		return errors.New("deviceID not exist")
	}

	userconn, ok := v.(*UserConn)
	if !ok {
		return errors.New("userconn error")
	}
	if userconn.tp {
		conn := s.tcp.GetConn(userconn.cid)
		if conn == nil {
			return errors.New("tcp conn not exist")
		}
		return conn.CompleteWrite(b)
	} else {
		return s.udp.WriteTo(b, userconn.address)
	}
}
