/*
Copyright © 2025 Adameji Israel <ireoluwa48@gmail.com>
*/
package cmd

import (
	"fmt"

	"cli/todo/tui"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get the list of all tasks",
	Long:    `Get the list of all tasks without any filter.`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		tasks := getTasks()

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Launch Bubble Tea TUI for task list
		tui := tui.New()
		model := tui.NewListModel(tasks)
		tui.Run(model)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
