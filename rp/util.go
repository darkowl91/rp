package rp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"

	"time"

	logging "github.com/op/go-logging"
)

var (
	log    = logging.MustGetLogger("rp.logger")
	format = logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s}%{color:reset} %{message}")
)

// InitLogger for structured logger output
func InitLogger() {
	logHandler := logging.NewLogBackend(os.Stdout, "", 0)

	formatter := logging.NewBackendFormatter(logHandler, format)

	logger := logging.AddModuleLevel(logHandler)
	logger.SetLevel(logging.ERROR, "")
	logger.SetLevel(logging.ERROR, "rp.logger")
	logging.SetBackend(logger, formatter)
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

// joinURL join URL parts as path segments
func joinURL(base string, parts ...string) string {
	u, err := url.Parse(base)
	if err != nil {
		log.Errorf("could not parse base '%s' url: %v", base, err)
		return base
	}
	u.Path = path.Join(u.Path, path.Join(parts...))
	return u.String()
}

// parseTimeStamp parsing with TimestampLayout
func parseTimeStamp(timeStr string) time.Time {
	t, err := time.Parse(xmlTimestampLayout, timeStr)
	if err != nil {
		log.Error(err)
	}
	return t
}

// converts seconds to duration
func secondsToDuration(sec float64) time.Duration {
	return time.Duration(int64(sec * float64(time.Second)))
}
