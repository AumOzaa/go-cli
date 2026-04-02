package funcs

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func printTable() {
	fmt.Println("In the tables function")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "Task", "Completed"})

	t.AppendRows([]table.Row{
		{1, "bath", 0},
		{2, "GYM", 1},
	})
	t.AppendSeparator()
	t.Render()
}
