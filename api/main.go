package main

import (
	"log"
	"net/http"
	"time"

	"gitthub.com/giulian-coding/cloudrift/internal/user"
)

func main() {
	userRepo := user.NewMemoryRepository()
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)

	mux := http.NewServeMux()
	userHandler.Register(mux)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Server is running...")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server failed to start", "error", err)
	}
}
