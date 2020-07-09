package config

import (
	"mpc"
	"testing"
)

func Test_Config(t *testing.T) {
	runMode := "dev"
	path := "/dev/application.yml"
	c, err := mpc.NewAppConfig(runMode, path)
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

func NewConfig(c mpc.Configer) error {
	return c.UnmarshalYaml(&Config)
}
