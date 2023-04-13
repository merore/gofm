package missevan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/merore/gofm/pkg/logger"
)

type Connection interface {
	Read() (FMMessage, error)
	Close() error
}

type connection struct {
	live int
	fms  chan FMMessage
	conn *websocket.Conn
}

func NewConnection(live int, conn *websocket.Conn) Connection {
	c := &connection{
		live: live,
		fms:  make(chan FMMessage, FMMessagesSize),
		conn: conn,
	}
	c.join()
	go c.heartbeat()
	go c.handle()
	return c
}

func (c *connection) Close() error {
	return nil
}

// Read read a FMMessage from connection.
// When error is not nil, the connection has been broken.
func (c *connection) Read() (FMMessage, error) {
	return <-c.fms, nil
}

func (c *Client) connect(live int) error {
	return nil
}

func (c *connection) join() error {
	msg := fmt.Sprintf(`{"action":"join","uuid":"%s","type":"room","room_id":%d}`, uuid.New().String(), c.live)
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		return err
	}
	return nil
}

func (c *connection) heartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		if err := c.conn.WriteMessage(websocket.TextMessage, []byte("❤️")); err != nil {
			logger.Error(err)
		}
	}
}

func (c *connection) handle() error {
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
			c.fms <- m
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
