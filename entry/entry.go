package YOUR_FUNCTION_NAME

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
)

func Initialize() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	FileServerWithIndexHandlerFallback(r, os.Getenv("PUBLIC_PATH"), os.Getenv("DIST_FOLDER"))

	return r
}

func FileServerWithIndexHandlerFallback(r chi.Router, path string, localpath string) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, localpath))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(res http.ResponseWriter, req *http.Request) {
		// find the file in the static file dir first
		f, err := filesDir.Open(req.URL.Path)
		if err != nil {
			// if not found then serve the index html with the route....let the frontend takes care of it
			http.ServeFile(res, req, filepath.Join(filepath.Join(workDir, localpath), "/"+os.Getenv("INDEX_HOME")))
			return
		}
		defer f.Close()
		http.FileServer(filesDir).ServeHTTP(res, req)
	})
}
