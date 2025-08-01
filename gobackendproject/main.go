package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list..."))
}

func (s *server) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			if _, err := w.Write([]byte("index page")); err != nil {
				return
			}
		case "/users":
			if _, err := w.Write([]byte("users page")); err != nil {
				return
			}
		default:
			if _, err := w.Write([]byte("404 error")); err != nil {
				log.Print("no such page")
				return
			}
		}
	case http.MethodPost:
		if _, err := w.Write([]byte("POST method")); err != nil {
			return
		}
	default:
		if _, err := w.Write([]byte("error")); err != nil {
			return
		}
	}
}

func main() {
	s := &server{
		addr: ":8080",
	}

	mux := http.NewServeMux() // router

	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", s.getUsersHandler)
	mux.HandleFunc("POST /users", s.createUsersHandler)

	if err := srv.ListenAndServe(); err != nil {
		return
	}

}
