package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/huynq0911/GoApiCassandra/Cassandra"
	"github.com/huynq0911/GoApiCassandra/User"
	"log"
	"net/http"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HeartBeat)
	router.HandleFunc("/users/new", User.Post)
	router.HandleFunc("/users", User.Get)
	router.HandleFunc("/users/{user_uuid}", User.GetOne)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HeartBeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
