package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
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

	ErrorRes struct {
		Error string `json:"error"`
	}
)

func MakeEndPoints(s Service) Endpoints {

	return Endpoints{
		Create: makeCreateEndPoint(s),
		GetAll: makeGetAllEndPoint(s),
		Get:    makeGetEndPoint(s),
		Update: makeUpdateEndPoint(s),
		Delete: makeDeleteEndPoint(s),
	}

}

func makeCreateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"invalid request format"})
			return
		}
		if req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"first name is required"})
			return
		}
		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"last name is required"})
			return
		}
		err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
		}
		json.NewEncoder(w).Encode(req)
	}
}

func makeGetAllEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get all users")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
