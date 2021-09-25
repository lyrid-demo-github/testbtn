package main

import (
    "net/http"

	entry "go1x_chi.template/entry"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := entry.Initialize()
	http.ListenAndServe(":3000", r)
}