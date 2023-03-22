package main

import (
	"github.com/merore/gofm/pkg/logger"
)

func main() {
	if err := robotCmd.Execute(); err != nil {
		logger.Error(err)
	}
}
