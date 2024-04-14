package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// _ "github.com/lib/pq" // PostgreSQL driver
)

type apiConfig struct {
	port string
}

func newApiServer(port string) *apiConfig {
	return &apiConfig{
		port: port,
	}
}

func (cfg *apiConfig) apiRun() {
	router := mux.NewRouter()

	//routing logic goes here-

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	}).Methods("GET")

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(w, r)
    }).Methods("GET")

	//routing End/Up
	corsMux := middlewareCors(router)
	srv := &http.Server{
		Addr:         ":" + cfg.port,
		Handler:      corsMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Println("Server Started listening on", cfg.port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}