package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/merore/gofm"
	"github.com/merore/gofm/pkg/logger"
	"github.com/merore/gofm/pkg/openai"
)

func main() {
	config, _ := gofm.UnmarshalConfig("config.yaml")
	logger.Debug(config)
	c := openai.NewClient(config.OpenAi.Key, config.OpenAi.Proxy) // TODO token
	c.Reset(openai.DefaultPrompt)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("[Sven]    ")
		bs, _ := reader.ReadBytes('\n')
		fmt.Printf("[Summer]  ")
		cs, err := c.ChatStream(string(bs))
		if err != nil {
			logger.Error(err)
			continue
		}
		for {
			s, closed := <-cs
			if !closed {
				fmt.Printf("\n")
				break
			}
			fmt.Printf(s)
		}
	}
}
