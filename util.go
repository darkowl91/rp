package rp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	logging "github.com/op/go-logging"
)

var (
	log    = logging.MustGetLogger("rp.logger")
	format = logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s}%{color:reset} %{message}")
)

// InitLogger for structured logger output
func InitLogger(level logging.Level) {
	logging.Reset()
	// stdout
	logHandler := logging.NewLogBackend(os.Stdout, "", 0)
	formatter := logging.NewBackendFormatter(logHandler, format)
	leveledHandler := logging.AddModuleLevel(logHandler)
	leveledHandler.SetLevel(level, "rp.logger")
	logging.SetBackend(leveledHandler, formatter)
}

// decodeError decodes an Error from an io.Reader.
func decodeError(r io.Reader) error {
	var e struct {
		Code    int    `json:"error_code, omitempty"`
		Message string `json:"message"`
	}
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return errors.New("couldn't decode responce error")
	}
	if e.Code == 0 {
		return errors.New("no responce error")
	}
	return fmt.Errorf("code: %d, msg: %s", e.Code, e.Message)
}
