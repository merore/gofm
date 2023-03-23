package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/merore/gofm"
	"github.com/merore/gofm/pkg/missevan"
	"github.com/spf13/cobra"
)

var (
	phoneReg         = regexp.MustCompile(`^1\d{10}$`)
	InvalidParameter = errors.New("invalid parameter")
)

func init() {
	robotCmd.AddCommand(tokenCmd)
}

var robotCmd = &cobra.Command{
	Use:   "gofm",
	Short: "Gofm is a robot for missevan.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := gofm.Config{
			MissevanToken: MissevanToken,
			OpenAIToken:   OpenAIToken,
			OpenAIAPI:     OpenAIAPI,
		}
		live, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		s := gofm.NewRobot(config, live)
		return s.Run()
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get a token from missevan.",
	RunE: func(cmd *cobra.Command, args []string) error {
		strs := phoneReg.FindAllString(args[0], -1)
		if strs == nil {
			return InvalidParameter
		}
		phone, _ := strconv.Atoi(strs[0])
		token, err := missevan.NewToken(phone, args[1])
		if err != nil {
			return err
		}
		fmt.Println(token)
		return nil
	},
}
