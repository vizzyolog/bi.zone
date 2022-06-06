package tcpController

import "app/domain"

type Encryptor interface {
	Encrypt(domain.Event) domain.Event
}

type Saver interface {
	Save(domain.Event) error
}

type tcpController struct {
	encryptor Encryptor
	saver     Saver
}

func New(encryptor Encryptor, saver Saver) *tcpController {
	//TODO
	t := tcpController{}
	return &t
}

func (t *tcpController) HandleTCP(data []byte) error {

	return nil
}
