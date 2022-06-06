package events

import (
	"context"

	"app/domain"
)

// Store implementation is responsible for managing persistance of posts.
type Store interface {
	Get(ctx context.Context, name string) (*domain.Event, error)
	Exists(ctx context.Context, name string) bool
	Save(ctx context.Context, post domain.Event) (*domain.Event, error)
	Delete(ctx context.Context, name string) (*domain.Event, error)
}

// userVerifier is responsible for verifying existence of a user.
type userVerifier interface {
	Exists(ctx context.Context, name string) bool
}
