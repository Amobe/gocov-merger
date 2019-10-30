package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amobe/gocov-merger/profile"
)

var usage = `Usage: gocov-merger [options] [files...]

Options:
	-q		quite mode
	-o file		specific output file
	-h		help page
`

func main() {
	err := mainWithCode()
	if err != nil {
		os.Exit(1)
	}
}

func mainWithCode() error {
	flag.Usage = func() {
		fmt.Println(usage)
	}
	outputFileName := flag.String("o", "", "specific output file")
	quiteMode := flag.Bool("q", false, "quite mode")
	flag.Parse()
	coverFiles := flag.Args()
	printReport := *outputFileName == ""

	merged := profile.NewEmptyProfile()
	mergeCnt := 0

	for _, file := range coverFiles {
		p, err := profile.NewProfile(file)
		if err != nil {
			fmt.Printf("failed to parse coverprofile %s, %v\n", file, err)
			continue
		}
		merged.Merge(p)
		mergeCnt++
	}

	if printReport {
		merged.Dump()
	} else {
		err := merged.Write(*outputFileName)
		if err != nil {
			fmt.Printf("failed to write file %s, %v\n", *outputFileName, err)
		}
		if !*quiteMode {
			fmt.Printf("Merged %d reports to %s\n", mergeCnt, *outputFileName)
		}
	}

	return nil
}
