package tools

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func IniDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("URL : %v\n", os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int
	var task string
	var completed int
	rows, err := conn.Query(context.Background(), "SELECT * from taskList")
	fmt.Printf("Rows : \n %v", rows)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, task, completed)
}
