package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (s *Store) routes() http.Handler {
	mux := chi.NewMux()
	return mux
}
