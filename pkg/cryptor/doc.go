package cryptor

import "app/domain"

type Encryptor interface {
	Encode(domain.Event) domain.Event
	Decode(domain.Event) domain.Event
}
