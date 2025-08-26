package main

import (
	"log/slog"
	"net/http"
	"os"

	"golang.org/x/net/context"
)

type config struct {
	port int
	env  string
}

type application struct {
	config *config
	logger *slog.Logger
}

func run(ctx context.Context) error {
	logOpt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpt))

	//create config
	cfg := &config{}

	//create application
	app := application{
		config: cfg,
		logger: logger,
	}

	//create http server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: http.HandlerFunc(app.getPostsHandler),
	}

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
