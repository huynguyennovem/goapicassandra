package User

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/huynq0911/GoApiCassandra/Cassandra"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUUID gocql.UUID

	user, errs := FormToUser(r)
	if len(errs) != 0 {
		fmt.Println("error, so can not create a new user")
		return
	}

	fmt.Println("creating a new user")
	var created bool = false
	gocqlUUID = gocql.TimeUUID()

	err := Cassandra.Session.Query("INSERT INTO users(id, firstname, lastname, email, city, age) VALUES (?,?,?,?,?,?)",
		gocqlUUID, user.FirstName, user.LastName, user.Email, user.City, user.Age).Exec()
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		created = true
	}

	if created {
		fmt.Println("user id", gocqlUUID)
		json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUUID})
	} else {
		fmt.Println("error to create a new user", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
