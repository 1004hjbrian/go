package main

import (
	"log"
	"net/http"
	"time"
	"web10/decoHandler"
	"web10/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Complete time", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := decoHandler.NewDecoHandler(mux, logger)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
