package clai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

type Clai struct {
	Key string
}

func New(key string) *Clai {
	return &Clai{
		Key: key,
	}
}

// get model: map string from viper to openai model, warn and default if not found

func (c *Clai) GetModel() string {
	m := map[string]string{

		"GPT432K0613":           openai.GPT432K0613,
		"GPT432K0314":           openai.GPT432K0314,
		"GPT432K":               openai.GPT432K,
		"GPT40613":              openai.GPT40613,
		"GPT40314":              openai.GPT40314,
		"GPT4TurboPreview":      openai.GPT4TurboPreview,
		"GPT4VisionPreview":     openai.GPT4VisionPreview,
		"GPT4":                  openai.GPT4,
		"GPT3Dot5Turbo1106":     openai.GPT3Dot5Turbo1106,
		"GPT3Dot5Turbo0613":     openai.GPT3Dot5Turbo0613,
		"GPT3Dot5Turbo0301":     openai.GPT3Dot5Turbo0301,
		"GPT3Dot5Turbo16K":      openai.GPT3Dot5Turbo16K,
		"GPT3Dot5Turbo16K0613":  openai.GPT3Dot5Turbo16K0613,
		"GPT3Dot5Turbo":         openai.GPT3Dot5Turbo,
		"GPT3Dot5TurboInstruct": openai.GPT3Dot5TurboInstruct,
	}

	if model, found := m[viper.GetString("model")]; found {
		return model
	}

	if viper.GetBool("verbose") && viper.GetString("model") != "" {
		fmt.Printf("Model %s not found, defaulting to %s\n", viper.GetString("model"), openai.GPT3Dot5Turbo)
	}
	return openai.GPT3Dot5Turbo
}

func (c *Clai) Work(query string) {
	client := openai.NewClient(c.Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Model: openai.GPT3Dot5Turbo,
			// Model: openai.GPT4,
			Model: c.GetModel(),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: viper.GetString("system"),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

}
