package api

import (
	"context"
)

type Application interface {
	SaveEvent(ctx context.Context, event Event) error
	GetEvent(ctx context.Context, startSearch, endSearch string) ([]Request, error)
	DeleteEvent(ctx context.Context, id int) error
}
