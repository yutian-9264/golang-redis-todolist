package cmd

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yutian-9264/golang-redis-todolist/db"
)

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "list all todos",
	Run:   List,
}

func init() {
	listCMD.Flags().String("status", "", "completed, pending or in-progress")
	rootCmd.AddCommand(listCMD)
}

func List(cmd *cobra.Command, args []string) {
	status := cmd.Flag("status").Value.String()

	var todos []db.Todo
	if status == "completed" || status == "pending" || status == "in-progress" || status == "" {
		todos = db.ListTodos(status)
	} else {
		log.Fatal("invalid status")
	}

	todoTable := [][]string{}
	for _, todo := range todos {
		todoTable = append(todoTable, []string{todo.ID, todo.Desc, todo.Status})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Description", "Status"})

	for _, v := range todoTable {
		table.Append(v)
	}
	table.Render()
}
