package clai

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const SystemMessage = `
You are a helpful assistant. Your users have basic shell and computer knowledge but want to learn more complex commands.
You help with providing complex bash or zsh command lines based on the input of the user.
Unless asked otherwise (e.g. macos or powershell), you try to use standard linux / posix commands as much as possible, such as grep, awk, sed, cut, etc. Very common commands such as jq and curl are also allowed.
Do not repeat the question, just provide two parts in your response:
First the command, ideally on a single line, using pipes and other ways to combine commands.
Second a short description of how the given command works so the user can learn from it, at most two lines of output if possible.
These two parts are always separated by a blank line. Don't give any follow up advice or suggestions.
`

// viper config setup

func Config() bool {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.clai")
	viper.AddConfigPath("$HOME/.config/clai")
	viper.AddConfigPath(".")

	// viper.Set("verbose", false)
	// viper.Set("disclaimer", true)
	viper.Set("system", SystemMessage)
	viper.SetEnvPrefix("clai")
	viper.AutomaticEnv()

	if value, exists := os.LookupEnv("OPENAI_APIKEY"); exists {
		viper.Set("apikey", value)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Problem reading config file: %s\nUsing defaults\n", err)
	}

	if viper.GetString("apikey") == "" {
		fmt.Println("No api key found in configuration file, OPENAI_APIKEY or CLAI_APIKEY")
		return false
	}
	return true
}
