package v1

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(r *http.ServeMux) {
	r.Handle("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
	r.HandleFunc("GET /users/{userID}", GetUsers)
	r.HandleFunc("POST /users/{userID}", AddUser)
	r.HandleFunc("PUT /users/{userID}", UpdateUser)
	r.HandleFunc("DELETE /users/{userID}", DeleteUser)
}
