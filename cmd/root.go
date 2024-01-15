package cmd

import (
	"os"

	"github.com/iivvoo/clai/clai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "clai",
	Short: "Command Line AI",
	Long:  `clai helps you write and learn more complex shell commands`,
	Run: func(cmd *cobra.Command, args []string) {
		clai.Config()
		c := clai.New(viper.GetString("apikey"))
		c.Work(args[0])
	},
}

var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// add positional arguments
	rootCmd.Args = cobra.MinimumNArgs(1)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
