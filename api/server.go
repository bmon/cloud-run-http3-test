package api

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type API struct{}

func NewServer() *API {
	return &API{}
}

func logAccess(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		h(rw, r)
		log.Printf("served %s", r.URL)
	}
}

func (a *API) Handler() http.Handler {
	// Setup request router.
	r := mux.NewRouter()
	r.HandleFunc("/", logAccess(a.HandleIndex)).Methods("GET")
	r.HandleFunc("/client-ping", logAccess(a.HandleClientPing)).Methods("POST")

	return r
}

//go:embed index.html
var indexHTML []byte

func (a *API) HandleIndex(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write(indexHTML)
}

func (a *API) HandleClientPing(rw http.ResponseWriter, r *http.Request) {
	serverTime := time.Now().UnixMilli()
	res := struct {
		ServerTime int64  `json:"server_time"`
		ClientIP   string `json:"client_ip"`
	}{serverTime, clientIP(r)}

	data, err := json.Marshal(res)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

func clientIP(req *http.Request) string {
	addr := req.Header.Get("X-Real-Ip")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
	}
	if addr == "" {
		addr = req.RemoteAddr
	}
	return addr
}
