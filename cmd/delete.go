package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yutian-9264/golang-redis-todolist/db"
)

var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "delete a todo",
	Run:   Delete,
}

func init() {
	deleteCMD.Flags().String("id", "", "id of the todo you want to delete")
	deleteCMD.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCMD)
}

func Delete(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	db.DeleteTodo(id)
}
