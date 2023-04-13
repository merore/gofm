package main

import (
	"flag"

	"github.com/merore/gofm"
	"github.com/merore/gofm/pkg/logger"
)

var (
	Version       string
	live          int
	missevanToken string
	openaiToken   string
)

func init() {
	flag.IntVar(&live, "live", 0, "--live=82727192")
	flag.StringVar(&missevanToken, "missevan-token", "", "--missevan-token=TOKEN")
	flag.StringVar(&openaiToken, "openai-token", "", "--openai-token=TOKEN")
	flag.Parse()
}

func main() {
	c := gofm.Config{
		Live:          live,
		MissevanToken: missevanToken,
		OpenAIToken:   openaiToken,
	}
	r, err := gofm.NewRobot(c)
	if err != nil {
		logger.Error(err)
		return
	}
	if err := r.Run(); err != nil {
		logger.Error(err)
	}
}
