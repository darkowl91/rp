package rp

import "net/http"

// Client is a client for working with the RP Web API.
type Client struct {
	baseURL    string
	authBearer string
	http       *http.Client
}

// Launch that identifies a test run.
type Launch struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Mode        string   `json:"mode,omitempty"`
	StartTime   string   `json:"start_time"`
	Tags        []string `json:"tags,omitempty"`
}

// TestItem identifies a test suite, test, test method (step) fot test run.
type TestItem struct {
	LaunchID    string   `json:"launch_id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	StartTime   string   `json:"start_time"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags,omitempty"`
}

// executionRresult is used to update executed TestItem.
type executionRresult struct {
	EndTime string `json:"end_time"`
	Status  string `json:"status"`
}

// LogMessage identifies test log.
type LogMessage struct {
	ItemID  string `json:"item_id"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Level   string `json:"level"`
}
