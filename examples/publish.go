package main

import (
	"github.com/owl/rp"
)

const (
	uuid      = "26169e7b-f3fc-46a8-8ef6-a41f8ee30fcd"
	project   = "WFR-API"
	reportDir = "report"
	apiURL    = "http://10.48.128.12:80/api/v1/"
)

func main() {
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
		Name:      "api test launch",
		StartTime: launchStart,
		Mode:      rp.ModeDebug,
		Tags:      []string{"R50", "debug", "go", "mw01"},
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

			// pos log in case we have failure
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
