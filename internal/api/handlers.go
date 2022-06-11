package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware func
func Middleware() {}

// Save event
func (a Api) Save(w http.ResponseWriter, r *http.Request) {

	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println(err)
		return
	}

	time.Date()

	err = a.app.SaveEvent(r.Context(), event)
	if err != nil {
		log.Println(err)
		return
	}

}

// Get events
func (a Api) Get(w http.ResponseWriter, r *http.Request) {

	var date Date

	err := json.NewDecoder(r.Body).Decode(&date)
	if err != nil {
		log.Println(err)
		return
	}

	if date.StartPeriod == "" && date.EndPeriod == "" {
		fmt.Println("bad request")
		return
	}

	result, err := a.app.GetEvent(r.Context(), date.StartPeriod, date.EndPeriod)
	if err != nil {
		fmt.Println("err app/repo layers")
		return
	}

	err = json.NewEncoder(w).Encode(result)

}

func (a Api) Delete(_ http.ResponseWriter, r *http.Request) {

	var id int

	err := json.NewDecoder(r.Body).Decode(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = a.app.DeleteEvent(r.Context(), id)
	if err != nil {
		fmt.Println("problem app/repo")
		return
	}

}

type (e Event) Validation()