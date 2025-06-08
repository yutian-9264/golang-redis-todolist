package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yutian-9264/golang-redis-todolist/db"
)

var createCMD = &cobra.Command{
	Use:   "create",
	Short: "create a todo with description",
	Run:   Create,
}

func init() {
	createCMD.Flags().String("description", "", "create a todo with description")
	createCMD.MarkFlagRequired("description")
	rootCmd.AddCommand(createCMD)
}

func Create(cmd *cobra.Command, args []string) {
	desc := cmd.Flag("description").Value.String()
	db.CreateTodo(desc)
}
