package main

import (
	"log"
	"net/http"
	
	"github.com/carloshcamilo/ego/controllers"
	"github.com/gorilla/mux"
)

func main()  {
	r := mux.NewRouter()
	
	uc := controllers.NewUserController()

	r.HandleFunc("/users", uc.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}