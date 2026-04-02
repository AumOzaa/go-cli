package funcs

import (
	"os"

	"github.com/AumOzaa/go-cli/my-cli/internal/tools"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable() {
	conn := tools.IniDB()
	todos := tools.ListTodoss(conn)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "Task", "Completed"})
	for i := 0; i < len(todos); i++ {
		t.AppendRow(table.Row{todos[i].Id, todos[i].Task, todos[i].Completed})
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			WidthMin: 90,
		},
	})
	t.AppendSeparator()
	t.SetAllowedRowLength(90)
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func InsertTodo(task string, completed int) {
	conn := tools.IniDB()
	tools.InsertValues(conn, task, completed)
}

func DeleteATodo(id int) {
	conn := tools.IniDB()
	tools.DeleteTask(conn, id)
}
