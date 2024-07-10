package main

import (
	"flag"
	"fmt"

	"github.com/pogorelof/ProjectPdf/internal/app/cli"
)

var (
	rootDir                   string
	outputFileName            string
	customProhibitedFilesName string
)

func init() {
	flag.StringVar(&rootDir, "dir", "", "Set project dir")
	flag.StringVar(&outputFileName, "o", "", "Set output file name without extension")
	flag.StringVar(&customProhibitedFilesName, "proh", "", "Use custom prohibited files")
}

func main() {
	cli := cli.New()

	flag.Parse()

	if rootDir != "" {
		cli.SetDir(rootDir)
	}
	if outputFileName != "" {
		cli.SetOutputFileName(outputFileName)
	}
	if customProhibitedFilesName != ""{
		cli.SetCustomProhibit(customProhibitedFilesName)
	}

	cli.Open()
	defer cli.Close()

	cli.StructureToFile()
	cli.RecursiveCopyFiles()

	fmt.Printf("Данные проекта успешно записаны в файл\n")
}
