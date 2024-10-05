package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Silence(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Flags().GetString("name"))
	fmt.Println(cmd.Flags().GetString("password"))
}
