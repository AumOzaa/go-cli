package funcs

import (
	"fmt"

	"github.com/AumOzaa/go-cli/my-cli/internal/tools"
	"github.com/AumOzaa/go-cli/my-cli/models"
)

func ListTodos() []models.Todo {
	fmt.Printf("Exisiting todos :\n %v\n", tools.MockTodos)
	return tools.MockTodos
}
