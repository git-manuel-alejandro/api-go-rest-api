package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
)

func MakeEndPoints() Endpoints {

	return Endpoints{
		Create: makeCreateEndPoint(),
		GetAll: makeGetAllEndPoint(),
		Get:    makeGetEndPoint(),
		Update: makeUpdateEndPoint(),
		Delete: makeDeleteEndPoint(),
	}

}

func makeCreateEndPoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			return
		}

		json.NewEncoder(w).Encode(req)
	}
}

func makeGetAllEndPoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get all user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndPoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndPoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndPoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
