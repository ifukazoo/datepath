package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yamanobori-old/datepath/common"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <file path>\n", os.Args[0])
		flag.PrintDefaults()
	}

}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := common.UpdateTimePath(flag.Arg(0)); err != nil {
		log.Fatal(err.Error())
	}
}
