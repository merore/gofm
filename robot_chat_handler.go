package gofm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/merore/gofm/pkg/logger"
	"github.com/merore/gofm/pkg/missevan"
	"github.com/merore/gofm/pkg/openai"
)

/*
一条 new message 消息的类型
- user []
- user [@someone content]
- user [@someone command]
*/

type NewMessageType int

const (
	CommonContent NewMessageType = iota
	MentiondContent
	MentiondCommand
)

var commands map[string]func()

const (
	CommandReset = "reset"
)

const chatRegexpStr = `^@([\p{Han}\w]+)\s+(.*)`

var reg *regexp.Regexp

func init() {
	var err error
	reg, err = regexp.Compile(chatRegexpStr)
	if err != nil {
		panic(err)
	}
}

type ChatMessage struct {
	User     string // who send the message
	Type     NewMessageType
	Mentiond string // @someone
	Content  string
}

func ParseChatMessage(msg missevan.FMMessage) (nm ChatMessage) {
	nm.User = msg.User.Username
	match := reg.FindStringSubmatch(msg.Message)
	if match == nil || len(match) != 3 {
		nm.Type = CommonContent
		nm.Content = msg.Message
		return
	}
	nm.Type = MentiondContent
	nm.Mentiond = match[1]
	nm.Content = match[2]

	// command
	if nm.Content == CommandReset {
		nm.Type = MentiondCommand
	}
	return
}

func (r *Robot) chatHandler(msg missevan.FMMessage) {
	logger.Info(msg.User.Username, "\n", msg.Message)
	nm := ParseChatMessage(msg)
	switch nm.Type {
	case CommonContent:
		r.commonContentHandler(nm)
	case MentiondContent:
		r.mentiondContentHandler(nm)
	case MentiondCommand:
		r.mentiondCommandHandler(nm)
	}
}

// mentiondContentHandler Handle the new message which begin with @someone.
// Trigger chatGPT when the robot is mentioned.
func (r *Robot) mentiondContentHandler(nm ChatMessage) {
	if nm.Mentiond != r.name {
		return
	}
	ans, err := r.openai.Chat(nm.Content)
	if err != nil {
		logger.Error(err)
		return
	}
	// trim " " and "\n" in ChatGPT return
	ans = strings.Trim(ans, " ")
	ans = strings.Trim(ans, "\n")
	// @someone message
	ans = fmt.Sprintf("@%s\n", nm.User) + ans
	if err := r.Send(ans); err != nil {
		logger.Error(err)
	}
	logger.Info("gofm", ans)
}

func (r *Robot) mentiondCommandHandler(nm ChatMessage) {
	if nm.Content == CommandReset {
		r.resetCommand(nm)
	}
}

func (s *Robot) commonContentHandler(nm ChatMessage) {
	return
}

func (r *Robot) resetCommand(nm ChatMessage) {
	logger.Debug("reset", nm)
	r.openai.Reset(openai.DefaultPrompt)
}
