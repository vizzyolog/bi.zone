package mongodb

import (
	"context"
	"time"

	"app/domain"

	"gopkg.in/mgo.v2"
)

const colEvents = "events"

// EventStore provides functions for persisting User entities in MongoDB.
type EventStore struct {
	db *mgo.Database
}

// NewEventStore initializes a users store with the given db handle.
func NewEventStore(db *mgo.Database) *EventStore {
	return &EventStore{
		db: db,
	}
}

// Save validates and persists the user.
func (events *EventStore) Save(ctx context.Context, event domain.Event) (*domain.Event, error) {
	event.SetDefaults()
	if err := event.Validate(); err != nil {
		return nil, err
	}
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	col := events.db.C(colEvents)
	if err := col.Insert(event); err != nil {
		return nil, err
	}
	return &event, nil
}
