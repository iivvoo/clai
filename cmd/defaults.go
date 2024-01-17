package cmd

import (
	"fmt"

	"github.com/iivvoo/clai/clai"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Show defaults",
	Long:  "Show defaults that can be used or changed in your config file",
	Run: func(cmd *cobra.Command, args []string) {
		Defaults()
	},
}

func Defaults() {
	fmt.Println("Available models")
	for k := range clai.Models {
		fmt.Println(k)
	}
	fmt.Println("\nDefault model: GPT3Dot5Turbo")
	fmt.Println("Default System Prompt:")
	fmt.Println(clai.SystemMessage)
}

func init() {
	rootCmd.AddCommand(defaultCmd)
}
