// Пакет handler содержит HTTP-обработчики анкеты.
package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ruauka/hw1/internal/domain"
)

// Questions обрабатывает GET /questions — возвращает список вопросов анкеты.
func (h *Handler) Questions(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.uc.Questions()) //nolint:errcheck
}

// Submit обрабатывает POST /answers — принимает ответы пользователя и сохраняет их.
func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	var answers []domain.Answer
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "пустое тело запроса", http.StatusBadRequest)
			return
		}
		http.Error(w, "невалидный JSON", http.StatusBadRequest)
		return
	}

	if err := h.uc.Submit(answers); err != nil {
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
