package events

import (
	"context"

	"app/domain"
)

// Store implementation is responsible for managing persistance of events.
type Store interface {
	Get(ctx context.Context, name string) (*domain.Event, error)
	Save(ctx context.Context, event domain.Event) (*domain.Event, error)
}
