package rp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	apiURL = "http://10.48.128.12:80/api/v1/"

	// TimestampLayout can be used with time.Parse to create time.Time values from strings.
	// It is an ISO 8601 UTC timestamp with a zero offset.
	TimestampLayout = "2006-01-02T15:04:05"

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
		log.Error("project could not be empty")
	}
	if len(uuid) == 0 {
		log.Error("uuid could not be empty")
	}
	return Client{
		baseURL:    apiURL + project,
		authBearer: "Bearer " + uuid,
		http:       new(http.Client),
	}
}

// createNewRequest used for building new http.Request to RP API with default headers
// apiUrl should start from "/" e.g. '/launch'
func (c *Client) createNewRequest(method string, apiURL string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+apiURL, bytes.NewBuffer(payload))
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
	log.Debugf("RP Request: %v", req)
	resp, err := c.http.Do(req)
	log.Debugf("RP Responce: %v", resp)
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
