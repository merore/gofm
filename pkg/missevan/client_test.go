package missevan

import (
	"encoding/json"
	"fmt"
	"testing"
)

var c *Client

const (
	TestLive  = 868858631
	TestToken = "6409cf69751eb1b91deaccbd|f528befe080f5994|1678364521|b0acb3c2a2e78743"
)

func init() {
	c = NewClient(TestToken)
}
func TestClientSend(t *testing.T) {
	type args struct {
		live int
		msg  string
	}
	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "test01",
			args: args{
				live: TestLive,
				msg:  "你好",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Send(tt.args.live, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientConnect(t *testing.T) {
	msgs := c.Connect(TestLive)
	for msg := range msgs {
		bs, _ := json.Marshal(msg)
		fmt.Printf("%v\n", string(bs))
	}
}

func TestClientGetUserInfo(t *testing.T) {
	fmUser, err := c.GetUserInfo()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", fmUser)
}
