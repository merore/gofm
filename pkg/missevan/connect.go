package missevan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/merore/gofm/pkg/logger"
)

func (c *Client) connect(live int) error {
	if err := c.dial(live); err != nil {
		return err
	}
	if err := c.join(live); err != nil {
		return err
	}
	go c.heartbeat()
	return c.handle()
}

func (c *Client) disconnect() error {
	return c.conn.Close()
}
func (c *Client) retryConnect(live int) error {
	for {
		if err := c.connect(live); err != nil {
			c.disconnect()
			logger.Error(err, "reconnect to "+strconv.Itoa(live))
		}
	}
}

func (c *Client) dial(live int) error {
	req, _ := c.NewRequest(http.MethodConnect, WebsocketUrl, nil)
	ws := &websocket.Dialer{}
	conn, resp, _ := ws.Dial(fmt.Sprintf("%s?room_id=%d", WebsocketUrl, live), req.Header)
	if resp.StatusCode != 101 {
		defer conn.Close()
		bs, _ := ioutil.ReadAll(resp.Body)
		return errors.New("Dial() " + string(bs))
	}
	c.conn = conn
	return nil
}

func (c *Client) join(live int) error {
	msg := fmt.Sprintf(`{"action":"join","uuid":"%s","type":"room","room_id":%d}`, uuid.New().String(), live)
	return c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (c *Client) heartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		if err := c.conn.WriteMessage(websocket.TextMessage, []byte("❤️")); err != nil {
			logger.Error(err)
		}
	}
}

func (c *Client) handle() error {
	for {
		t, bs, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}
		if len(bs) == 0 || t == websocket.TextMessage {
			continue
		}
		if t == websocket.BinaryMessage {
			bs, _ = BrotliDecompress(bs)
		}
		ms := parseFMMessage(bs)
		for _, m := range ms {
			c.cms <- m
		}
	}
}

func parseFMMessage(msgData []byte) []FMMessage {
	// Make sure the JSON object is an array.
	buf := bytes.Buffer{}
	if msgData[0] != '[' {
		buf.WriteByte('[')
		buf.Write(msgData)
		buf.WriteByte(']')
	} else {
		buf.Write(msgData)
	}
	var textMsgs []FMMessage
	_ = json.Unmarshal(buf.Bytes(), &textMsgs)
	return textMsgs
}
