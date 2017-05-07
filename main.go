package main

import (
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/darkowl91/rp-client/rp"
)

var (
	// Version for the application, passed on build
	Version = "No Version Provided"

	reportDir string

	debugFlag   bool
	hostFlag    string
	projectFlag string
	launchFlag  string
	tagsFlag    string
	uuidFlag    string
	helpFlag    bool
	versionFlag bool
)

func main() {
	flag.Usage = func() {
		u := " Usage:\n"
		u += "  rp-client [OPTIONS] (DIR|FILE)\n\n"
		u += " Options:\n"
		u += "	-r	--rp		Report Portal host\n"
		u += "	-d	--debug		Report Portal debug mode\n"
		u += "	-p	--project	Report Portal project\n"
		u += "	-l	--launch	Report Portal launch name\n"
		u += "	-t	--tags		Report Portal launch tags\n"
		u += "	-id	--uuid		Report Portal user id\n"
		u += "	-h,	--help		Print usage\n"
		u += "	-v,	--version	Print version information and quit\n\n"
		u += " Example:\n"
		u += "	rp-client -r http://example.com/api/v1/ -p PROJECT -l LAUNCH -t tag1,tag2,tag3 -id your_id ./examples/report"
		fmt.Fprintf(os.Stdout, u)
	}

	flag.StringVar(&hostFlag, "r", "", "Report Portal host. Example: http://example.com/api/v1/.")
	flag.StringVar(&hostFlag, "rp", "", "Report Portal host. Example: http://example.com/api/v1/.")

	flag.BoolVar(&debugFlag, "d", false, "Enable Report Portal debug mode.")
	flag.BoolVar(&debugFlag, "debug", false, "Enable Report Portal debug mode.")

	flag.StringVar(&projectFlag, "p", "", "Report Portal project.")
	flag.StringVar(&projectFlag, "project", "", "Report Portal project.")

	flag.StringVar(&launchFlag, "l", "", "Report Portal launch name.")
	flag.StringVar(&launchFlag, "launch", "", "Report Portal launch.")

	flag.StringVar(&tagsFlag, "t", "", "Report Portal launch tags.")
	flag.StringVar(&tagsFlag, "tags", "", "Report Portal tags.")

	flag.StringVar(&uuidFlag, "id", "", "Report Portal user id.")
	flag.StringVar(&uuidFlag, "uuid", "", "Report Portal user id.")

	flag.BoolVar(&helpFlag, "h", false, "Print usage.")
	flag.BoolVar(&helpFlag, "help", false, "Print usage.")

	flag.BoolVar(&versionFlag, "v", false, "Print version.")
	flag.BoolVar(&versionFlag, "version", false, "Print version.")
	flag.Parse()

	if versionFlag {
		fmt.Printf("rp version: %s\n", Version)
		os.Exit(1)
	}

	if helpFlag {
		flag.Usage()
		os.Exit(1)
	}

	reportDir = flag.Arg(0)
	if len(reportDir) == 0 {
		fmt.Println("specify directory with tests results")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(reportDir); os.IsNotExist(err) {
		fmt.Printf("invalid tests results directory path '%s'\n", reportDir)
		flag.Usage()
		os.Exit(1)
	}

	if len(hostFlag) == 0 {
		fmt.Println("report portal host cold not be empty")
		flag.Usage()
		os.Exit(1)
	}

	if len(projectFlag) == 0 {
		fmt.Println("report portal project cold not be empty")
		flag.Usage()
		os.Exit(1)
	}

	if len(uuidFlag) == 0 {
		fmt.Println("report portal user id could not be empty")
		flag.Usage()
		os.Exit(1)
	}

	report, err := rp.LoadXMLReport(reportDir)

	if err != nil {
		fmt.Printf("could not load report")
		os.Exit(1)
	}

	if len(launchFlag) == 0 {
		fmt.Println("report portal launch name could not be empty")
		flag.Usage()
		os.Exit(1)
	}

	launchStart := report.LaunchStartTime()
	fmt.Printf("launch start: %s\n", launchStart.Format(rp.TimestampLayout))
	launch := &rp.Launch{}
	launch.StartTime = launchStart
	launch.Name = launchFlag
	if debugFlag {
		rp.InitLogger()
		launch.Mode = rp.ModeDebug
	}
	if len(tagsFlag) != 0 {
		launch.Tags = strings.Split(tagsFlag, ",")
	}

	rpClient := rp.NewClient(hostFlag, projectFlag, uuidFlag)

	launchID := rpClient.StartLaunch(launch)
	if launchID == nil {
		fmt.Println("could not start launch")
		os.Exit(1)
	}

	for i := 0; i < report.SuitesCount(); i++ {
		suite := report.Suite(i)
		suite.LaunchID = launchID.ID
		suiteID := rpClient.StartTestItem("", suite)

		for j := 0; j < report.TesCaseCount(i); j++ {
			tCase := report.TestCase(i, j)
			tCase.LaunchID = launchID.ID
			tCaseID := rpClient.StartTestItem(suiteID.ID, tCase)

			if report.HasTestCasefailure(i, j) {
				tFailure := report.TestCasefailure(i, j)
				tFailure.ItemID = tCaseID.ID
				rpClient.SendMesssage(tFailure)
				fDetails := report.TestCaseFailureDetails(i, j)
				fDetails.ItemID = tCaseID.ID
				rpClient.SendMesssage(fDetails)
			}

			tResult := report.TestCaseResult(i, j)
			rpClient.FinishTestItem(tCaseID.ID, tResult)
		}

		suiteResult := report.SuiteResult(i)
		rpClient.FinishTestItem(suiteID.ID, suiteResult)
	}

	launchEnd := report.LaunchEndTime()
	fmt.Printf("launch end: %s\n", launchEnd.Format(rp.TimestampLayout))
	rpClient.FinishLaunch(launchID.ID, &rp.ExecutionResult{
		EndTime: launchEnd,
	})

	os.Exit(0)
}
