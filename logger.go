package mpc

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	globalLogger *AppLogger
	loggerOnce   sync.Once
)

type AppLogger struct {
	mux  sync.Mutex
	zlog zerolog.Logger
	id   string
}

func SetupLogger(option *LoggerConfig) {
	NewAppLogger(option)
}

func NewAppLogger(options ...*LoggerConfig) *AppLogger {
	switch len(options) {
	case 0:
	default:
		loggerOnce.Do(func() {
			var w = ioutil.Discard
			fmt.Println(options[0].LoggerFileName())
			wd, _ := os.Getwd()
			fd, err := os.OpenFile(wd+options[0].LoggerFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				panic(err)
			}
			w = fd
			w = diode.NewWriter(w, 1000, 10*time.Millisecond, func(missed int) {
				os.Stderr.Write([]byte("[WARN] Logger dropped " + strconv.Itoa(missed) + " messages."))
			})
			log := zerolog.MultiLevelWriter(w, os.Stderr)
			globalLogger = &AppLogger{
				zlog: zerolog.New(log).Level(options[0].LoggerLevel()),
			}
		})
	}
	return &AppLogger{
		zlog: globalLogger.zlog.With().Timestamp().Logger(),
	}
}

func (a *AppLogger) RequestID() string {
	return a.id
}

func (a *AppLogger) New(id string) Logger {
	a.mux.Lock()
	defer a.mux.Unlock()

	return &AppLogger{
		zlog: globalLogger.zlog.With().Timestamp().Logger(),
		id:   id,
	}
}

func (a *AppLogger) Reuse(l Logger) {

}

func (a *AppLogger) caller() *zerolog.Logger {
	log := a.zlog.With().Caller().Str("request_id", a.id).Logger()
	return &log
}

func (a *AppLogger) Print(v ...interface{}) {
	a.caller().Log().Msgf("%+v", v...)
}

func (a *AppLogger) Printf(format string, v ...interface{}) {
	a.caller().Log().Msgf(format, v...)
}

func (a *AppLogger) Debug(v ...interface{}) {
	a.caller().Debug().Msgf("%+v", v...)
}

func (a *AppLogger) Debugf(format string, v ...interface{}) {
	a.caller().Debug().Msgf(format, v...)
}

func (a *AppLogger) Info(v ...interface{}) {
	a.caller().Info().Msgf("%+v", v...)
}

func (a *AppLogger) Infof(format string, v ...interface{}) {
	a.caller().Info().Msgf(format, v...)
}

func (a *AppLogger) Warn(v ...interface{}) {
	a.caller().Warn().Msgf("%+v", v...)
}

func (a *AppLogger) Warnf(format string, v ...interface{}) {
	a.caller().Warn().Msgf(format, v...)
}

func (a *AppLogger) Error(v ...interface{}) {
	a.caller().Error().Msgf("%+v", v...)
}

func (a *AppLogger) Errorf(format string, v ...interface{}) {
	a.caller().Error().Msgf(format, v...)
}

func (a *AppLogger) Fatal(v ...interface{}) {
	a.caller().Fatal().Msgf("%+v", v...)
}

func (a *AppLogger) Fatalf(format string, v ...interface{}) {
	a.caller().Fatal().Msgf(format, v...)
}
