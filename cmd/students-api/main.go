package main

import (
	"fmt"
	"net/http"

	"github.com/Harikrishnasinh/go-students-api/internal/config"
)

func main() {
	// Entry point for the students-api application
	fmt.Println("Hello welcome students!!!")
	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)

	// connection of database
	// > currently it is in pending state

	// creating routes
	routes := http.NewServeMux()

	routes.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling root route")
		w.Write([]byte("Welcome to Students API"))
	})

	// creating server
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: routes,
	}

	fmt.Printf("Server is listening on %s", cfg.HttpServer.Address)

	// listening on the port (localhost:8080)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}

}
