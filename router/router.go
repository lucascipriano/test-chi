package router

import (
	"net/http"

	"apirest.com/controller"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func StartRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controller.GetAlgumacoisa)

	http.ListenAndServe(":3000", r)
}
