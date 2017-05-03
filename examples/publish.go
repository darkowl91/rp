package main

import (
	"time"

	"github.com/owl/rp"
)

const (
	uuid    = "26169e7b-f3fc-46a8-8ef6-a41f8ee30fcd"
	project = "default_project"
)

func main() {
	// enable logging
	rp.InitLogger()

	// create new Report Portal client
	rpClient := rp.NewClient(project, uuid)

	// create new launch before start test execution
	launchID := rpClient.StartLaunch(&rp.Launch{
		Name:      "Go Tests Launch",
		Mode:      rp.ModeDebug,
		StartTime: time.Now(),
		Tags:      []string{"R50", "debug", "go"},
	})

	// create new test suite
	suiteID := rpClient.StartTestItem("", &rp.TestItem{
		LaunchID:    launchID.ID,
		Name:        "workflow_functional",
		StartTime:   time.Now(),
		Description: "companies.id.timesheets.id.action._alias.PUT",
		Type:        rp.TestItemTypeSuite,
	})

	// create new test
	stepID := rpClient.StartTestItem(suiteID.ID, &rp.TestItem{
		Name:        "Prepare",
		Description: "companies.id.timesheets.id.action._alias.PUT.workflow_functional",
		Type:        rp.TestItemTypeTest,
		StartTime:   time.Now(),
	})

	// send test log
	rpClient.SendMesssage(&rp.LogMessage{
		ItemID:  stepID.ID,
		Level:   rp.LogLevelError,
		Message: "Unexpected Status Code. Expected: 200, Actual: 403;",
		Time:    time.Now().Add(time.Duration(15)),
	})

	// update current test to failure
	rpClient.FinishTestItem(stepID.ID, &rp.ExecutionResult{
		Status:  rp.ExecutionStatusFailed,
		EndTime: time.Now(),
	})

	// update test suite to failure
	rpClient.FinishTestItem(suiteID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),
		Status:  rp.ExecutionStatusFailed,
	})

	// update launch to completed state
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),
		Status:  rp.ExecutionStatusFailed,
	})
}
