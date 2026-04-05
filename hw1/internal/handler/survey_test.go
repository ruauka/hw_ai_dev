package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ruauka/hw1/internal/domain"
)

// mockUsecase — тестовая заглушка интерфейса Usecase.
type mockUsecase struct {
	questions  []domain.Question
	submitFunc func(answers []domain.Answer) error
}

func (m *mockUsecase) Questions() []domain.Question {
	return m.questions
}

func (m *mockUsecase) Submit(answers []domain.Answer) error {
	if m.submitFunc != nil {
		return m.submitFunc(answers)
	}
	return nil
}

func TestHandler_Questions(t *testing.T) {
	t.Parallel()

	uc := &mockUsecase{
		questions: []domain.Question{
			{ID: 1, Text: "Как вас зовут?"},
		},
	}

	h := New(uc)
	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	rec := httptest.NewRecorder()

	h.Questions(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want %q", ct, "application/json")
	}
	if body := rec.Body.String(); !strings.Contains(body, "Как вас зовут?") {
		t.Errorf("body does not contain question text: %s", body)
	}
}

func TestHandler_Submit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		body       string
		submitFunc func([]domain.Answer) error
		wantCode   int
	}{
		{
			name:     "ok",
			body:     `[{"question_id":1,"text":"Go"}]`,
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid json",
			body:     `not json`,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "empty body",
			body:     ``,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "usecase error",
			body: `[{"question_id":1,"text":"Go"}]`,
			submitFunc: func(_ []domain.Answer) error {
				return errors.New("store unavailable")
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := &mockUsecase{submitFunc: tt.submitFunc}
			h := New(uc)

			req := httptest.NewRequest(http.MethodPost, "/answers", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			h.Submit(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantCode)
			}
		})
	}
}
