package cmd

import (
	"log/slog"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func Logger() *slog.Logger {
	verbose := viper.GetBool("verbose")

	plogger := pterm.DefaultLogger.WithTime(false)
	if verbose {
		plogger = plogger.WithLevel(pterm.LogLevelTrace).WithCaller(true)
	}

	// Create a new slog handler with the default PTerm logger
	handler := pterm.NewSlogHandler(plogger)

	// Create a new slog logger with the handler
	return slog.New(handler)
}
