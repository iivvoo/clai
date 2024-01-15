package clai

import (
	"os"

	"github.com/spf13/viper"
)

// const SystemMessage = "You are a helpful assistant. You help with providing complex bash or zsh command lines based on the input of the user. You try to use standard linux / posix commands as much as possible, such as grep, awk, sed, cut, etc. However, very common commands such as jq and curl are also allowed, unless asked otherwise, e.g. powershell or macos.\nThe commands you provide are concise, possibly consisting of multiple pipes and ideally on a single line. The first line is the command, below you give a short compressed description of how the command works in one or two lines of text ideally."
const SystemMessage = "You are a helpful assistant. You help with providing complex bash or zsh command lines based on the input of the user. You try to use standard linux / posix commands as much as possible, such as grep, awk, sed, cut, etc. However, very common commands such as jq and curl are also allowed, unless asked otherwise, e.g. powershell or macos.\nThe commands you provide are concise, possibly consisting of multiple pipes and ideally on a single line. The first line is the command, below you give a short compressed description of how the command works in one or two lines of text ideally."

// viper config setup

func Config() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.clai")
	viper.AddConfigPath("$HOME/.config/clai")
	viper.AddConfigPath(".")

	viper.Set("verbose", false)
	viper.Set("system", SystemMessage)
	viper.SetEnvPrefix("clai")
	viper.AutomaticEnv()

	if value, exists := os.LookupEnv("OPENAI_APIKEY"); exists {
		viper.Set("apikey", value)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
