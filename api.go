package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() {
	// Initialize the router with path prefix
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Store holds the repository methods
	store := NewStore(s.db)
	mailer := NewSendGridMailer(Envs.SendGridAPIKey, Envs.SendGridFromEmail)

	// Controller holds the API methods
	controller := NewController(store, mailer)
	controller.RegisterRoutes(subrouter)

	log.Println("Starting API server on ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
