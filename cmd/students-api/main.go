package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Harikrishnasinh/go-students-api/internal/config"
	"github.com/Harikrishnasinh/go-students-api/internal/http/handlers/student"
)

func main() {
	// Entry point for the students-api application
	fmt.Println("Hello welcome students!!!")
	cfg := config.MustLoad()

	// connection of database
	// > currently it is in pending state

	// creating routes
	routes := http.NewServeMux()

	routes.HandleFunc("POST /", student.New())

	// creating server
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: routes,
	}

	slog.Info("Server is listening on", "address", cfg.HttpServer.Address)

	// graceful shutdown, making done channel for notifying the shutdown signal i.e interrupt

	// Flow of Execution (Step by Step)
	// 1️⃣ Program starts
	// done channel is created
	// OS signal listener is registered

	// 2️⃣ Server starts
	// Server runs in a separate goroutine
	// Main goroutine continues

	// 3️⃣ Main goroutine waits
	// Execution stops at:
	// <-done

	// Program is idle but still running

	// 4️⃣ User presses Ctrl + C
	// OS sends SIGINT
	// Go delivers it to done channel

	// 5️⃣ Program unblocks
	// <-done receives the signal
	// Main function continues
	// Program exits (unless more code follows)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		// listening on the port (localhost:8080)
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Server failed to start", "error", err.Error())
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5 seconds
	defer cancel()
	doneContext := ctx.Done()
	slog.Info("Shutting down server... it will be down in 5 seconds")
	<-doneContext
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", "error", err.Error())
	}
	slog.Info("Server exited properly")
}
