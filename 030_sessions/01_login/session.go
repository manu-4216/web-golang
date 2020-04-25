package main

import (
	"net/http"
)

// hasSessionCookie checks if there is already a session cookie in the request
func hasSessionCookie(r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		return false
	}

	return true
}

// getSavedUser gets the saved user from a request
func getSavedUser(r *http.Request) user {
	cookie, err := r.Cookie("session")
	if err != nil {
		return user{}
	}
	uuid := cookie.Value

	username := getUsernameFromSession(uuid)
	u, ok := getUser(username)
	if !ok {
		return user{}
	}
	return u
}

// setSessionCookie sets the session cookie to a new uuid value
func setSessionCookie(w http.ResponseWriter, uuid string) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: uuid,
	}
	http.SetCookie(w, cookie)
}

// Utils for getting the username, and the user
// A map is used the represent a DB of users and of sessions

// getUsernameFromSession gets the username from the uuid
func getUsernameFromSession(uuid string) string {
	username, _ := dbSessions[uuid]
	return username
}

// attachSessionToUsername attaches the session (uuid) to the username
func attachSessionToUsername(uuid, username string) {
	dbSessions[uuid] = username
}

// getUser gets the user from the username
func getUser(username string) (user, bool) {
	u, ok := dbUsers[username]
	return u, ok
}

// setUser sets the user in the map
func setUser(username string, u user) {
	dbUsers[username] = u
}
