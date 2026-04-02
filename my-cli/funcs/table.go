package funcs

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable() {
	todos := ListTodos()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "Task", "Completed"})
	for i := 0; i < len(todos); i++ {
		t.AppendRow(table.Row{todos[i].Id, todos[i].Task, todos[i].Completed})
	}
	t.AppendSeparator()
	t.Render()
}
