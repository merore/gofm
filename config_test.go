package gofm

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var config = Config{
	MissEvanConfig{
		Live:  868854559,
		Token: "token|token",
	},
	OpenAIConfig{
		Key:   "key-key",
		Proxy: "http://localhost:7890",
	},
}

var rawconfig = `missevan:
    live: 868854559
    token: token|token
openai:
    key: key-key
    proxy: http://localhost:7890
`

func TestUnmarshalConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "test01",
			args: args{
				filename: "config-template.yaml",
			},
			want: config,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalConfig(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{
				config: config,
			},
			want: rawconfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarshalConfig(tt.args.config)
			if got != tt.want {
				t.Errorf("MarhalConfig() \ngot\n%v\nwant\n%v", got, tt.want)
				gs := hex.EncodeToString([]byte(got))
				ws := hex.EncodeToString([]byte(tt.want))
				t.Errorf("hex different\ngot\n%v\nwant\n%v", gs, ws)
			}
		})
	}
}
