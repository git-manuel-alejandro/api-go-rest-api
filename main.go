package main

import (
	"api/internal/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userEnd := user.MakeEndPoints()

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE")

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
