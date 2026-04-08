package httpx

import (
	"net/http"
	"path/filepath"

	"github.com/szey/Mirai-Nihongo/backend/internal/content"
)

type Handler struct {
	loader *content.Loader
}

func NewHandler() *Handler {
	return &Handler{
		loader: content.NewLoader(filepath.Join("..", "content", "pre-n5", "v1")),
	}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetManifest(w http.ResponseWriter, _ *http.Request) {
	var data map[string]any
	if err := h.loader.Read("manifest.json", &data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *Handler) GetLessons(w http.ResponseWriter, _ *http.Request) {
	var data []map[string]any
	if err := h.loader.Read("lessons.json", &data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *Handler) GetKana(w http.ResponseWriter, _ *http.Request) {
	var data []map[string]any
	if err := h.loader.Read("kana.json", &data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *Handler) GetVocab(w http.ResponseWriter, _ *http.Request) {
	var data []map[string]any
	if err := h.loader.Read("vocab.json", &data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *Handler) GetQuizzes(w http.ResponseWriter, _ *http.Request) {
	var data []map[string]any
	if err := h.loader.Read("quizzes.json", &data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}
