package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/hsmtkk/sturdy-giggle/repo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	reader := env.NewReader()
	pgConfig, err := reader.PostgresConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	conn, err := repo.ConnectPostgres(ctx, pgConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	userRepo := repo.NewUser(conn)
	alpha, err := userRepo.NewUser(ctx, "alpha", "alhpa@example.com", "secret")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", alpha)
}
