package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yamanobori-old/datepath/common"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <dir path> [job name]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	if err := os.Chdir(flag.Arg(0)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	jobName := ""
	if flag.NArg() > 1 {
		jobName = flag.Arg(1)
	}

	if err := common.CreateEmptyFile(jobName); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
