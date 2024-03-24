package logging

import (
	"os"

	"github.com/rs/zerolog"
)

// zerolog logger instance
var logger zerolog.Logger

func init() {
    // Create a new zerolog logger with desired output (in this case, console)
    logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}

func DefaultLogger() zerolog.Logger {
	return logger
}