package tcpSrv

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
	CONN_URL  = CONN_HOST + ":" + CONN_PORT
)

type TCPSRV struct {
	handler TcpContoller
}

type Config interface {
	GetString(string) string
}

type TcpContoller interface {
	HandleTCP(data []byte) error
}

func New(tcpContoller TcpContoller, cfg Config) *TCPSRV {

	return nil
}

func (t *TCPSRV) StartTCP() {
	// Listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_URL)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when this application closes
	defer l.Close()

	fmt.Println("Listening on " + CONN_URL)
	var i int
	for {
		// Listen for connections
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			os.Exit(1)
		}
		i++
		if i > 2 {
			os.Exit(0)
		}
		go t.handleRequest(conn)
	}
}

func (t *TCPSRV) handleRequest(conn net.Conn) {
	// Buffer that holds incoming information
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		err = t.handler.HandleTCP(buf)

	}
}
