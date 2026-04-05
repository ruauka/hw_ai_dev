# CLAUDE.md

## Проект
- Проект: 
  - Backend: Go 1.24
  - Frontend: Чистый HTML + JS, React/Vue
- Стек: [Базовый ServeMux, docker]
- Архитектура: [clean architecture]

## Команды
- `make build`  — сборка
- `make test`   — запуск тестов
- `make lint`   — golangci-lint

## Coding standards
- Go: идиоматичный Go, errors as values, no panic в lib-коде
- Коммиты: Conventional Commits format

## Что Claude делает неправильно
- Не используй --race только для lint, всегда запускай с тестами
- При изменении одного файла не гоняй весь test suite: `go test ./path/...`

## Скиллы cc-skills-golang

### В фазах workflow (вызывать всегда)
| Фаза | Скилл | Когда |
|------|-------|-------|
| `/research` | `/golang-project-layout` | Перед проектированием структуры |
| `/design` | `/golang-structs-interfaces` | При проектировании интерфейсов |
| `/design` | `/golang-naming` | При именовании типов/функций/пакетов |
| `/design` | `/golang-design-patterns` | При выборе паттернов |
| `/design` | `/golang-dependency-injection` | При проектировании DI |
| `/implement` | `/golang-error-handling` | Перед реализацией error handling |
| `/implement` | `/golang-testing` | Перед написанием тестов |
| `/implement` | `/golang-safety` | Перед реализацией каждой фазы |
| после фазы | `/golang-code-style` | Code review после реализации |
| после фазы | `/golang-modernize` | Проверить на идиоматичность |

### По контексту (вызывать при необходимости)
| Тема | Скилл |
|------|-------|
| goroutines / channel / mutex | `/golang-concurrency` |
| context.Context | `/golang-context` |
| логирование / метрики | `/golang-observability`, `/golang-samber-slog` |
| производительность | `/golang-performance`, `/golang-benchmark`, `/golang-data-structures` |
| работа с БД | `/golang-database` |
| gRPC | `/golang-grpc` |
| CLI-инструменты | `/golang-cli` |
| security review | `/golang-security` |
| CI/CD | `/golang-continuous-integration` |
| документация | `/golang-documentation` |
| дебаг / баги | `/golang-troubleshooting` |
| добавление зависимостей | `/golang-dependency-management`, `/golang-popular-libraries` |
| samber/lo, do, mo, ... | соответствующий `/golang-samber-*` |
| настройка линтера | `/golang-linter` |
