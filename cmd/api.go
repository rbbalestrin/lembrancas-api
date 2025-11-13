package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	//logger
	//db driver
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//middleware
	r.Use(middleware.RequestID) // rate limiting
	r.Use(middleware.RealIP)    // rate limiting e analytics e tracing
	r.Use(middleware.Logger)    // logger bonito
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("não tao no teto... ta suave."))
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server iniciado em addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string //endereço da api ex 8000
	db   dbConfig
}

type dbConfig struct {
	dsn string //domain string para conectar na database
}
