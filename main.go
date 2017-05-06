package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// Version for the application, passed on build
	Version = "No Version Provided"

	reportDir string
)

func main() {
	flag.Usage = func() {
		u := " Usage:\n"
		u += "  rp [OPTIONS] (DIR|FILE)\n\n"
		u += " Options:\n"
		u += "	-r,	--report	XML Report dir path\n"
		u += "	-H	--rpHost	Report Portal host\n"
		u += "	-m	--mode		Report Portal mode\n"
		u += "	-p	--project	Report Portal project\n"
		u += "	-l	--launch	Report Portal launch name\n"
		u += "	-id	--uuid		Report Portal user id\n"
		u += "	-h,	--help		Print usage\n"
		u += "	-v,	--version	Print version information and quit\n\n"
		fmt.Fprintf(os.Stdout, u)
	}

	flag.Usage()

	fmt.Print(Version)

}
