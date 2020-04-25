package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{} // uuid, username
var dbUsers = map[string]user{}      // username, user

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	seedWithMockUser()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	var u user
	if hasSessionCookie(r) {
		u = getSavedUser(r)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")

		// Is the username taken ?
		_, ok := getUser(username)
		if ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// Create new user
		uuid := uuid.New().String()
		setSessionCookie(w, uuid)
		attachSessionToUsername(uuid, username)
		setUser(username, user{
			Username: username,
			Password: hashPassword(r.FormValue("password")),
			First:    r.FormValue("firstname"),
			Last:     r.FormValue("lastname"),
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Does the username exist ?
		u, ok := getUser(username)
		if !ok {
			http.Error(w, "Username doesn't exist", http.StatusForbidden)
			return
		}

		match := checkPassword(password, u.Password)
		if !match {
			http.Error(w, "Password is wrong", http.StatusForbidden)
			return
		}

		// Persist the user
		uuid := uuid.New().String()
		setSessionCookie(w, uuid)
		attachSessionToUsername(uuid, username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "signin.gohtml", nil)
}

// seedWithMockUser facilitates login testing by creating an existing user
func seedWithMockUser() {
	dbUsers["test@test.com"] = user{
		Username: "test@test.com",
		Password: hashPassword("password"),
		First:    "manuel",
		Last:     "m",
	}
}

// hashPassword hashes a password with bcrypt
func hashPassword(password string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes
}

// checkPassword checks that the password matches with the real hashed one
func checkPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
