/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/msaboor35/xidcli/internal/xid"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate xid",
	Long:  `Generate a new xid.`,
	Run: func(cmd *cobra.Command, args []string) {
		xid := xid.GenerateXID()
		fmt.Println(xid)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
