package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	repository "github.com/leonardotomascostadasilva/X9/internal/repositories"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", s.healthHandler)
	r.Get("/v1/messages", s.getMessagesHandler)
	r.Get("/v1/messages/{squad}", s.getMessagesBySquadHandler)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	messages, err := repository.GetMessagesLast30Days(ctx)
	if err != nil {
		http.Error(w, "Erro ao buscar mensagens", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Erro ao serializar resposta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func (s *Server) getMessagesBySquadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	squad := chi.URLParam(r, "squad")

	messages, err := repository.GetMessagesBySquad(ctx, squad)
	if err != nil {
		http.Error(w, "Erro ao buscar mensagens", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Erro ao serializar resposta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}
