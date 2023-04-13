package gofm

import (
	"fmt"

	"github.com/merore/gofm/pkg/logger"
	"github.com/merore/gofm/pkg/missevan"
	"github.com/merore/gofm/pkg/openai"
)

type Config struct {
	Live          int
	MissevanToken string
	OpenAIToken   string
}

type Robot struct {
	name     string // The username of missevan account used by robot.
	liveName string // The username of live owner.
	openai   *openai.Client
	missevan *missevan.Client
	config   Config
}

func NewRobot(config Config) (*Robot, error) {
	mc, err := missevan.NewClient(config.MissevanToken)
	if err != nil {
		return nil, err
	}
	oc, err := openai.NewClient(config.OpenAIToken)
	if err != nil {
		return nil, err
	}
	r := &Robot{
		config:   config,
		openai:   oc,
		missevan: mc,
	}
	if err := r.init(); err != nil {
		return nil, err
	}
	return r, nil
}

// init get missevan info
func (r *Robot) init() error {
	fmt.Println(r.config)
	user, err := r.missevan.GetUserInfo()
	if err != nil {
		return err
	}
	r.name = user.Username

	live, err := r.missevan.GetLiveInfo(r.config.Live)
	if err != nil {
		return err
	}
	r.liveName = live.CreatorUsername
	return nil

}

func (r *Robot) Run() error {
	conn, err := r.missevan.Connect(r.config.Live)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("%s connect to %d successfully.", r.name, r.config.Live))
	for {
		msg, _ := conn.Read()
		r.dispatcher(msg)
	}
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
	return r.missevan.Send(r.config.Live, msg)
}
