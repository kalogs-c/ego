package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carloshcamilo/ego/entities"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := entities.User{}

	json.NewDecoder(r.Body).Decode(&u)

	user, err := u.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	return
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {

}
