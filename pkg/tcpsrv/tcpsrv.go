package tcpsrv

import (
	"net"
	"os"

	"app/pkg/logger"
)

var (
	CONN_TYPE string = "tcp"
	CONN_HOST string
	CONN_PORT string

	CONN_URL string = CONN_HOST + ":" + CONN_PORT
)

type TCPSRV struct {
	logger  logger.Logger
	handler TcpContoller
}

type Config interface {
	GetString(string) string
}

type TcpContoller interface {
	HandleTCP(data []byte) error
}

func New(logger logger.Logger, tcpContoller TcpContoller) *TCPSRV {

	return nil
}

func (t *TCPSRV) StartTCP() {
	// Listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_URL)

	if err != nil {
		t.logger.Errorf("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when this application closes
	defer l.Close()

	var i int
	for {
		// Listen for connections
		conn, err := l.Accept()

		if err != nil {
			t.logger.Errorf("Error accepting connection:", err.Error())
			os.Exit(1)
		}
		i++
		if i > 2 {
			os.Exit(0)
		}
		go t.serveTCP(conn)
	}
}

func (t *TCPSRV) serveTCP(conn net.Conn) {
	// Buffer that holds incoming information
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)

		if err != nil {
			t.logger.Errorf("Error reading:", err.Error())
			break
		}

		err = t.handler.HandleTCP(buf)

	}
}
