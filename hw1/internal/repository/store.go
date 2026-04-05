// Пакет repository содержит in-memory реализацию хранилища анкеты.
package repository

import (
	"sync"

	"github.com/ruauka/hw1/internal/domain"
)

// Compile-time проверка соответствия интерфейсу Store.
var _ Store = (*store)(nil)

// hardcodedQuestions — захардкоженный список вопросов анкеты.
var hardcodedQuestions = []domain.Question{
	{ID: 1, Text: "Как вас зовут?"},
	{ID: 2, Text: "Сколько лет вы программируете на Go?"},
	{ID: 3, Text: "Какой ваш любимый инструмент разработки?"},
	{ID: 4, Text: "Какие темы вас интересуют в этом курсе?"},
	{ID: 5, Text: "Чего вы ожидаете от обучения?"},
}

// store — потокобезопасное in-memory хранилище анкеты.
type store struct {
	mu      sync.Mutex
	answers [][]domain.Answer
}

// New создаёт новый экземпляр хранилища.
func New() *store {
	return &store{}
}

// Questions возвращает список вопросов анкеты.
func (s *store) Questions() []domain.Question {
	return hardcodedQuestions
}

// Save сохраняет набор ответов пользователя в памяти.
func (s *store) Save(answers []domain.Answer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.answers = append(s.answers, answers)
	return nil
}
