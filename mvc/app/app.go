package app

import (
	"net/http"

	"github.com/manu-4216/web-golang/mvc/controllers"
)

// StartApp starts the application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
