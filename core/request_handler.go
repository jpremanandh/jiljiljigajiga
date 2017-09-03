package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bsm/openrtb"
	"github.com/gorilla/mux"
)

func ListenToServer(c chan bool) {
	var server *http.Server
	router := mux.NewRouter()
	server = &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", "127.0.0.1", "9999"),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	var err error
	var bidFunc = func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header()["Content-Type"] = []string{"application/json"}
		var req *openrtb.BidRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": "unable to deserialize config"})
			return
		}
		json.NewEncoder(w).Encode(req)
	}
	router.HandleFunc("/bid", func(w http.ResponseWriter, r *http.Request) {
		bidFunc(w, r)
	}).Methods("POST")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
