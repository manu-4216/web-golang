package domain

import (
	"fmt"
	"net/http"

	"github.com/manu-4216/web-golang/mvc/utils"
)

// mock of a DB
var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Manu", LastName: "Micu", Email: "myemail@me.com"},
	}
)

// GetUser gets the user from the DB. Use a pointer, so that we can return nil for the user, and not a User{}
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
