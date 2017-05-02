package main

import (
	"time"

	logging "github.com/op/go-logging"
	"github.com/owl/rp"
)

const (
	uuid    = "26169e7b-f3fc-46a8-8ef6-a41f8ee30fcd"
	project = "default_project"
)

func main() {
	// enable log level
	rp.InitLogger(logging.DEBUG)

	// Create new Report Portal client
	rpClient := rp.NewClient(project, uuid)

	// Create new launch before start test execution
	launchID := rpClient.StartLaunch(&rp.Launch{
		Name:      "Go Tests Launch",
		Mode:      rp.ModeDebug,
		StartTime: time.Now(),
		Tags:      []string{"R50", "debug", "go"},
	})

	// update launch to completed state
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: time.Now(),
		Status:  rp.ExecutionStatusPassed,
	})
}
