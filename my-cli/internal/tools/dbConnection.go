package tools

import (
	"context"
	"fmt"
	"github.com/AumOzaa/go-cli/my-cli/models"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func IniDB() *pgx.Conn {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func ListTodoss(conn *pgx.Conn) []models.Todo {
	var todos []models.Todo
	var todo models.Todo

	rows, err := conn.Query(context.Background(), "SELECT * FROM tasklist")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&todo.Id, &todo.Task, &todo.Completed)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		todos = append(todos, todo)
	}
	return todos
}

func InsertValues(conn *pgx.Conn, task string, completed int) {
	result, err := conn.Exec(
		context.Background(),
		"INSERT INTO tasklist (task,completed) VALUES ($1,$2)", task, completed)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task added successfully")
	fmt.Printf("%v\n", result)
}

func DeleteTask(conn *pgx.Conn, id int) {
	result, err := conn.Exec(
		context.Background(),
		"DELETE FROM tasklist WHERE id = ($1)", id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", result)
}
