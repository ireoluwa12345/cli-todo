package cmd

import (
	"cli/todo/tui"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create task",
	Long:    `Create a new task with required details.`,
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {

		t := tui.New()
		model := t.NewCreateModel()
		t.Run(model)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
