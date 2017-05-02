package main

import "github.com/owl/rp"
import "time"

const (
	uuid    = ""
	project = "default_project"
)

func main() {

	// Create new Report Portal client
	rpClient := rp.NewClient(project, uuid)

	// Create new launch before start test execution
	launchID := rpClient.StartLaunch(rp.Launch{
		Name: "Go Tests Launch",
		Mode: rp.ModeDebug,
		Tags: []string{"R50", "debug", "go"},
	})

	// Create new test suite for test launch
	suiteID := rpClient.StartTestItem("", rp.TestItem{
		LaunchID: launchID,
		Name:     "Go Test Suite",
		Type:     rp.TestItemTypeSuite,
	})

	// Create new test for given test suite
	testID := rpClient.StartTestItem(suiteID, rp.TestItem{
		LaunchID: launchID,
		Name:     "Go Test",
		Type:     rp.TestItemTypeTest,
	})

	// Create new test step for given test
	stepID := rpClient.StartTestItem(testID, rp.TestItem{
		LaunchID: launchID,
		Name:     "Go Test step (method)",
		Type:     rp.TestItemTypeStep,
	})

	time.Sleep(time.Second * 5) // Do the work

	// Add step log info
	rpClient.SendMesssage(rp.LogMessage{
		ItemID:  stepID,
		Level:   rp.LogLevelDebug,
		Message: "Go Debug log message",
		Time:    time.Now().Format(time.RFC3339),
	})

	// Add one more log msg
	rpClient.SendMesssage(rp.LogMessage{
		ItemID:  stepID,
		Level:   rp.LogLevelInfo,
		Message: "Go Info log message",
		Time:    time.Now().Format(time.RFC3339),
	})

	// update test step (method) to completed state
	rpClient.FinishTestItem(stepID, rp.ExecutionStatusPassed)
	// update test to completed state
	rpClient.FinishTestItem(testID, rp.ExecutionStatusPassed)
	// update test suite to completed state
	rpClient.FinishTestItem(suiteID, rp.ExecutionStatusPassed)

	// update launch to completed state
	rpClient.FinishLaunch(launchID, rp.ExecutionStatusPassed)
}
