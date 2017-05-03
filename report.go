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

// XMLReport identifies JUnit XML format specification that Hudson supports
type XMLReport struct {
	xmlSuites []xmlSuite
}

type xmlSuite struct {
	XMLName     string        `xml:"testsuite"`
	ID          int           `xml:"id,attr"`
	Name        string        `xml:"name,attr"`
	PackageName string        `xml:"package,attr"`
	TimeStamp   string        `xml:"timestamp,attr"`
	Time        float64       `xml:"time,attr"`
	HostName    string        `xml:"hostname,attr"`
	Tests       int           `xml:"tests,attr"`
	Failures    int           `xml:"failures,attr"`
	Errors      int           `xml:"errors,attr"`
	Properties  xmlProperties `xml:"properties"`
	Cases       []xmlTest     `xml:"testcase"`
	SystemOut   string        `xml:"system-out"`
	SystemErr   string        `xml:"system-err"`
}

type xmlProperties struct {
}

type xmlTest struct {
	Name      string      `xml:"name,attr"`
	ClassName string      `xml:"classname,attr"`
	Time      float64     `xml:"time,attr"`
	Failure   *xmlFailure `xml:"failure,omitempty"`
}

type xmlFailure struct {
	Type    string `xml:"type,attr"`
	Message string `xml:"message,attr"`
	Details string `xml:",chardata"`
}

// LoadXMLReport is used for loading JUnit XML report from specified directory
func LoadXMLReport(dirName string) (*XMLReport, error) {
	report, err := parseXMLReport(dirName)
	if err != nil {
		return nil, err
	}
	return &XMLReport{
		xmlSuites: report,
	}, nil
}

//
func (report *XMLReport) GetLaunchStartTime() time.Time {
	return parseTimeStamp(report.xmlSuites[0].TimeStamp)
}

//
func (report *XMLReport) GetLaunchEndTime() time.Time {
	lastIndex := len(report.xmlSuites) - 1
	return parseTimeStamp(report.xmlSuites[lastIndex].TimeStamp)
}

// parseXMLReport is used for parsing xml report sorted by suite start time
func parseXMLReport(reportDir string) ([]xmlSuite, error) {

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

		b, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			log.Error(err)
			continue
		}

		var xSuite xmlSuite
		xml.Unmarshal(b, &xSuite)
		xSuites[i] = xSuite
	}

	// sort by start time
	sort.Slice(xSuites, func(i, j int) bool {
		t1 := parseTimeStamp(xSuites[i].TimeStamp)
		t2 := parseTimeStamp(xSuites[j].TimeStamp)
		return t1.Before(t2)
	})

	return xSuites, nil
}
