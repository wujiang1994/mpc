package mpc

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type AppConfig struct {
	Name    string                  `yaml:"name"`
	Mode    RunMode                 `yaml:"mode"`
	Logger  *LoggerConfig           `yaml:"logger"`
	Section map[RunMode]interface{} `yaml:"section"`

	mux      sync.RWMutex
	filename string
}

type LoggerConfig struct {
	Output   string `yaml:"output"`
	Filename string `yaml:"filename"`
	Ext      string `yaml:"ext"`
	Level    string `yaml:"level"`
}

func NewAppConfig(runMode, cfgPath string) (config *AppConfig, err error) {
	config = new(AppConfig)
	cfgPath = filepath.Clean(cfgPath)
	root, err := os.Getwd()
	if err != nil {
		return config, err
	}
	filename := "application.yml"
	filename = filepath.Join(root, cfgPath, "config", runMode, filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err = yaml.Unmarshal(b, config); err != nil {
		return config, err
	}
	config.filename = filename
	return
}

func (f *AppConfig) RunMode() RunMode {
	return f.Mode
}

func (f *AppConfig) RunName() string {
	return f.filename
}

func (f *AppConfig) SetMode(mode RunMode) {
	f.mux.Lock()
	defer f.mux.Unlock()

	if !mode.IsValid() {
		return
	}
	f.Mode = mode
}

func (f *AppConfig) UnmarshalYaml(v interface{}) error {
	f.mux.Lock()
	defer f.mux.Unlock()

	section, ok := f.Section[f.Mode]
	if !ok {
		return fmt.Errorf("error config section")
	}
	b, err := yaml.Marshal(section)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, v)
}

func (l *LoggerConfig) LoggerFileName() string {
	return l.Output + l.Filename + l.Ext
}

func (l *LoggerConfig) LoggerLevel() zerolog.Level {
	switch l.Level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	}
	return zerolog.InfoLevel
}
