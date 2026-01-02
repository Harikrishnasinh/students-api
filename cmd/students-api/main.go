package main

import (
	"fmt"

	"github.com/Harikrishnasinh/go-students-api/internal/config"
)

func main() {
	// Entry point for the students-api application
	fmt.Println("Hello welcome students!!!")

	cfg := config.MustLoad()
	fmt.Printf("Loaded config: %+v\n", cfg)
}
