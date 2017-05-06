package rp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TestItemType string
type ExecutionStatus string
type LogLevel string
type Mode string

const (
	// TimestampLayout can be used with time.Parse to create time.Time values from strings.
	TimestampLayout = "2006-01-02T15:04:05.000-07:00"

	// TestItemTypeSuite - SUITE
	TestItemTypeSuite TestItemType = "SUITE"
	// TestItemTypeStep - STEP
	TestItemTypeStep TestItemType = "STEP"
	// TestItemTypeStory - STORY
	TestItemTypeStory TestItemType = "STORY"
	// TestItemTypeTest - TEST
	TestItemTypeTest TestItemType = "TEST"
	// TestItemTypeScenario - SCENARIO
	TestItemTypeScenario TestItemType = "SCENARIO"

	// ExecutionStatusPassed - PASSED
	ExecutionStatusPassed ExecutionStatus = "PASSED"
	// ExecutionStatusFailed - FAILED
	ExecutionStatusFailed ExecutionStatus = "FAILED"
	// ExecutionStatusSkipped - SKIPPED
	ExecutionStatusSkipped ExecutionStatus = "SKIPPED"

	// LogLevelTrace - TRACE
	LogLevelTrace LogLevel = "TRACE"
	// LogLevelDebug - DEBUG
	LogLevelDebug LogLevel = "DEBUG"
	// LogLevelInfo - INFO
	LogLevelInfo LogLevel = "INFO"
	// LogLevelWarn - WARN
	LogLevelWarn LogLevel = "WARN"
	// LogLevelError - ERROR
	LogLevelError LogLevel = "ERROR"

	// ModeDebug - DEBUG
	ModeDebug Mode = "DEBUG"
	// ModeDefault - DEFAULT
	ModeDefault Mode = "DEFAULT"
)

// NewClient creates a RP Client for specified project and user unique id
func NewClient(apiURL, project, uuid string) Client {
	if len(project) == 0 {
		log.Error("project could not be empty")
	}
	if len(uuid) == 0 {
		log.Error("uuid could not be empty")
	}
	return Client{
		baseURL:    joinURL(apiURL, project),
		authBearer: "Bearer " + uuid,
		http:       new(http.Client),
	}
}

// createNewRequest is used for building new http.Request to RP API with default headers
// apiUrl should start from "/" e.g. '/launch'
func (c *Client) createNewRequest(method string, apiURL string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, joinURL(c.baseURL, apiURL), bytes.NewBuffer(payload))
	req.Header.Add("Authorization", c.authBearer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	return req, err
}

// request is used to send api request to rp
func (c *Client) request(method, apiURL string, payload []byte) (*http.Response, error) {
	req, err := c.createNewRequest(method, apiURL, payload)
	if err != nil {
		return nil, err
	}
	log.Debugf("rp request: %v", req)
	resp, err := c.http.Do(req)
	log.Debugf("rp responce: %v", resp)
	return resp, err
}

// post request
func (c *Client) post(apiURL string, body interface{}) (*http.Response, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return c.request("POST", apiURL, payload)
}

// put request
func (c *Client) put(apiURL string, body interface{}) (*http.Response, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return c.request("PUT", apiURL, payload)
}
