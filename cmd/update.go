package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yutian-9264/golang-redis-todolist/db"
)

var updateCMD = &cobra.Command{
	Use:   "update",
	Short: "update description or status of todos",
	Run:   Update,
}

func init() {
	updateCMD.Flags().String("id", "", "id of the todo you want to update")
	updateCMD.MarkFlagRequired("id")
	updateCMD.Flags().String("description", "", "new description")
	updateCMD.Flags().String("status", "", "new status: completed, pending, in-progress")
	rootCmd.AddCommand(updateCMD)
}

func Update(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	desc := cmd.Flag("description").Value.String()
	status := cmd.Flag("status").Value.String()

	if desc == "" && status == "" {
		log.Fatalln("您没有输入想要修改的描述和状态")
	}

	if status == "completed" || status == "pending" || status == "in-progress" || status == "" {
		db.UpdateTodo(id, desc, status)
	} else {
		log.Fatalln("输入的状态无效")
	}
}
