package rp

import (
	"encoding/json"
	"net/http"
)

// StartLaunch creates new launch
func (c *Client) StartLaunch(launch Launch) (launchID *ResponceID) {
	resp, err := c.post("/launch", launch)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode == http.StatusCreated {
		err := json.NewDecoder(resp.Body).Decode(&launchID)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error(decodeError(resp.Body))
	}
	return
}

// FinishLaunch update specified launch to passed (completed state)
func (c *Client) FinishLaunch(launchID string, result ExecutionResult) {
	if len(launchID) == 0 {
		log.Error("launchID could not be empty")
		return
	}

	resp, err := c.put("/launch/"+launchID+"/finish", result)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Error(decodeError(resp.Body))
	}
}
