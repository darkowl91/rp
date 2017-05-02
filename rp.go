package rp

import (
	"net/http"
	"os"

	logging "github.com/op/go-logging"
)

const (
	// APIURL to access RP API
	APIURL = "http://10.48.128.12:80/api/v1/"

	// TestItemTypeSuite - SUITE
	TestItemTypeSuite = "SUITE"
	// TestItemTypeStep - STEP
	TestItemTypeStep = "STEP"
	// TestItemTypeStory - STORY
	TestItemTypeStory = "STORY"
	// TestItemTypeTest - TEST
	TestItemTypeTest = "TEST"
	// TestItemTypeScenario - SCENARIO
	TestItemTypeScenario = "SCENARIO"

	// ExecutionStatusPassed - PASSED
	ExecutionStatusPassed = "PASSED"
	// ExecutionStatusFailed - FAILED
	ExecutionStatusFailed = "FAILED"
	// ExecutionStatusSkipped - SKIPPED
	ExecutionStatusSkipped = "SKIPPED"

	// LogLevelTrace - TRACE
	LogLevelTrace = "TRACE"
	// LogLevelDebug - DEBUG
	LogLevelDebug = "DEBUG"
	// LogLevelInfo - INFO
	LogLevelInfo = "INFO"
	// LogLevelWarn - WARN
	LogLevelWarn = "WARN"
	// LogLevelError - ERROR
	LogLevelError = "ERROR"

	// ModeDebug - DEBUG
	ModeDebug = "DEBUG"
	// ModeDefault - DEFAULT
	ModeDefault = "DEFAULT"
)

// setup logger
var log = logging.MustGetLogger("rp.logger")

// setup logger format
var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

//InitLogger - initiate logger
func InitLogger() {
	logHandler := logging.NewLogBackend(os.Stderr, "rp ", 0)
	formatter := logging.NewBackendFormatter(logHandler, logFormat)
	logging.SetBackend(logHandler, formatter)
}

// NewClient creates a RP Client for specified project and user unique id
func NewClient(project, uuid string) Client {
	if len(project) == 0 {
		log.Error("project could not be empty")
	}
	if len(uuid) == 0 {
		log.Error("uuid could not be empty")
	}
	return Client{
		baseURL:    APIURL + project,
		authBearer: "Bearer " + uuid,
		http:       new(http.Client),
	}
}
