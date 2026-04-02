/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/AumOzaa/go-cli/my-cli/funcs"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Calling the table functinon
		funcs.PrintTable()
		// title := args[0]
		// funcs.ListTodos()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		funcs.PrintTable()
	},
}
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Make a new Todo",
	Run: func(cmd *cobra.Command, args []string) {
		// Call your function from the funcs package
		// fmt.Println("Is this even being called?")
		title := args[0]
		completed, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Printf("An unknown error occured %v\n", err)
		}

		funcs.InsertTodo(title, completed)
	},
}

func init() {
	// fmt.Println("Over here")
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
