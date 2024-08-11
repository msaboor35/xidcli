/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of xidcli",
	Long:  `This command will show the version of xidcli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("xidcli version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
