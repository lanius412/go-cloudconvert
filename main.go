package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lanius412/go-cloudconvert/convert"
	"github.com/lanius412/go-cloudconvert/pdf"
)

type Flags struct {
	InputFile    string
	OutputFormat string
}

var flags Flags

var allFormats = [...]string{"7z", "ACE", "ALZ", "ARC", "ARJ", "BZ", "BZ2", "CAB", "CPIO", "DEB", "DMG", "GZ", "IMG", "ISO", "JAR", "LHA", "LZ", "LZMA", "LZO", "RAR", "RPM", "RZ", "TAR", "TAR.7z", "TAR.BZ", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "TAR.XZ", "TAR.Z", "TBZ", "TBZ2", "TGZ", "TZ", "TZO", "XZ", "Z", "ZIP", "AAC", "AC3", "AIF", "AIFC", "AIFF", "AMR", "AU", "CAF", "FLAC", "M4A", "M4B", "MP3", "OGA", "SFARK", "VOC", "WAV", "WEBA", "WMA", "DWG", "DXF", "ABW", "DJVU", "DOC", "DOCM", "DOCX", "DOT", "DOTX", "HTML", "HWP", "LWP", "MD", "ODT", "PAGES", "PDF", "RST", "RTF", "SDW", "TEX", "TXT", "WPD", "WPS", "ZABW", "AZW", "AZW3", "AZW4", "CBC", "CBR", "CBZ", "CHM", "EPUB", "FB2", "HTM", "HTMLZ", "LIT", "LRF", "MOBI", "PDB", "PML", "PRC", "RB", "SNB", "TCR", "TXTZ", "EOT", "OTF", "TTF", "WOFF", "WOFF2", "3FR", "ARW", "AVIF", "BMP", "CR2", "CR3", "CRW", "DCR", "DNG", "EPS", "ERF", "GIF", "HEIC", "ICNS", "ICO", "JFIF", "JPEG", "JPG", "MOS", "MRW", "NEF", "ODD", "ORF", "PEF", "PNG", "PPM", "PS", "PSD", "RAF", "RAW", "RW2", "TIF", "TIFF", "WEBP", "X3F", "XCF", "XPS", "DPS", "KEY", "ODP", "POT", "POTX", "PPS", "PPSX", "PPT", "PPTM", "PPTX", "SDA", "CSV", "ET", "NUMBERS", "ODS", "SDC", "XLS", "XLSM", "XLSX", "AI", "CDR", "CGM", "EMF", "SK", "SK1", "SVG", "SVGZ", "VSD", "WMF", "3G2", "3GP", "3GPP", "AVI", "CAVS", "DV", "DVR", "FLV", "M2TS", "M4V", "MKV", "MOD", "MOV", "MP4", "MPEG", "MPG", "MTS", "MXF", "OGG", "RM", "RMVB", "SWF", "TS", "VOB", "WEBM", "WMV", "WTV"}

func main() {
	err := CheckInputError()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	outFile, err := convert.Do(flags.InputFile, strings.ToLower(flags.OutputFormat))
	if err != nil {
		log.Fatalln("convert error: ", err)
	}

	if flags.OutputFormat == "pdf" {
		err = pdf.SplitPdf(outFile)
		if err != nil {
			log.Fatalln("split pdf error: ", err)
		}
	}

	fmt.Println("Complete")
}

//checkInputError return error if find unexpected flag argument
func CheckInputError() error {
	if flags.InputFile == "" {
		return errors.New("Input file path")

	}
	if flags.OutputFormat == "" {
		return errors.New("Input output format")
	}

	for _, format := range allFormats {
		if !strings.Contains(filepath.Ext(flags.InputFile), strings.ToLower(format)) || !strings.Contains(strings.ToLower(flags.OutputFormat), strings.ToLower(format)) {
			return errors.New("Unsupported File. Check supported format: option -h")
		}
	}
	return nil
}

func init() {
	inputFlag := flag.String("i", "", "Required : input file")
	outputFlag := flag.String("f", "", "Required : output file format like 'pdf' ")
	helpFlag := flag.Bool("h", false, "print supported format")

	flag.Parse()

	if *helpFlag {
		printFormats()
		os.Exit(0)
	}

	flags.InputFile = *inputFlag
	flags.OutputFormat = *outputFlag
}

//printFormats print Supported Format when call help flag
func printFormats() {
	formatKinds := [...]string{"Archive", "Audio", "Cad", "Document", "EBook", "Font", "Image", "PResentation", "Spredsheet", "Vector", "Video"}

	var formats [11][]string
	formats[0] = []string{"7z", "ACE", "ALZ", "ARC", "ARJ", "BZ", "BZ2", "CAB", "CPIO", "DEB", "DMG", "GZ", "IMG", "ISO", "JAR", "LHA", "LZ", "LZMA", "LZO", "RAR", "RPM", "RZ", "TAR", "TAR.7z", "TAR.BZ", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "TAR.XZ", "TAR.Z", "TBZ", "TBZ2", "TGZ", "TZ", "TZO", "XZ", "Z", "ZIP"}
	formats[1] = []string{"AAC", "AC3", "AIF", "AIFC", "AIFF", "AMR", "AU", "CAF", "FLAC", "M4A", "M4B", "MP3", "OGA", "SFARK", "VOC", "WAV", "WEBA", "WMA"}
	formats[2] = []string{"DWG", "DXF"}
	formats[3] = []string{"ABW", "DJVU", "DOC", "DOCM", "DOCX", "DOT", "DOTX", "HTML", "HWP", "LWP", "MD", "ODT", "PAGES", "PDF", "RST", "RTF", "SDW", "TEX", "TXT", "WPD", "WPS", "ZABW"}
	formats[4] = []string{"AZW", "AZW3", "AZW4", "CBC", "CBR", "CBZ", "CHM", "EPUB", "FB2", "HTM", "HTMLZ", "LIT", "LRF", "MOBI", "PDB", "PML", "PRC", "RB", "SNB", "TCR", "TXTZ"}
	formats[5] = []string{"EOT", "OTF", "TTF", "WOFF", "WOFF2"}
	formats[6] = []string{"3FR", "ARW", "AVIF", "BMP", "CR2", "CR3", "CRW", "DCR", "DNG", "EPS", "ERF", "GIF", "HEIC", "ICNS", "ICO", "JFIF", "JPEG", "JPG", "MOS", "MRW", "NEF", "ODD", "ORF", "PEF", "PNG", "PPM", "PS", "PSD", "RAF", "RAW", "RW2", "TIF", "TIFF", "WEBP", "X3F", "XCF", "XPS"}
	formats[7] = []string{"DPS", "KEY", "ODP", "POT", "POTX", "PPS", "PPSX", "PPT", "PPTM", "PPTX", "SDA"}
	formats[8] = []string{"CSV", "ET", "NUMBERS", "ODS", "SDC", "XLS", "XLSM", "XLSX"}
	formats[9] = []string{"AI", "CDR", "CGM", "EMF", "SK", "SK1", "SVG", "SVGZ", "VSD", "WMF"}
	formats[10] = []string{"3G2", "3GP", "3GPP", "AVI", "CAVS", "DV", "DVR", "FLV", "M2TS", "M4V", "MKV", "MOD", "MOV", "MP4", "MPEG", "MPG", "MTS", "MXF", "OGG", "RM", "RMVB", "SWF", "TS", "VOB", "WEBM", "WMV", "WTV"}

	for i, kind := range formatKinds {
		fmt.Println(kind + " >")
		for j, format := range formats[i] {
			fmt.Print("・" + format)
			if j%8 == 0 && j != 0 {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}
