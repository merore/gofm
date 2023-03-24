package gofm

import (
	"fmt"

	"github.com/merore/gofm/pkg/logger"
	"github.com/merore/gofm/pkg/missevan"
	"github.com/merore/gofm/pkg/openai"
)

type Robot struct {
	name     string // The username of missevan account used by robot.
	openai   *openai.Client
	missevan *missevan.Client
	config   Config
}

func NewRobot(config Config) *Robot {
	s := &Robot{
		config:   config,
		openai:   openai.NewClient(config.OpenAIToken, config.OpenAIAPI),
		missevan: missevan.NewClient(config.MissevanToken),
	}
	s.openai.Reset(openai.DefaultPrompt)
	return s
}

func (r *Robot) Run() error {
	user, err := r.missevan.GetUserInfo()
	if err != nil {
		return err
	}
	r.name = user.Username
	msgs := r.missevan.Connect(r.config.MissevanLive)
	logger.Info(fmt.Sprintf("%s connect to %d successfully.", r.name, r.config.MissevanLive))
	for msg := range msgs {
		r.dispatcher(msg)
	}
	return nil
}

func (r *Robot) dispatcher(msg missevan.FMMessage) {
	logger.Debug(msg)
	switch {
	case msg.Type == missevan.TypeMessage && msg.Event == missevan.EventNew:
		r.chatHandler(msg)
	case msg.Type == missevan.TypeMember && msg.Event == missevan.EventJoinQueue:
		r.joinHandler(msg)
	}
}

func (r *Robot) Send(msg string) error {
	return r.missevan.Send(r.config.MissevanLive, msg)
}
