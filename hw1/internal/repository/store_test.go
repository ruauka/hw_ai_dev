package repository

import (
	"sync"
	"testing"

	"github.com/ruauka/hw1/internal/domain"
)

func TestStore_Questions(t *testing.T) {
	t.Parallel()

	s := New()
	got := s.Questions()

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

func TestStore_Save(t *testing.T) {
	t.Parallel()

	s := New()
	answers := []domain.Answer{{QuestionID: 1, Text: "Go"}}

	if err := s.Save(answers); err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	s.mu.Lock()
	got := len(s.answers)
	s.mu.Unlock()

	if got != 1 {
		t.Errorf("len(answers) = %d, want 1", got)
	}
}

func TestStore_Save_Multiple(t *testing.T) {
	t.Parallel()

	s := New()
	answers := []domain.Answer{{QuestionID: 1, Text: "Go"}}

	for range 3 {
		if err := s.Save(answers); err != nil {
			t.Fatalf("Save() unexpected error: %v", err)
		}
	}

	s.mu.Lock()
	got := len(s.answers)
	s.mu.Unlock()

	if got != 3 {
		t.Errorf("len(answers) = %d, want 3", got)
	}
}

func TestStore_Save_Concurrent(t *testing.T) {
	t.Parallel()

	s := New()
	answers := []domain.Answer{{QuestionID: 1, Text: "Go"}}

	const goroutines = 10

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			if err := s.Save(answers); err != nil {
				t.Errorf("Save() unexpected error: %v", err)
			}
		}()
	}

	wg.Wait()

	s.mu.Lock()
	got := len(s.answers)
	s.mu.Unlock()

	if got != goroutines {
		t.Errorf("len(answers) = %d, want %d", got, goroutines)
	}
}
