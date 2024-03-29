package gofm

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MissEvan MissEvanConfig `yaml:"missevan"`
	OpenAi   OpenAIConfig   `yaml:"openai"`
}

type MissEvanConfig struct {
	Live  int    `yaml:"live"`
	Token string `yaml:"token"`
}

type OpenAIConfig struct {
	Key   string `yaml:"key"`
	Proxy string `yaml:"proxy"`
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
