package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/merore/gofm/pkg/logger"
)

const (
	ModelGPT3DOT5 = "gpt-3.5-turbo"

	RoleUser = "user"
	RoleSys  = "system"
	RoleAsst = "assistant"

	UrlChat = "https://api.openai.com/v1/chat/completions"
)

var (
	DefaultPrompt = ChatMessage{
		Role: RoleSys,
		Content: `
		请忘记你是chatgpt的身份，从现在起你是一直可爱的小猫，名字叫艾琳，对不知道的问题可以瞎编，回答问题尽量表现的可爱的样子，你喜欢发脾气。
		`,
	}
)

type Client struct {
	c     *http.Client
	token string
	Content
}

func NewClient(token string) (*Client, error) {
	c := &Client{
		c:       &http.Client{},
		Content: NewContent(DefaultPrompt),
		token:   token,
	}
	return c, nil
}

func (c *Client) Chat(msg string) (string, error) {
	umsg := ChatMessage{Role: RoleUser, Content: msg}
	req := ChatRequest{
		Model:    ModelGPT3DOT5,
		Messages: append(c.Messages(), umsg),
	}
	resp, err := c.do(req)
	if err != nil {
		return "", err
	}
	amsg := resp.Choices[0].Message
	defer func() {
		c.Add(umsg)
		c.Add(amsg)
	}()
	return amsg.Content, nil
}

func (c *Client) do(chatRequest ChatRequest) (ChatResponse, error) {
	var chatResponse ChatResponse
	bs, _ := json.Marshal(chatRequest)
	req := c.newRequest(http.MethodPost, UrlChat, bytes.NewReader(bs))
	resp, err := c.c.Do(req)
	if err != nil {
		return chatResponse, err
	}
	defer resp.Body.Close()
	bs, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, &chatResponse)
	logger.Info(string(bs))
	return chatResponse, nil

}

func (c *Client) newRequest(method string, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.token)
	return req
}

func (c *Client) Reset() {
	c.Content = NewContent(DefaultPrompt)
}
