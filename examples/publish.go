package main

import (
	"github.com/owl/rp"
)

const (
	uuid      = "26169e7b-f3fc-46a8-8ef6-a41f8ee30fcd"
	project   = "default_project"
	reportDir = "C:/project/zeyt/test-api/report"
)

func main() {
	// enable logging
	rp.InitLogger()

	// create new Report Portal client
	rpClient := rp.NewClient(project, uuid)

	var launch = rp.Launch{
		Name: "Go test launch NoW",
		Mode: rp.ModeDebug,
		Tags: []string{"R50", "debug", "go"},
	}

	rp.PublishReport(reportDir, &launch, rpClient)
}
