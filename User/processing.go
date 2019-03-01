package User

import (
	"net/http"
	"strconv"
)

func FormToUser(r *http.Request) (User, []string) {
	var user User
	var errStr, ageStr string
	var errs []string
	var err error

	user.FirstName, errStr = processFormField(r, FirstName)
	errs = appendError(errs, errStr)
	user.LastName, errStr = processFormField(r, LastName)
	errs = appendError(errs, errStr)
	user.Email, errStr = processFormField(r, Email)
	errs = appendError(errs, errStr)
	user.City, errStr = processFormField(r, City)
	errs = appendError(errs, errStr)

	ageStr, errStr = processFormField(r, Age)
	if len(ageStr) == 0 {
		errs = append(errs, errStr)
	} else {
		user.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			errs = append(errs, "Age is not integer")
		}
	}
	return user, errs
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing" + field
	}
	return fieldData, ""
}
