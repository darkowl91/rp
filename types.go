package rp

import (
	"encoding/json"
	"net/http"
	"time"
)

// Client is a client for working with the RP Web API.
type Client struct {
	baseURL    string
	authBearer string
	http       *http.Client
}

// Launch that identifies a test run.
type Launch struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Mode        string    `json:"mode,omitempty"`
	StartTime   time.Time `json:"start_time"`
	Tags        []string  `json:"tags,omitempty"`
}

// MarshalJSON with custom time format
func (launch *Launch) MarshalJSON() ([]byte, error) {
	type Alias Launch
	return json.Marshal(&struct {
		StartTime string `json:"start_time"`
		*Alias
	}{
		StartTime: launch.StartTime.Format(TimestampLayout),
		Alias:     (*Alias)(launch),
	})
}

// TestItem identifies a test suite, test, test method (step) fot test run.
type TestItem struct {
	LaunchID    string    `json:"launch_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags,omitempty"`
}

// MarshalJSON with custom time format
func (item *TestItem) MarshalJSON() ([]byte, error) {
	type Alias TestItem
	return json.Marshal(&struct {
		*Alias
		StartTime string `json:"start_time"`
	}{
		Alias:     (*Alias)(item),
		StartTime: item.StartTime.Format(TimestampLayout),
	})
}

// ExecutionResult is used to update executed TestItem.
type ExecutionResult struct {
	EndTime time.Time `json:"end_time"`
	Status  string    `json:"status"`
}

// MarshalJSON with custom time format
func (result *ExecutionResult) MarshalJSON() ([]byte, error) {
	type Alias ExecutionResult
	return json.Marshal(&struct {
		*Alias
		EndTime string `json:"end_time"`
	}{
		Alias:   (*Alias)(result),
		EndTime: result.EndTime.Format(TimestampLayout),
	})
}

// ResponceID of created item
type ResponceID struct {
	ID string `json:"id"`
}

// LogMessage identifies test log.
type LogMessage struct {
	ItemID  string    `json:"item_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
	Level   string    `json:"level"`
}

// MarshalJSON with custom time format
func (msg *LogMessage) MarshalJSON() ([]byte, error) {
	type Alias LogMessage
	return json.Marshal(&struct {
		*Alias
		Time string `json:"time"`
	}{
		Alias: (*Alias)(msg),
		Time:  msg.Time.Format(TimestampLayout),
	})
}
