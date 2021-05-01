package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-cloudconvert/convert"
)

type Flags struct {
	InputFile    string
	OutputFormat string
}

var flags Flags

func main() {
	err := CheckCanConvert(strings.ReplaceAll(filepath.Ext(flags.InputFile), ".", ""), flags.OutputFormat)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Check convertible formats corresponding to input : option -l arg")
		os.Exit(0)
	}

	outFile, err := convert.Do(flags.InputFile, strings.ToLower(flags.OutputFormat))
	if err != nil {
		log.Fatalln("convert error: ", err)
	}

	fmt.Println("Output:", outFile)

	fmt.Println("Complete")
}

func init() {
	inputFlag := flag.String("i", "", "Required : input file")
	outputFlag := flag.String("f", "", "Required : output file format like 'pdf' ")
	convertibleformatFlag := flag.String("l", "", "Print covertible formats corresponding to input")
	helpFlag := flag.Bool("h", false, "print supported all formats")

	flag.Parse()

	if *convertibleformatFlag != "" {
		PrintConvertibleFormat(*convertibleformatFlag)
		os.Exit(0)
	}

	if *helpFlag {
		PrintAllFormats()
		os.Exit(0)
	}

	if *inputFlag == "" {
		fmt.Println("Input file path")
		os.Exit(0)
	}
	if *outputFlag == "" {
		fmt.Println("Input output format")
		os.Exit(0)
	}

	flags.InputFile = *inputFlag
	flags.OutputFormat = *outputFlag
}
