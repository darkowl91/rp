package examples

import (
	"github.com/darkowl91/rp-client/rp"
)

var reportDir = "report"

func examplePublishReport() {
	// enable logging
	rp.InitLogger()

	// load xml results from folder
	report, err := rp.LoadXMLReport(reportDir)
	if err != nil {
		panic(err)
	}

	// start, end for launch
	launchStart := report.LaunchStartTime()
	launchEnd := report.LaunchEndTime()

	// create new Report Portal client
	rpClient := rp.NewClient(apiURL, project, uuid)

	// post new launch to rp
	launchID := rpClient.StartLaunch(&rp.Launch{
		Name:      "EXAMPLE PUBLISH LAUNCH NAME",
		StartTime: launchStart,
		Mode:      rp.ModeDebug,
		Tags:      []string{"example", "test", "go"},
	})

	// start post report
	for i := 0; i < report.SuitesCount(); i++ {
		// create suite
		suite := report.Suite(i)
		suite.LaunchID = launchID.ID
		suiteID := rpClient.StartTestItem("", suite)

		// start post cases
		for j := 0; j < report.TesCaseCount(i); j++ {
			//create test case
			tCase := report.TestCase(i, j)
			tCase.LaunchID = launchID.ID
			caseID := rpClient.StartTestItem(suiteID.ID, tCase)

			// post logs when have failure
			if report.HasTestCasefailure(i, j) {
				// post message
				failure := report.TestCasefailure(i, j)
				failure.ItemID = caseID.ID
				rpClient.SendMesssage(failure)
			}
			tResult := report.TestCaseResult(i, j)
			rpClient.FinishTestItem(caseID.ID, tResult)
		}
		suiteResult := report.SuiteResult(i)
		rpClient.FinishTestItem(suiteID.ID, suiteResult)
	}

	// update launch to completed state
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: launchEnd,
	})

}
