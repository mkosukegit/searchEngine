package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	db "search/src/middleware/db"
	web "search/src/middleware/web"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	envLoad()

	db.DbConnect()

	web.Connect()

	log.Print("service start ...")
	err := http.ListenAndServe(":8002", nil)

	if err != nil {
		if ctx.Err() != nil {
			fmt.Printf("error: %v\n", ctx.Err())
		} else {
			fmt.Printf("error: %v\n", err)
		}
		return 1
	}

	log.Println("shut down ...")
	return 0
}

func envLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
		return
	}

	log.Print("get env ...")
}
