package User

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/huynq0911/GoApiCassandra/Cassandra"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var userList []User
	m := map[string]interface{}{}

	query := "SELECT * FROM users"
	iter := Cassandra.Session.Query(query).Iter()

	for iter.MapScan(m) {
		userList = append(userList, User{
			ID:        m[Id].(gocql.UUID),
			FirstName: m[FirstName].(string),
			LastName:  m[LastName].(string),
			Email:     m[Email].(string),
			City:      m[City].(string),
			Age:       m[Age].(int),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllUserResponse{Users: userList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	var user User
	var errs []string
	var found = false

	vars := mux.Vars(r)
	id := vars["user_uuid"]

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		m := map[string]interface{}{}
		query := "SELECT * FROM users WHERE id=? LIMIT 1"
		iter := Cassandra.Session.Query(query, uuid).Consistency(gocql.One).Iter()
		for iter.MapScan(m) {
			found = true
			user = User{
				ID:        m[Id].(gocql.UUID),
				FirstName: m[FirstName].(string),
				LastName:  m[LastName].(string),
				Email:     m[Email].(string),
				City:      m[City].(string),
				Age:       m[Age].(int),
			}
		}
		if !found {
			errs = append(errs, "User is not found")
		}
	}

	if found {
		json.NewEncoder(w).Encode(GetUserResponse{User: user})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func Enrich(uuids []gocql.UUID) map[string]string {
	if len(uuids) <= 0 {
		fmt.Println("list UUID is empty")
		return map[string]string{}
	}
	names := map[string]string{}
	m := map[string]interface{}{}
	query := "SELECT id, firstname, lastname FROM users WHERE id IN ?"
	iter := Cassandra.Session.Query(query, uuids).Iter()
	for iter.MapScan(m) {
		fmt.Println("m:", m)
		userid := m[Id].(gocql.UUID)
		names[userid.String()] = fmt.Sprintf("%s %s", m[FirstName], m[LastName])
		m = map[string]interface{}{}
	}
	return names
}
