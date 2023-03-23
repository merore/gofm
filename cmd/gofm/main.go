package main

import (
	"os"

	"github.com/merore/gofm/pkg/logger"
)

const (
	EnvMissevanToken = "MISSEVAN_TOKEN"
	EnvOpenAIToken   = "OPENAI_TOKEN"
	EnvOpenAIAPI     = "OPENAI_API"
)

var (
	Version       string
	MissevanToken = os.Getenv(EnvMissevanToken)
	OpenAIToken   = os.Getenv(EnvOpenAIToken)
	OpenAIAPI     = os.Getenv(EnvOpenAIAPI)
)

func main() {
	if err := robotCmd.Execute(); err != nil {
		logger.Error(err)
	}
}
