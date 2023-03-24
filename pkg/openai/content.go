package openai

import (
	"container/list"
)

type Content interface {
	Add(ChatMessage)
	Messages() []ChatMessage
}

func NewContent(prompt ChatMessage) Content {
	return newListContent(prompt)
}

type ListContent struct {
	prompt   ChatMessage
	messages *list.List
	token    int
}

func newListContent(promot ChatMessage) *ListContent {
	return &ListContent{
		prompt:   promot,
		messages: list.New(),
		token:    len(promot.Content),
	}
}

func (lc *ListContent) Add(message ChatMessage) {
	lc.messages.PushBack(message)
	lc.token += len(message.Content)

	for lc.overflow() {
		lc.remove()
	}
}

func (lc *ListContent) Messages() []ChatMessage {
	ms := make([]ChatMessage, lc.messages.Len()+1)
	ms[0] = lc.prompt
	i := 1
	for e := lc.messages.Front(); e != nil; e = e.Next() {
		ms[i] = e.Value.(ChatMessage)
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
	lc.token -= len(e.Value.(ChatMessage).Content)
}
