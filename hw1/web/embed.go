// Пакет web содержит встроенные статические файлы фронтенда.
package web

import "embed"

// FS — файловая система со встроенным фронтендом.
//
//go:embed index.html
var FS embed.FS
