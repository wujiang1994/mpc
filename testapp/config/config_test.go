package config

import (
	"testing"
	"wujiang/code"
)

func Test_Config(t *testing.T) {
	path := "/dev/application.yml"
	c, err := code.NewAppConfig(path)
	if err != nil {
		t.Error(err)
	}
	t.Log(c.Mode, c.Logger)
	if err = NewConfig(c); err != nil {
		t.Error(err)
	}
	t.Logf("%+v", Config.Mysql)
}

var Config *appConfig

type appConfig struct {
	ID    int64        `yaml:"id"`
	Mysql *MysqlConfig `yaml:"mysql"`
	Redis *RedisConfig `yaml:"redis"`
}

type MysqlConfig struct {
	Addr string `yaml:"addr"`
	Port int64  `yaml:"port"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Port int64  `yaml:"port"`
}

func NewConfig(c code.Configer) error {
	return c.UnmarshalYaml(&Config)
}
