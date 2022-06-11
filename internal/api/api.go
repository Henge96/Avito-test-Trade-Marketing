package api

import (
	"github.com/gorilla/mux"
)

// Api for
type Api struct {
	app Application
}

func NewRouter(app Application) *mux.Router {

	api := &Api{
		app: app,
	}

	r := mux.NewRouter()

	r.HandleFunc("/event", api.Save).Methods("POST")
	r.HandleFunc("/event", api.Get).Methods("GET")
	r.HandleFunc("/event", api.Delete).Methods("DELETE")

	return r

}
