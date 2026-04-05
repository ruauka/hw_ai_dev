// Пакет usecase содержит бизнес-логику анкеты.
package usecase

import (
	"fmt"

	"github.com/ruauka/hw1/internal/domain"
	"github.com/ruauka/hw1/internal/repository"
)

// Survey — usecase для работы с анкетой.
type Survey struct {
	store repository.Store
}

// New создаёт новый Survey с переданным хранилищем.
func New(store repository.Store) *Survey {
	return &Survey{store: store}
}

// Questions возвращает список вопросов анкеты из хранилища.
func (s *Survey) Questions() []domain.Question {
	return s.store.Questions()
}

// Submit сохраняет ответы пользователя через хранилище.
func (s *Survey) Submit(answers []domain.Answer) error {
	if err := s.store.Save(answers); err != nil {
		return fmt.Errorf("submit answers: %w", err)
	}
	return nil
}
