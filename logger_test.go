package logger

import (
	"github.com/haderianous/go-logger/logger"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := logger.NewLogger(logger.DebugLevel, logger.JsonEncoding)
	data := struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Location string `json:"location"`
	}{Name: "ali", Age: 20, Location: "tehran"}
	log.DebugF("debug log")
	log.With(logger.Field{
		"test": data,
	}).InfoF("log with extra data")
	log.WarnF("warn error")
	log.ErrorF("test error log")
}
