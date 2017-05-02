package rp

import (
	"encoding/json"
	"net/http"
)

// StartTestItem is used to create new test suite for specified launch
func (c *Client) StartTestItem(parentItemID string, testItem TestItem) (testItemID *ResponceID) {
	apiURL := "/item"
	if len(parentItemID) > 0 {
		apiURL = apiURL + "/" + parentItemID
	}

	resp, err := c.post(apiURL, testItem)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode == http.StatusCreated {
		err := json.NewDecoder(resp.Body).Decode(&testItemID)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error(decodeError(resp.Body))
	}
	return
}

// FinishTestItem update specified test item to passed (completed state)
func (c *Client) FinishTestItem(testItemID string, result ExecutionResult) {
	if len(testItemID) == 0 {
		log.Error("suiteID could not be empty")
		return
	}

	resp, err := c.put("/item/"+testItemID, result)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Error(decodeError(resp.Body))
	}
}

// SendMesssage create new log entry for provided item
func (c *Client) SendMesssage(lgoMessage LogMessage) (messageID *ResponceID) {
	resp, err := c.post("/log", lgoMessage)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode == http.StatusCreated {
		err := json.NewDecoder(resp.Body).Decode(&messageID)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error(decodeError(resp.Body))
	}
	return
}
