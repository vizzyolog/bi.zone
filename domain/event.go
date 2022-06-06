package domain

import (
	"fmt"
	"time"
)

type Event struct {
	EventID    int
	Created    time.Time
	SystemName string
	Message    string
	Encrypted  bool
}

// Validate performs basic validation of user information.
func (event Event) Validate() error {
	if event.Message == "" {
		return fmt.Errorf("Empty message")
	}

	return nil
}
