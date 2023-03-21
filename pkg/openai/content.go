package openai

import (
	"container/list"

	openai "github.com/sashabaranov/go-openai"
)

type Content interface {
	Add(openai.ChatCompletionMessage)
	Messages() []openai.ChatCompletionMessage
}

func NewContent(prompt openai.ChatCompletionMessage) Content {
	return newListContent(prompt)
}

type ListContent struct {
	prompt   openai.ChatCompletionMessage
	messages *list.List
	token    int
}

func newListContent(promot openai.ChatCompletionMessage) *ListContent {
	return &ListContent{
		prompt:   promot,
		messages: list.New(),
		token:    len(promot.Content),
	}
}

func (lc *ListContent) Add(message openai.ChatCompletionMessage) {
	lc.messages.PushBack(message)
	lc.token += len(message.Content)

	for lc.overflow() {
		lc.remove()
	}
}

func (lc *ListContent) Messages() []openai.ChatCompletionMessage {
	ms := make([]openai.ChatCompletionMessage, lc.messages.Len()+1)
	ms[0] = lc.prompt
	i := 1
	for e := lc.messages.Front(); e != nil; e = e.Next() {
		ms[i] = e.Value.(openai.ChatCompletionMessage)
		i++
	}
	return ms
}

func (lc *ListContent) overflow() bool {
	if lc.token*3/4+1 > 4096 {
		return true
	}
	return false
}

func (lc *ListContent) remove() {
	e := lc.messages.Front()
	lc.messages.Remove(e)
	lc.token -= len(e.Value.(openai.ChatCompletionMessage).Content)
}
