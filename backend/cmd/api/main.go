package main

import (
	"log"
	"net/http"

	httpx "github.com/szey/Mirai-Nihongo/backend/internal/http"
)

func main() {
	handler := httpx.NewHandler()
	router := httpx.NewRouter(handler)

	addr := ":8080"
	log.Printf("Mirai Nihongo backend listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
