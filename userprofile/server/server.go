package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermebr/devfest/userprofile/user"
)

var users = user.NewUserManager()

func RegisterHandlers() {

	r := mux.NewRouter()
	r.HandleFunc("/users/", errorHandler(ListUsers)).Methods("GET")
	r.HandleFunc("/users/", errorHandler(NewUserHandler)).Methods("POST")
	r.HandleFunc("/users/"+"{id}", errorHandler(GetUser)).Methods("GET")
	r.HandleFunc("/users/"+"{id}", errorHandler(UpdateUser)).Methods("PUT")

	http.Handle("/users/", r)
}

type badRequest struct{ error }
type notFound struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request) error {
	res := users.All()
	return json.NewEncoder(w).Encode(res)
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) error {
	req := struct {
		Nome      string
		Sobrenome string
		Idade     int
	}{}
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return fmt.Errorf("Vish: ", err)
	}
	u, err := user.NewUser(req.Nome, req.Sobrenome, req.Idade)
	if err != nil {
		return err
	}
	return users.Save(u)
}

func parseID(r *http.Request) (int64, error) {
	txt, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("user id not found")
	}
	return strconv.ParseInt(txt, 10, 0)
}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	log.Println("User is ", id)
	if err != nil {
		return err
	}
	u, ok := users.Find(id)
	log.Println("Found", ok)

	if !ok {
		return fmt.Errorf("Not found")
	}
	return json.NewEncoder(w).Encode(u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return err
	}
	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}
	if u.ID != id {
		return fmt.Errorf("inconsistent user IDs")
	}
	if _, ok := users.Find(id); !ok {
		return fmt.Errorf("Not found")
	}
	return users.Save(&u)
}
