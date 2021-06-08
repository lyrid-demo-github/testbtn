package YOUR_FUNCTION_NAME

import (
	"github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "net/http"
)

func Initialize() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    		w.Write([]byte("hello world"))
    	})

	return r
}