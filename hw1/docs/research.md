# Research: Мини-анкета (Mini Questionnaire)

## Тип задачи
[x] Новый проект — HTTP-сервис + фронтенд

---

## Контекст

Изучено:
- `README.md` — требования к проекту
- `.claude/CLAUDE.md` — стек, архитектура, команды
- Makefile — build/test/lint команды
- Пустая директория `docs/`
- golang-project-layout skill — рекомендации по структуре

Стек по CLAUDE.md: Go 1.24, базовый ServeMux, Docker, Clean Architecture.
Существующего кода нет — чистый старт.

---

## Понимание задачи

Простое full-stack приложение "анкета":

**Backend (Go):**
- `GET /questions` — возвращает список 3-5 захардкоженных вопросов в JSON
- `POST /answers` — принимает JSON с ответами, хранит в памяти (slice/map)

**Frontend (HTML + JS):**
- Загружает вопросы с backend
- Отображает форму
- Отправляет ответы через POST /answers
- Показывает "Спасибо!" после отправки

**Хранилище:** только in-memory, данные теряются при рестарте.

---

## Затронутые слои clean architecture

- **Entity (domain):**
  - `Question` — {ID, Text} — хардкодед список вопросов
  - `Answer` — {QuestionID, Text} — ответ пользователя
  - `Survey` — набор ответов ([]Answer)

- **Usecase:**
  - `GetQuestions() []Question`
  - `SubmitAnswers(answers []Answer) error`

- **Repository:**
  - Интерфейс `SurveyRepository` с методом `Save(answers []Answer) error`
  - Реализация `inMemoryRepo` (sync.Mutex + slice)

- **Handler / ServeMux:**
  - `GET /questions` → JSON список вопросов
  - `POST /answers` → принять JSON, сохранить, ответить 200 OK
  - `GET /` → встроенный фронтенд (go:embed)

---

## Структура проекта

```
hw1/
├── cmd/server/main.go          # точка входа, wiring
├── internal/
│   ├── domain/
│   │   └── survey.go           # entity: Question, Answer
│   ├── usecase/
│   │   ├── survey.go           # бизнес-логика
│   │   └── survey_test.go
│   ├── repository/
│   │   ├── store.go            # in-memory реализация
│   │   └── store_test.go
│   └── handler/
│       ├── survey.go           # HTTP handlers
│       └── survey_test.go
├── web/
│   ├── index.html              # фронтенд
│   └── embed.go                # go:embed директива
├── go.mod
├── Makefile
└── docs/
    └── research.md
```

---

## Go-пакеты / зависимости

**Stdlib only:**
- `net/http` — HTTP сервер, ServeMux
- `encoding/json` — JSON encode/decode
- `embed` — встраивание фронтенда
- `sync` — Mutex для thread-safe хранилища
- `testing` — тесты

**Внешних зависимостей нет.**

---

## ServeMux маршруты

```
GET  /questions    → questionsHandler
POST /answers      → answersHandler
GET  /             → fs.FS (embed static frontend)
```

---

## Риски и неизвестные

1. **Thread safety** — in-memory store должен быть защищён `sync.Mutex`. Без этого — data race при конкурентных запросах.
2. **CORS** — решается раздачей фронтенда с того же сервера через `go:embed`.
3. **Метод проверки на ServeMux** — Go 1.22+ поддерживает `GET /path` синтаксис для method-based routing, что упрощает код.

---

## Уточняющие вопросы и ответы

1. **Формат вопросов** → Свободный текст (textarea)
2. **Frontend serving** → `go:embed` в бинарник
3. **Docker** → Не нужен, только локальный запуск
4. **GET /answers** → Не нужен, только POST

---

## Итоговое понимание задачи

Go HTTP-сервис (ServeMux) с clean architecture:

- `GET /questions` → JSON с 3-5 вопросами свободного текста
- `POST /answers` → принять JSON-массив ответов, сохранить в `sync.Mutex`-защищённый slice
- `GET /` → встроенный `go:embed` HTML/JS фронтенд

Frontend: один `index.html` с JS — загружает вопросы, рендерит textarea, POST ответы, показывает "Спасибо!".

Тесты: unit-тесты для usecase и handler слоёв.

---

## Следующий шаг

`/design` — проектирование интерфейсов, entity, sequence diagram
