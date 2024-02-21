package main

import (
	"context"
	"fmt"

	"main.go/api"
	"main.go/config"
	"main.go/service"
	"main.go/storage/postgres"
)

func main() {
	cfg := config.Load()
	fmt.Println("d")

	pgStore, err := postgres.New(context.Background(), cfg)
	fmt.Println("s")
	if err != nil {
		return
	}
	fmt.Println("rrr")
	defer pgStore.Close()

	services := service.New(pgStore)

	server := api.New(services)

	if err = server.Run("localhost:8058"); err != nil {
		panic(err)
	}
}
