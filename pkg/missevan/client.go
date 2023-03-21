package missevan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/merore/gofm/pkg/logger"
)

/*
 * 与 missevan 交互，功能如下
 * (1) 向 live 发送消息
 * (2) 获取某一 live 的信息流
 */

const (
	FM_SESS int = iota
	FM_SESS_SIG
	TOKEN
)
const (
	UserInfoUrl    = "https://fm.missevan.com/api/user/info"
	LoginUrl       = "https://app.missevan.com/member/login"
	MessageSendUrl = "https://fm.missevan.com/api/chatroom/message/send"
	WebsocketUrl   = "wss://im.missevan.com/ws"
	LiveOnlineUrl  = "https://fm.missevan.com/api/v2/chatroom/online"
)

const (
	MessagesSize = 100
	CookieSize   = 3
)

type Client struct {
	c       *http.Client
	cookies []*http.Cookie
	conn    *websocket.Conn
	cms     chan FMMessage
}

func NewClient(token string) *Client {
	c := &Client{
		c:       &http.Client{},
		cookies: make([]*http.Cookie, CookieSize),
		cms:     make(chan FMMessage, MessagesSize),
	}
	// store cookies
	baseCookies := getBaseCookies()
	c.cookies[FM_SESS] = baseCookies[FM_SESS]
	c.cookies[FM_SESS_SIG] = baseCookies[FM_SESS_SIG]
	c.cookies[TOKEN] = &http.Cookie{
		Name:  "token",
		Value: token,
	}
	return c
}

func (c *Client) Send(live int, msg string) error {
	type message struct {
		RoomID    int    `json:"room_id"`
		Message   string `json:"message"`
		MessageID string `json:"msg_id"`
	}
	rb, _ := json.Marshal(message{
		RoomID:    live,
		Message:   SafeMessage(msg),
		MessageID: MessageID(),
	})
	req, _ := c.NewRequest(http.MethodPost, MessageSendUrl, bytes.NewReader(rb))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Context-Length", strconv.Itoa(len(rb)))
	resp, _ := c.c.Do(req)

	var fmResp FMResp
	bs, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, &fmResp)
	if fmResp.Code != 0 {
		return errors.New(string(bs))
	}
	return nil
}

func (c *Client) Connect(live int) <-chan FMMessage {
	if err := c.online(live); err != nil {
		logger.Error(err)
	}
	go c.retryConnect(live)
	return c.cms
}

func (c *Client) GetUserInfo() (FMUser, error) {
	var fmUser FMUser
	req, _ := c.NewRequest(http.MethodGet, UserInfoUrl, nil)
	resp, err := c.c.Do(req)
	if err != nil {
		return fmUser, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	var fmResp FMResp
	bs, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bs, &fmResp)
	if fmResp.Info == nil || fmResp.Info.User == nil {
		return fmUser, errors.New("GetUserInfo() " + string(bs))
	}
	return *fmResp.Info.User, nil
}

func (c *Client) Online(live int) error {
	if err := c.online(live); err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			if err := c.online(live); err != nil {
				logger.Error(err)
			}
		}
	}()
	return nil
}

func (c *Client) online(live int) error {
	rb := fmt.Sprintf(`{"room_id":%d,"counter":1}`, live)
	req, _ := c.NewRequest(http.MethodPost, LiveOnlineUrl, strings.NewReader(rb))
	// must set [referer]
	req.Header.Add("referer", "https://fm.missevan.com/live/"+strconv.Itoa(live))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(len([]byte(rb))))
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	var fmResp FMResp
	bs, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bs, &fmResp); err != nil {
		return errors.New("online() " + err.Error() + string(bs))
	}
	if fmResp.Code != 0 {
		return errors.New(string(bs))
	}
	return nil
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.141 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,ja;q=0.5")
	req.AddCookie(c.cookies[FM_SESS])
	req.AddCookie(c.cookies[FM_SESS_SIG])
	req.AddCookie(c.cookies[TOKEN])
	return req, nil
}
