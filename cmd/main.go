package main

import (
	"log/slog"
	"os"
)

func main() {

	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := &application{
		config: cfg,
	}

	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//log estruturado
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("servidor falhou em inicia", "error", err)
		os.Exit(1)
	}
}
