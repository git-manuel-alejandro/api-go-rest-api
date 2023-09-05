package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

	srv := &http.Server{
		Handler: router,
		// Handler:      http.TimeoutHandler(router, 1*time.Second, "Timeout !!"),
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)

		return

	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(6 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
