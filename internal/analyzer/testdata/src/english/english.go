package english

import (
	"log/slog"
)

func english() {
	slog.Info("запуск сервера") // want "^log message must be in English$"
}
