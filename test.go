package rp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

// StartTestItem is used to create new test suite for specified launch
func (c *Client) StartTestItem(parentItemID string, testItem TestItem) (testItemID string) {
	apiURL := c.baseURL + "/item"
	if len(parentItemID) > 0 {
		apiURL = apiURL + "/" + parentItemID
	}

	if len(testItem.StartTime) == 0 {
		testItem.StartTime = time.Now().Format(time.RFC3339)
	}

	payload, err := json.Marshal(testItem)
	if err != nil {
		log.Error(err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", c.authBearer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	log.Infof("RP Request: %v", req)

	resp, err := c.http.Do(req)
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	log.Infof("RP Responce: %v", resp)
	if resp.StatusCode >= 400 {
		log.Error(decodeError(resp.Body))
	} else if resp.StatusCode == http.StatusCreated {
		var testItemResponce struct {
			ID string `json:"id"`
		}
		err := json.NewDecoder(resp.Body).Decode(&testItemResponce)
		if err != nil {
			log.Error(err)
		}
		testItemID = testItemResponce.ID
	}

	return
}

// FinishTestItem update specified test item to passed (completed state)
func (c *Client) FinishTestItem(testItemID, executionStatus string) {
	if len(testItemID) == 0 {
		log.Error("suiteID could not be empty")
	}
	apiURL := c.baseURL + "/item/" + testItemID

	result := new(executionResult)
	result.EndTime = time.Now().Format(time.RFC3339)
	result.Status = executionStatus

	payload, err := json.Marshal(result)
	if err != nil {
		log.Error(err)
	}

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", c.authBearer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	log.Infof("RP Request: %v", req)

	resp, err := c.http.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	log.Infof("RP Responce: %v", resp)
	if resp.StatusCode >= 400 {
		log.Error(decodeError(resp.Body))
	}
}

// SendMesssage create new log entry for provided item
func (c *Client) SendMesssage(lgoMessage LogMessage) {
	apiURL := c.baseURL + "/log"

	payload, err := json.Marshal(lgoMessage)
	if err != nil {
		log.Error(err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", c.authBearer)
	req.Header.Add("Content-Type", "application/json")

	log.Infof("RP Request: %v", req)

	resp, err := c.http.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	log.Infof("RP Responce: %v", resp)
	if resp.StatusCode >= http.StatusBadRequest {
		log.Error(decodeError(resp.Body))
	}

}
