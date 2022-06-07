package agentConsumer

import "app/domain"

type Encryptor interface {
	Encrypt(domain.Event) domain.Event
}

type Saver interface {
	Save(domain.Event) error
}

type agentConsumer struct {
	encryptor Encryptor
	saver     Saver
}

func New(encryptor Encryptor, saver Saver) *agentConsumer {
	//TODO
	t := agentConsumer{}
	return &t
}

func (t *agentConsumer) Handle(data []byte) error {

	return nil
}
