package repo

import (
	"context"
	"fmt"

	"depends/internal/api"
)

func (c *Core) SaveByEvent(ctx context.Context, event api.Event) error {

	const query = "insert into events (date, views, clicks, cost, cpc, cpm) values $1, $2,$3,$4, $5, $6"
	const thousand float64 = 1000
	_, err := c.db.ExecContext(ctx, query, event.Date, event.Views, event.Clicks, event.Cost, event.Cost/float64(event.Clicks), event.Cost/float64(event.Views)*thousand)
	if err != nil {
		fmt.Println("problem with db")
		return err
	}

	return nil

}

func (c *Core) GetEventByDates(ctx context.Context, startSearch, endSearch string) (r []api.Request, err error) {

	const query = "select * from events where date between $1 and $2"

	result, err := c.db.QueryContext(ctx, query, startSearch, endSearch)
	if err != nil {
		fmt.Println("problem with db")
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		e := api.Request{}
		err := result.Scan(&e.Date, &e.Views, &e.Clicks, &e.Cost, &e.Cpc, &e.Cpm)
		if err != nil {
			fmt.Println("problem with db")
			return nil, err
		}

		r = append(r, e)

	}

	return r, nil

}

// DeleteEventById ,,
func (c *Core) DeleteEventById(ctx context.Context, id int) error {

	const exec = "delete from events where id = $1"
	_, err := c.db.ExecContext(ctx, exec, id)
	if err != nil {
		fmt.Println("problem with db")
		return err
	}

	return nil
}
