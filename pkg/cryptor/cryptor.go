package cryptor

import "app/domain"

type Cryptor struct{}

func New() Cryptor {
	return Cryptor{}
}

func (c *Cryptor) Encode(event domain.Event) (domain.Event, error) {
	event.Encrypted = true
	return event, nil
}

func (c *Cryptor) Decode(event domain.Event) (domain.Event, error) {
	event.Encrypted = false
	return event, nil
}
