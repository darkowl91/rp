package rp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

// StartLaunch create new launch for specified project
func (c *Client) StartLaunch(launch Launch) (launchID string) {
	apiURL := c.baseURL + "/launch"

	if len(launch.StartTime) == 0 {
		launch.StartTime = time.Now().Format(time.RFC3339)
	}

	payload, err := json.Marshal(launch)
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

	log.Info("RP Responce: %v", resp)
	if resp.StatusCode >= 400 {
		log.Error(decodeError(resp.Body))
	} else if resp.StatusCode == http.StatusCreated {
		var launchResponce struct {
			ID string `json:"id"`
		}
		err := json.NewDecoder(resp.Body).Decode(&launchResponce)
		if err != nil {
			log.Error(err)
		}
		launchID = launchResponce.ID
	}

	return
}

// FinishLaunch update specified lauch to passed (completed state)
func (c *Client) FinishLaunch(launchID, executionStatus string) {
	if len(launchID) == 0 {
		log.Error("launchID could not be empty")
	}
	apiURL := c.baseURL + "/launch/" + launchID + "/finish"

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
