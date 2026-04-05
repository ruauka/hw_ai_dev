package usecase

import (
	"errors"
	"testing"

	"github.com/ruauka/hw1/internal/domain"
)

// mockStore — тестовая заглушка интерфейса Store.
type mockStore struct {
	questionsFunc func() []domain.Question
	saveFunc      func(answers []domain.Answer) error
}

func (m *mockStore) Questions() []domain.Question {
	if m.questionsFunc != nil {
		return m.questionsFunc()
	}
	return nil
}

func (m *mockStore) Save(answers []domain.Answer) error {
	if m.saveFunc != nil {
		return m.saveFunc(answers)
	}
	return nil
}

func TestSurvey_Questions(t *testing.T) {
	t.Parallel()

	uc := New(&mockStore{
		questionsFunc: func() []domain.Question {
			return []domain.Question{
				{ID: 1, Text: "Как вас зовут?"},
				{ID: 2, Text: "Сколько лет вы программируете на Go?"},
			}
		},
	})
	got := uc.Questions()

	if len(got) == 0 {
		t.Fatal("Questions() вернул пустой список")
	}

	for i, q := range got {
		if q.ID == 0 {
			t.Errorf("questions[%d].ID == 0", i)
		}
		if q.Text == "" {
			t.Errorf("questions[%d].Text пустой", i)
		}
	}
}

func TestSurvey_Submit(t *testing.T) {
	t.Parallel()

	errStore := errors.New("store error")

	tests := []struct {
		name     string
		storeErr error
		answers  []domain.Answer
		wantErr  error
	}{
		{
			name:    "ok",
			answers: []domain.Answer{{QuestionID: 1, Text: "Go"}},
			wantErr: nil,
		},
		{
			name:     "store error propagated",
			answers:  []domain.Answer{{QuestionID: 1, Text: "Go"}},
			storeErr: errStore,
			wantErr:  errStore,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := New(&mockStore{
				saveFunc: func(_ []domain.Answer) error {
					return tt.storeErr
				},
			})

			err := uc.Submit(tt.answers)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
