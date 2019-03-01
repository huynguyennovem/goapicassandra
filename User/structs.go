package User

import "github.com/gocql/gocql"

const Id string = "id"
const FirstName string = "firstname"
const LastName string = "lastname"
const Email string = "email"
const City string = "city"
const Age string = "age"

type User struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	City      string     `json:"city"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type AllUserResponse struct {
	Users []User `json:"users"`
}

type NewUserResponse struct {
	ID gocql.UUID `json:"id"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
