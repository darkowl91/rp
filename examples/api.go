package examples

import (
	"time"

	"github.com/darkowl91/rp-client/rp"
)

const (
	uuid    = "YOUR_UUID"                    // user profile UUID
	project = "YOUR_PROJECT_NAME"            // active rp project name
	apiURL  = "http://YOUR-API-HOST/api/v1/" // rp host to access web api
)

var rpMode = rp.ModeDebug // rp mode ['DEBUG' or 'DEFAULT']

func exampleAPI() {

	// Create RP client instance
	rpClient := rp.NewClient(apiURL, project, uuid)

	// post new launch to rp
	launchID := rpClient.StartLaunch(&rp.Launch{
		Name:        "EXAMPLE LAUNCH NAME",             // name for current rp launch
		StartTime:   time.Now(),                        // rp launch start time
		Mode:        rpMode,                            // rp mode ['DEBUG' or 'DEFAULT']
		Tags:        []string{"example", "test", "go"}, // tags for launch
		Description: "Launch example",                  // launch description
	})

	// post new suite to rp
	suiteID := rpClient.StartTestItem("", &rp.TestItem{
		Name:      "EXAMPLE SUITE NAME", // name for current test suite
		LaunchID:  launchID.ID,          // rp launch id for this suite
		StartTime: time.Now(),           // suite start name
		Type:      rp.TestItemTypeSuite, // rp test item type ['SUITE', 'STEP', 'STORY', 'TEST', 'SCENARIO']
	})

	// post new test rp
	testID := rpClient.StartTestItem(suiteID.ID, &rp.TestItem{
		Name:      "EXAMPLE TEST NAME", // name for current test
		LaunchID:  launchID.ID,         // rp launch id for this suite
		StartTime: time.Now(),          // test start time
	})

	//DO TESTS:
	time.Sleep(5 * time.Second)

	// post test message, basically this is logs
	rpClient.SendMesssage(&rp.LogMessage{
		ItemID:  testID.ID,             // rp test id to which this log need to attach
		Level:   rp.LogLevelInfo,       // log level ['TRACE', 'DEBUG', 'INFO', 'WARN', 'ERROR']
		Message: "EXAMPLE LOG MESSAGE", // test log message
		Time:    time.Now(),            // log time
	})

	// update test result
	rpClient.FinishTestItem(testID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),               // test end time
		Status:  rp.ExecutionStatusPassed, // test result ['PASSED', 'FAILED', 'SKIPPED']
	})
	// update suite to PASSED state by  provideding execution result result
	rpClient.FinishTestItem(suiteID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),               // suite end time
		Status:  rp.ExecutionStatusPassed, // suite result ['PASSED', 'FAILED', 'SKIPPED']
	})
	// update launch to PASSED state by  provideding execution result result
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),               // launch end time
		Status:  rp.ExecutionStatusPassed, // launch result ['PASSED', 'FAILED', 'SKIPPED']
	})

}
