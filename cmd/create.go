package cmd

import "github.com/spf13/cobra"

var createCMD = &cobra.Command{Use: "create", Short: "create a todo with description", Run: Create}

func Create(cmd *cobra.Command, args []string) {
	desc := cmd.Flag("description").Value.String()
	db.Create(desc)
}
