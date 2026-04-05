// Команда server запускает HTTP-сервер анкеты.
package main

import (
	"log"
	"net/http"

	"github.com/ruauka/hw1/internal/handler"
	"github.com/ruauka/hw1/internal/repository"
	"github.com/ruauka/hw1/internal/usecase"
	"github.com/ruauka/hw1/web"
)

func main() {
	// Ручной DI: repository → usecase → handler.
	store := repository.New()
	survey := usecase.New(store)
	h := handler.New(survey)

	mux := http.NewServeMux()
	h.Register(mux)
	mux.Handle("/", http.FileServerFS(web.FS))

	log.Println("сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
