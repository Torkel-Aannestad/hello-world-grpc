package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"golang.org/x/net/context"
)

type config struct {
	port int
}

type application struct {
	config *config
	logger *slog.Logger
}

func run(_ context.Context) error {
	logOpt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpt))

	//create config
	cfg := &config{
		port: 4000,
	}

	//create application
	app := application{
		config: cfg,
		logger: logger,
	}

	//create muxer
	router := app.routes()

	//create http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.config.port),
		Handler: router,
	}

	app.logger.Info("starting server", "port", app.config.port)
	//start server
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("failed to start the app", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
