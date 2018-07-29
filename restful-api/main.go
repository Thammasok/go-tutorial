package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type UserLists []User

func getAllUser(w http.ResponseWriter, r *http.Request) {
	users := UserLists{
		User{Name: "Peter", Gender: "male"},
		User{Name: "Sara", Gender: "female"},
	}
	fmt.Println("Endpoint Hit: getAllUser")

	json.NewEncoder(w).Encode(users)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Hit: getUserById")
	// Show the key receiving from API
	fmt.Fprintf(w, "Key: "+key)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	var1 := vars["var1"]
	var2 := vars["var2"]

	fmt.Println("Var 1: " + var1)
	fmt.Println("Var 2: " + var2)
	fmt.Fprintf(w, "Key: "+key)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", getAllUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequestsV2() {
	Router := mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/", homePage)
	Router.HandleFunc("/user", getAllUser)
	Router.HandleFunc("/user/{id}", getUserById)
	Router.HandleFunc("/article/{key}/{var1}/{var2}/", getUser)
	log.Fatal(http.ListenAndServe(":8081", Router))
}

func main() {
	// Rest API v1.0 - Basic API
	// handleRequests()

	// Rest API v2.0 - Mux Routers
	handleRequestsV2()
}
