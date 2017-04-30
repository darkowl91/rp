package rp

import (
	"log"
	"net/http"
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

// NewClient creates a RP Client for specified project and user unique id
func NewClient(project, uuid string) Client {
	if len(project) == 0 {
		log.Fatal("project could not be empty")
	}
	if len(uuid) == 0 {
		log.Fatal("uuid could not be empty")
	}
	return Client{
		baseURL:    APIURL + project,
		authBearer: "Bearer " + uuid,
		http:       new(http.Client),
	}
}
