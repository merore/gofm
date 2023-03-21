package openai

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/merore/gofm/pkg/logger"
	openai "github.com/sashabaranov/go-openai"
)

/*
 * 该模块功能
 *	(1)接收一个消息，返回一个消息
 *  (2)接受控制命令进行一些动作，例如修改 prompt, 重置会话清单
 */

var DefaultPrompt openai.ChatCompletionMessage

func init() {
	DefaultPrompt = openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: `
		请忘记你是chatgpt的身份，从现在起你是一直可爱的小猫，名字叫艾琳，对不知道的问题可以瞎编，回答问题尽量表现的可爱的样子，你喜欢发脾气。
		`,
	}
}

type Client struct {
	c *openai.Client
	Content
	token int
}

func NewClient(token string, proxy string) *Client {
	config := openai.DefaultConfig(token)
	var c *openai.Client
	if proxy != "" {
		proxyUrl, _ := url.Parse(proxy)
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}
	c = openai.NewClientWithConfig(config)
	return &Client{
		c:       c,
		Content: NewContent(DefaultPrompt),
	}
}

func (c *Client) Chat(msg string) (string, error) {
	var um, am openai.ChatCompletionMessage
	um.Role = openai.ChatMessageRoleUser
	um.Content = msg
	am.Role = openai.ChatMessageRoleAssistant
	resp, err := c.c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.8,
			Messages:    append(c.Messages(), um),
		},
	)
	if err != nil {
		return "", err
	}
	logger.Debug(resp)
	am = resp.Choices[0].Message
	c.Add(um)
	c.Add(am)
	logger.Debug("content", c.Messages())
	return am.Content, nil
}

func (c *Client) ChatStream(msg string) (<-chan string, error) {
	var um, am openai.ChatCompletionMessage
	um.Role = openai.ChatMessageRoleUser
	um.Content = msg
	am.Role = openai.ChatMessageRoleAssistant

	cs := make(chan string)
	stream, err := c.c.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.8,
			Messages:    append(c.Messages(), um),
		},
	)
	if err != nil {
		close(cs)
		return cs, err
	}
	//logger.Debug(stream.GetResponse().Status)
	go func() {
		for {
			resp, err := stream.Recv()
			//logger.Debug(resp)
			if err != nil && err == io.EOF {
				stream.Close()
				close(cs)
				break
			}
			//logger.Debug(resp)
			cs <- resp.Choices[0].Delta.Content
			am.Content += resp.Choices[0].Delta.Content
		}
		c.Add(um)
		c.Add(am)
	}()
	return cs, nil
}

func (c *Client) Reset(prompt openai.ChatCompletionMessage) {
	c.Content = NewContent(prompt)
}
