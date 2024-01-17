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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !clai.Config() {
			os.Exit(1)
		}
		c := clai.New(viper.GetString("apikey"))
		c.Work(args[0])
	},
}

var Verbose bool
var Disclaimer bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	rootCmd.PersistentFlags().BoolVarP(&Disclaimer, "disclaimer", "d", true, "show disclaimer")
	viper.BindPFlag("disclaimer", rootCmd.PersistentFlags().Lookup("disclaimer"))

	// rootCmd.Args = cobra.MinimumNArgs(1)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
