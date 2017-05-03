package main

import (
	"github.com/owl/rp"
)

const (
	uuid      = "26169e7b-f3fc-46a8-8ef6-a41f8ee30fcd"
	project   = "WFR-API"
	reportDir = "report"
	apiURL    = "http://10.48.128.12:80/api/v1"
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
	launchStart := report.GetLaunchStartTime()
	launchEnd := report.GetLaunchEndTime()

	// create new Report Portal client
	rpClient := rp.NewClient(apiURL, project, uuid)

	// post new launch to rp
	launchID := rpClient.StartLaunch(&rp.Launch{
		Name:      "api test launch",
		StartTime: launchStart,
		Mode:      rp.ModeDebug,
		Tags:      []string{"R50", "debug", "go"},
	})

	// TODO:

	// update launch to completed state
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: launchEnd,
		Status:  rp.ExecutionStatusPassed, // launch always passed
	})

}
