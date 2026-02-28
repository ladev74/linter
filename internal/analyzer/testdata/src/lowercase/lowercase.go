package lowercase

import (
	"log/slog"
)

func lowercase() {
	slog.Info("Starting server") // want "^log message must start with lowercase letter$"
}
