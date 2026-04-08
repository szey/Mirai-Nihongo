package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", h.Health)
	r.Route("/api/content/pre-n5", func(api chi.Router) {
		api.Get("/manifest", h.GetManifest)
		api.Get("/lessons", h.GetLessons)
		api.Get("/kana", h.GetKana)
		api.Get("/vocab", h.GetVocab)
		api.Get("/quizzes", h.GetQuizzes)
	})

	return r
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
