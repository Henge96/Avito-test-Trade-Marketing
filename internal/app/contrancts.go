package app

import (
	"context"
	"depends/internal/api"
)

// Repo ..
type Repo interface {
	SaveByEvent(ctx context.Context, event api.Event) error
	GetEventByDates(ctx context.Context, startSearch, endSearch string) (r []api.Request, err error)
	DeleteEventById(ctx context.Context, id int) error
}
