package rp

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func PublishReport(reportDir string, launch *Launch, c Client) {
	xSuites, err := parseJunitXMLReport(reportDir)
	if err != nil {
		log.Critical(err)
		return
	}
	launchStartTime, _ := time.Parse(TimestampLayout, xSuites[0].TimeStamp)
	launch.StartTime = launchStartTime
	launchID := c.StartLaunch(launch)
	if launchID == nil {
		log.Critical("could not start launch")
		return
	}

	for i := 0; i < len(xSuites); i++ {
		xSuite := xSuites[i]
		suiteStartTime, _ := time.Parse(TimestampLayout, xSuite.TimeStamp)
		// TODO: end time
		suiteID := c.StartTestItem("", &TestItem{
			LaunchID:    launchID.ID,
			Type:        TestItemTypeSuite,
			StartTime:   suiteStartTime,
			Name:        xSuite.Name,
			Description: xSuite.PackageName,
		})

		//finis suite
		c.FinishTestItem(suiteID.ID, &ExecutionResult{
			Status:  ExecutionStatusFailed,
			EndTime: suiteStartTime.Add(time.Duration(xSuite.Time)),
		})
	}

	launchEndTime, _ := time.Parse(TimestampLayout, xSuites[len(xSuites)].TimeStamp)
	c.FinishLaunch(launchID.ID, &ExecutionResult{
		Status:  ExecutionStatusFailed,
		EndTime: launchEndTime,
	})
}

// parseJunitXMLReport is used for parsing Junit xml report is a sorted by suite time order
func parseJunitXMLReport(reportDir string) ([]xmlSuite, error) {

	if len(reportDir) == 0 {
		return nil, errors.New("report dir could not be empty")
	}
	files, err := ioutil.ReadDir(reportDir)
	if err != nil {
		return nil, err
	}

	n := len(files)
	xSuites := make([]xmlSuite, n)

	// read all files in report dir
	for i := 0; i < n; i++ {
		f := files[i]
		xmlFile, err := os.Open(filepath.Join(reportDir, f.Name()))
		defer xmlFile.Close()
		if err != nil {
			log.Error(err)
			continue
		}
		b, _ := ioutil.ReadAll(xmlFile)
		var xSuite xmlSuite
		xml.Unmarshal(b, &xSuite)
		xSuites[i] = xSuite
	}

	// sort by time
	sort.Slice(xSuites, func(i, j int) bool {
		t1, _ := time.Parse(TimestampLayout, xSuites[i].TimeStamp)
		t2, _ := time.Parse(TimestampLayout, xSuites[j].TimeStamp)
		return t1.Before(t2)
	})

	// debug
	for i := 0; i < len(xSuites); i++ {
		log.Debugf("Suite %d time: %s \n", xSuites[i].ID, xSuites[i].TimeStamp)
	}

	return xSuites, nil
}
