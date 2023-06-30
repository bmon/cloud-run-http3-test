package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type API struct{}

func NewServer() *API {
	return &API{}
}

func (a *API) Handler() http.Handler {
	// Setup request router.
	r := mux.NewRouter()
	r.HandleFunc("/", a.HandleIndex).Methods("GET")
	r.HandleFunc("/client-ping", a.HandleClientPing).Methods("POST")

	return r
}

func (a *API) HandleIndex(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("<h1>Hello world!</h1>"))
}

func (a *API) HandleClientPing(rw http.ResponseWriter, r *http.Request) {
	serverTime := time.Now().UnixMilli()
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf(`{"server_time": %d}`, serverTime)))
}
