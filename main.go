package main

import (
    "net/http"

	entry "YOUR_APP_NAME.YOUR_MODULE_NAME/YOUR_FUNCTION_NAME"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := entry.Initialize()
	http.ListenAndServe(":3000", r)
}