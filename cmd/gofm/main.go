package main

import (
	"github.com/merore/gofm"
	"github.com/merore/gofm/pkg/logger"
)

var defaultConfigPath = "config.yaml"

func main() {
	logger.Info("Start gofm")
	c, _ := gofm.UnmarshalConfig(defaultConfigPath)
	s := gofm.NewRobot(c)
	s.Run()
}
