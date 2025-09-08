package cmd

import (
	"cli/todo/tui"
	"fmt"

	"github.com/spf13/cobra"
)

var detailsCmd = &cobra.Command{
	Use:     "details",
	Short:   "View task details",
	Long:    `View detailed information about a specific task by ID or title.`,
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		task, err := getTask(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		tui := tui.New()
		model := tui.NewDetailsModel(task, 0, 0)
		tui.Run(model)
	},
}

func init() {
	rootCmd.AddCommand(detailsCmd)
}
