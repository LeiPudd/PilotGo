package network

import (
	"fmt"
	"net"

	"openeluer.org/PilotGo/PilotGo/pkg/logger"
)

type SocketServer struct {
	// MessageProcesser *protocol.MessageProcesser
	OnAccept func(net.Conn)
	OnStop   func()
}

func (s *SocketServer) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	logger.Debug("Waiting for agents")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		s.OnAccept(conn)

		// go SendHandle(conn)
	}
}

// func SendHandle(conn net.Conn) {
// 	for {
// 		fmt.Println("请输入指令：")
// 		inputReader := bufio.NewReader(os.Stdin)
// 		input, err := inputReader.ReadString('\n')
// 		if err != nil {
// 			continue
// 		}
// 		data := &protocol.Message{
// 			Type: protocol.RunScript,
// 			Body: []byte(input),
// 		}

// 		// data := msg.Encode()
// 		// sendData := protocol.TlvEncode(data)

// 		err = pnet.Send(conn, protocol.TlvEncode(data.Encode()))
// 		if err != nil {
// 			fmt.Println("send error:", err)
// 		}

// 		time.Sleep(time.Second)
// 	}
// }

func (s *SocketServer) Stop() {

}
