package gofm

import (
	"fmt"

	"github.com/merore/gofm/pkg/missevan"
	"github.com/mozillazg/go-pinyin"
)

const (
	AnonUser = "匿名大佬"
)

var pyArg pinyin.Args

func init() {
	pyArg = pinyin.NewArgs()
	pyArg.Fallback = func(r rune, a pinyin.Args) []string {
		return nil
	}
	pyArg.Style = pinyin.Tone
	pyArg.Separator = " "
}

func (r *Robot) joinHandler(msg missevan.FMMessage) {
	for i := range msg.Queue {
		name := msg.Queue[i].Username
		// anonymous user
		if name == "" {
			name = AnonUser
		}
		wm := welcomeMessage(name)
		r.Send(wm)
	}
}

func welcomeMessage(name string) string {
	py := hansToPinyin(name)
	msg := fmt.Sprintf("欢迎@%s 进入直播间\n\n[%s]", name, py)
	return msg
}

func hansToPinyin(name string) string {
	return pinyin.Slug(name, pyArg)
}
