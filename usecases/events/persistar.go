package events

import (
	"context"

	"app/domain"
	"app/pkg/logger"
)

// NewPersitar initializes the events usecase.
func NewPersistar(lg logger.Logger, store Store) *Persistar {
	return &Persistar{
		Logger: lg,
		store:  store,
	}
}

// Publication implements the events usecases.
type Persistar struct {
	logger.Logger

	store Store
}

// Publish validates and persists the post into the store.
func (pub *Persistar) Publish(ctx context.Context, event domain.Event) (*domain.Event, error) {
	if err := event.Validate(); err != nil {
		return nil, err
	}

	saved, err := pub.store.Save(ctx, event)
	if err != nil {
		pub.Warnf("failed to save post to the store: %+v", err)
	}

	return saved, nil
}
