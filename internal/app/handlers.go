package app

import (
	"context"
	"depends/internal/api"
)

func (a *App) SaveEvent(ctx context.Context, event api.Event) error {

	return a.repo.SaveByEvent(ctx, event)

}

func (a *App) GetEvent(ctx context.Context, startSearch, endSearch string) (r []api.Request, err error) {

	return a.repo.GetEventByDates(ctx, startSearch, endSearch)

}

func (a *App) DeleteEvent(ctx context.Context, id int) error {

	return a.repo.DeleteEventById(ctx, id)

}
