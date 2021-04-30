package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func SplitPdf(pdfFile string) error {
	config := pdfcpu.NewDefaultConfiguration()
	err := api.SplitFile(pdfFile, ".", 1, config)
	return err
}
