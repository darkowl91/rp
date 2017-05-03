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

	// create new Report Portal client
	rpClient := rp.NewClient(apiURL, project, uuid)

}
