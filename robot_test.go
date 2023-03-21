package gofm

import "testing"

func TestRobotRun(t *testing.T) {
	config, _ := UnmarshalConfig("config.yaml")
	s := NewRobot(config)
	s.Run()
}

func TestWelcokeMessage(t *testing.T) {
	t.Logf("\n" + welcomeMessage("中国人"))
	t.Logf("\n" + welcomeMessage("中国人abdc"))
}
