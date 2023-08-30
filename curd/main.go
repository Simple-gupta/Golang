package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func intializeRouter(){
	r:=mux.NewRouter()
	r.HandleFunc("/users",GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")
	r.HandleFunc("/usercustom",GetUserCustom).Methods("GET")
	r.HandleFunc("/customquery",GetUserByCustomQuery).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",r))
}
func main(){
	InitialMigration()
	intializeRouter()
	
}