# Plan: Мини-анкета (Mini Questionnaire)

> Основан на `docs/design.md`. Каждая фаза компилируется и тестируется самостоятельно.

---

# Phase 1: Project scaffold + Domain

## Цель
Инициализировать модуль и определить доменные типы.

## Файлы
- создать: `go.mod` (`go mod init github.com/ruauka/hw1`)
- создать: `internal/domain/survey.go`

## Что делаем
- `Question{ID int, Text string}` с `json:"..."` тегами
- `Answer{QuestionID int, Text string}` с `json:"..."` тегами

## Тесты
Нет — чистые типы-данные, тестировать нечего.

## Критерий Done
- [ ] `go build ./...` проходит без ошибок
- [ ] `make lint` чист

## Статус
[x] Выполнена

---

# Phase 2: Usecase

## Цель
Реализовать бизнес-логику: хардкоженный список вопросов и сохранение ответов через интерфейс.

## Файлы
- создать: `internal/usecase/survey.go`
- создать: `internal/usecase/survey_test.go`

## Что делаем
- Интерфейс `Store` с методом `Save(answers []domain.Answer) error` — определяем здесь (где потребляется)
- Struct `Survey` с полем `store Store`
- Конструктор `New(store Store) *Survey`
- Метод `Questions() []domain.Question` — возвращает хардкоженный список из 5 вопросов
- Метод `Submit(answers []domain.Answer) error` — вызывает `s.store.Save(answers)`

## Тесты
- `TestSurvey_Questions` — проверить длину и непустоту ответа
- `TestSurvey_Submit_OK` — mock Store возвращает nil, Submit возвращает nil
- `TestSurvey_Submit_StoreError` — mock Store возвращает ошибку, Submit пробрасывает её

Mock Store определяем inline в `_test.go` (не импортировать mocker-библиотек).

## Критерий Done
- [ ] `go test ./internal/usecase/...` проходит
- [ ] `make lint` чист

## Статус
[x] Выполнена

---

# Phase 3: Repository

## Цель
Реализовать in-memory хранилище ответов, защищённое от race conditions.

## Файлы
- создать: `internal/repository/store.go`
- создать: `internal/repository/store_test.go`

## Что делаем
- Unexported struct `store` с полями `mu sync.Mutex` и `answers [][]domain.Answer`
- Конструктор `New() *store`
- Метод `Save(answers []domain.Answer) error` — lock/unlock + append
- Compile-time check: `var _ usecase.Store = (*store)(nil)`

## Тесты
- `TestStore_Save` — сохранить ответы, проверить что `len(answers) == 1`
- `TestStore_Save_Multiple` — сохранить 3 раза, проверить `len == 3`
- `TestStore_Save_Concurrent` — 10 goroutine пишут параллельно, запустить с `-race`, итог `len == 10`

## Критерий Done
- [ ] `go test -race ./internal/repository/...` проходит
- [ ] `make lint` чист

## Статус
[x] Выполнена

---

# Phase 4: Handler

## Цель
Реализовать HTTP-обработчики для двух endpoints.

## Файлы
- создать: `internal/handler/survey.go`
- создать: `internal/handler/survey_test.go`

## Что делаем
- Интерфейс `Usecase` с методами `Questions() []domain.Question` и `Submit([]domain.Answer) error` — определяем здесь
- Struct `Handler` с полем `uc Usecase`
- Конструктор `New(uc Usecase) *Handler`
- `Handler.Questions(w, r)`:
  - Устанавливает `Content-Type: application/json`
  - `json.NewEncoder(w).Encode(uc.Questions())`
- `Handler.Submit(w, r)`:
  - `json.NewDecoder(r.Body).Decode(&answers)` → при ошибке `400 Bad Request`
  - `uc.Submit(answers)` → при ошибке `500 Internal Server Error`
  - При успехе `200 OK`

## Тесты (table-driven)

`TestHandler_Questions`:
- `questions_ok` — GET, ожидаем 200, непустой JSON-массив

`TestHandler_Submit`:
- `submit_ok` — POST с валидным JSON, ожидаем 200
- `submit_invalid_json` — POST с `"not json"`, ожидаем 400
- `submit_empty_body` — POST без тела, ожидаем 400
- `submit_store_error` — mock Usecase.Submit возвращает ошибку, ожидаем 500

Использовать `httptest.NewRecorder()` и `httptest.NewRequest()`.
Mock Usecase определить inline в `_test.go`.

## Критерий Done
- [ ] `go test ./internal/handler/...` проходит
- [ ] `make lint` чист

## Статус
[x] Выполнена

---

# Phase 5: Integration (web + main)

## Цель
Собрать приложение: встроить фронтенд и запустить сервер.

## Файлы
- создать: `web/embed.go`
- создать: `web/index.html`
- создать: `cmd/server/main.go`

## Что делаем

**`web/embed.go`:**
```
//go:embed index.html
var FS embed.FS
```

**`web/index.html`:**
- JS загружает `GET /questions`
- Рендерит `<textarea>` для каждого вопроса
- Кнопка отправляет `POST /answers`
- После ответа показывает "Спасибо!"

**`cmd/server/main.go`:**
- Ручной DI: `repository.New() → usecase.New() → handler.New()`
- `mux.HandleFunc("GET /questions", h.Questions)`
- `mux.HandleFunc("POST /answers", h.Submit)`
- `mux.Handle("/", http.FileServerFS(web.FS))`
- `http.ListenAndServe(":8080", mux)`

## Тесты
Нет unit-тестов — проверка вручную:
- `curl http://localhost:8080/questions`
- `curl -X POST http://localhost:8080/answers -d '[{"question_id":1,"text":"Go"}]'`
- Открыть браузер на `http://localhost:8080`

## Критерий Done
- [ ] `make build` компилируется без ошибок
- [ ] `make test` проходит весь suite
- [ ] `make lint` чист
- [ ] `curl /questions` возвращает JSON-массив вопросов
- [ ] `curl -X POST /answers` с валидным телом возвращает 200
- [ ] Браузер показывает форму и "Спасибо!" после отправки

## Статус
[x] Выполнена

---

## Review плана

### Проверено

| Аспект | Статус | Комментарий |
|--------|--------|-------------|
| **Edge cases** | ✅ | Пустое тело, невалидный JSON → 400; store error → 500 |
| **Конкурентность** | ✅ | `sync.Mutex` в Phase 3, `-race` в тесте |
| **Идиоматика** | ✅ | errors as values, no panic, `json.NewDecoder(r.Body)` |
| **Интерфейсы** | ✅ | Store в usecase, Usecase в handler; зависимости направлены внутрь |
| **Тесты** | ✅ | Table-driven в Phase 4 (4 кейса), mock inline в `_test.go` |
| **Compile-time check** | ✅ | `var _ usecase.Store = (*store)(nil)` в Phase 3 |
| **Метод ServeMux** | ✅ | Go 1.22+ синтаксис `"GET /questions"` — method mismatch → 405 автоматически |

### Предостережения при реализации

1. `json.NewDecoder(r.Body).Decode(&answers)` вернёт `io.EOF` для пустого тела — нужно обработать оба случая (`io.EOF` и другие ошибки) → оба ведут к 400.
2. В `store.Save()` делать `append` к копии среза (defensively copy), либо принимать как есть — при in-memory хранении достаточно append без копирования.
3. `http.FileServerFS(web.FS)` нужно зарегистрировать последним — мux-порядок важен для Go 1.22+.
4. В тестах handler использовать `httptest.NewRequest("POST", "/answers", strings.NewReader(...))`, не `http.NewRequest`.
