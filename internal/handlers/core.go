package handlers

import (
	"database/sql"
	"github.com/borntodie-new/question-go/internal/config"
	"github.com/borntodie-new/question-go/internal/db/sqlc"
	"log"
	"net/http"
)

type Store struct {
	store  sqlc.Store
	config config.Config
	mux    http.Handler
}

func NewCore(db *sql.DB, config config.Config) (*Store, error) {
	store := &Store{
		store:  sqlc.NewStore(db),
		config: config,
	}
	store.mux = store.routes()
	return store, nil
}

func (s *Store) Run() error {
	log.Printf("staring application at %s\n", s.config.ServerAddress)
	return http.ListenAndServe(s.config.ServerAddress, s.mux)
}
