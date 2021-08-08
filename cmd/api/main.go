package main

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"

	"github.com/deepakvashist/golang-fx-sample/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Module = fx.Options(
	fx.Provide(internal.NewRepository),
	fx.Provide(internal.NewUsecase),
	fx.Provide(internal.NewRestInfra),
	fx.Invoke(runHTTPServer),
)

func runHTTPServer(lifecycle fx.Lifecycle, restInfra *internal.RestInfra) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(c context.Context) error {
				go func() {
					r := chi.NewRouter()
					r.Use(middleware.Logger)
					r.Get("/users/{id}", restInfra.GetByID)
					http.ListenAndServe(":8000", r)
				}()

				return nil
			},
		},
	)
}

func main() {
	app := fx.New(Module)
	err := app.Err()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
