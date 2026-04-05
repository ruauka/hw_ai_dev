package repository

import "github.com/ruauka/hw1/internal/domain"

// Store — интерфейс хранилища анкеты.
type Store interface {
	Questions() []domain.Question
	Save(answers []domain.Answer) error
}
