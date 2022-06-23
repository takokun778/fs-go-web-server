package main

import (
	"log"
	"net/http"
)

const (
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

func main() {
	http.Handle("/", BasicAuth(http.FileServer(http.Dir("public"))))

	log.Println("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, pass, ok := r.BasicAuth()

		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		if id != ADMIN_USER || pass != ADMIN_PASSWORD {
			w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	})
}
