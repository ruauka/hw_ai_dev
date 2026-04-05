package usecase

import "github.com/ruauka/hw1/internal/domain"

// Usecase — интерфейс бизнес-логики анкеты.
type Usecase interface {
	Questions() []domain.Question
	Submit(answers []domain.Answer) error
}
