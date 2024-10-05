/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bravo",
	Short: "Silences database notifications",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: Silence,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "", "The name of the database")
	rootCmd.Flags().StringP("password", "p", "", "Password of the database")
}
