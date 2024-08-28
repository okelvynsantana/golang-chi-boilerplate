package main

import (
	"log"
	"mix-online-api/internal/server"
)

func main() {

	// Configura o roteador
	router := server.SetupRouter()

	// Cria e configura o servidor HTTP
	srv := server.NewServer(router)

	// Inicia o servidor
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
