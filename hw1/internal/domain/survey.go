// Пакет domain содержит доменные типы анкеты без внешних зависимостей.
package domain

// Question — вопрос анкеты с уникальным идентификатором и текстом.
type Question struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Answer — ответ пользователя на конкретный вопрос анкеты.
type Answer struct {
	QuestionID int    `json:"question_id"`
	Text       string `json:"text"`
}
