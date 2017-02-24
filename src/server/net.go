package server

import "net"
type Handle func(conn net.Conn)
type Server struct {
	Port string
	exitCh chan interface{}
	Handle Handle

}

func (s *Server) Start() error {

	listen, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}
	defer func() {
		//close socket
		listen.Close()
	}()
	for {
		select {
		case <-s.exitCh:
			break
		}

		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go s.Handle(conn)
	}
}

func (s *Server) Stop()error{
	s.exitCh<-nil
	return nil
}

func ClientConn(ip string,port string)(conn net.Conn,err error){
	conn,err=net.Dial("tcp",ip+":"+port)
	if err!=nil{
		return
	}
	return
}
