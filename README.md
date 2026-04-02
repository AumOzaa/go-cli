# Todo CLI

A command-line interface application for managing your todo tasks, built with Go.

## Overview

Todo CLI is a simple and efficient command-line tool that allows you to manage your tasks directly from the terminal. It stores your todos in a PostgreSQL database and displays them in a beautifully formatted table.

## Features

- Add new tasks with title and completion status
- List all tasks in a formatted table
- Mark tasks as completed (delete tasks)
- PostgreSQL database backend
- Beautiful table output using go-pretty

## Tech Stack

- **Go** - Programming language
- **[Cobra](https://github.com/spf13/cobra)** - CLI framework for Go
- **[go-pretty](https://github.com/jedib0t/go-pretty)** - Library for pretty printing tables
- **[pgx](https://github.com/jackc/pgx)** - PostgreSQL driver for Go
- **[godotenv](https://github.com/joho/godotenv)** - Library for loading .env files

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.26 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (any recent version)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/AumOzaa/go-cli
cd my-cli
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Database

Create a `.env` file in the `my-cli` directory with your PostgreSQL connection string:

```env
DATABASE_URL=postgres://username:password@localhost:5432/tododb?sslmode=disable
```

Replace the values with your PostgreSQL credentials:
- `username` - Your PostgreSQL username
- `password` - Your PostgreSQL password
- `localhost:5432` - Your PostgreSQL host and port
- `tododb` - Your database name

### 4. Create Database Table

Connect to your PostgreSQL database and create the required table:

```sql
CREATE TABLE tasklist (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    completed INTEGER NOT NULL
);
```

## Running the Application

### Linux/macOS

```bash
# Build the application
go build -o todo-cli

# Run
./todo-cli <command>
```

### Windows

```bash
# Build the application
go build -o todo-cli.exe

# Run
todo-cli.exe <command>
```

## Available Commands

### List All Tasks

```bash
./todo-cli list
```

Displays all tasks in a formatted table.

### Add a New Task

```bash
./todo-cli add <task_title> <completed_status>
```

Arguments:
- `<task_title>` - The title of your task (required)
- `<completed_status>` - 0 for incomplete, 1 for completed (required)

Example:
```bash
./todo-cli add "Learn Go" 0
```

### Mark Task as Completed

```bash
./todo-cli completed <task_id>
```

Arguments:
- `<task_id>` - The ID of the task to mark as completed (delete)

Example:
```bash
./todo-cli completed 1
```

## Usage Examples

```bash
# Add a new task
./todo-cli add "Build a CLI app" 0

# Add another task
./todo-cli add "Read a book" 0

# List all tasks
./todo-cli list

# Mark task #1 as completed
./todo-cli completed 1
```

## Project Structure

```
my-cli/
├── cmd/
│   ├── root.go       # Root command configuration
│   └── hello.go      # Command implementations
├── funcs/
│   ├── table.go      # Table rendering functions
│   └── getTodos.go   # Todo retrieval functions
├── internal/tools/
│   ├── database.go  # Database operations
│   └── dbConnection.go # Database connection utilities
├── models/
│   ├── todo.go      # Todo model
│   └── view.go      # View model
├── main.go          # Application entry point
```
