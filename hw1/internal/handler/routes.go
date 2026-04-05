package handler

import (
	"net/http"

	"github.com/ruauka/hw1/internal/usecase"
)

// Handler — HTTP-обработчик запросов анкеты.
type Handler struct {
	uc usecase.Usecase
}

// New создаёт новый Handler с переданным usecase.
func New(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Register регистрирует маршруты анкеты в переданном ServeMux.
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /questions", h.Questions)
	mux.HandleFunc("POST /answers", h.Submit)
}
