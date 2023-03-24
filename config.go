package gofm

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MissevanToken string
	MissevanLive  int
	OpenAIToken   string
	OpenAIAPI     string
}

func UnmarshalConfig(filename string) (Config, error) {
	c := Config{}
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(bs, &c)
	return c, err
}

func MarshalConfig(config Config) string {
	bs, _ := yaml.Marshal(config)
	return string(bs)
}
