// Package pdfops is a thin, intention-revealing wrapper around pdfcpu.
//
// It keeps the cobra command layer free of pdfcpu's lower-level types and
// centralizes the (deliberately offline, side-effect-free) configuration.
package pdfops

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func init() {
	// Never read or write ~/.config/pdfcpu. folio stays fully self-contained
	// and touches nothing outside the files you point it at.
	api.DisableConfigDir()
}

// conf returns a fresh configuration with relaxed validation, so folio happily
// processes the slightly-malformed PDFs that real-world tools produce.
func conf() *model.Configuration {
	c := model.NewDefaultConfiguration()
	c.ValidationMode = model.ValidationRelaxed
	return c
}

// Merge concatenates inFiles into a single outFile, in order.
func Merge(inFiles []string, outFile string) error {
	return api.MergeCreateFile(inFiles, outFile, false, conf())
}

// Split writes one PDF per span-sized chunk of inFile into outDir.
// span == 1 produces a file per page.
func Split(inFile, outDir string, span int) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	return api.SplitFile(inFile, outDir, span, conf())
}

// Compress losslessly optimizes inFile (de-duplicates objects, prunes waste).
func Compress(inFile, outFile string) error {
	return api.OptimizeFile(inFile, outFile, conf())
}

// Rotate turns the selected pages by angle (a multiple of 90; negative = CCW).
// Empty pages means every page.
func Rotate(inFile, outFile string, angle int, pages []string) error {
	return api.RotateFile(inFile, outFile, angle, pages, conf())
}

// Protect encrypts inFile so it requires password to open. AES-256.
// If ownerPW is empty it mirrors the user password.
func Protect(inFile, outFile, userPW, ownerPW string) error {
	c := conf()
	c.UserPW = userPW
	if ownerPW == "" {
		ownerPW = userPW
	}
	c.OwnerPW = ownerPW
	c.EncryptUsingAES = true
	c.EncryptKeyLength = 256
	return api.EncryptFile(inFile, outFile, c)
}

// Unlock removes password protection from inFile using the supplied password.
func Unlock(inFile, outFile, password string) error {
	c := conf()
	c.UserPW = password
	c.OwnerPW = password
	return api.DecryptFile(inFile, outFile, c)
}

// Watermark stamps semi-transparent diagonal text across the selected pages.
func Watermark(inFile, outFile, text string, pages []string) error {
	desc := "fontname:Helvetica, points:48, scalefactor:0.9 rel, rotation:45, fillcolor:0.5 0.5 0.5, opacity:0.3, position:c"
	return api.AddTextWatermarksFile(inFile, outFile, pages, true, text, desc, conf())
}

// PageNumbers stamps "<n> of <total>" centered at the bottom of each page.
func PageNumbers(inFile, outFile string, pages []string) error {
	desc := "fontname:Helvetica, points:11, scalefactor:1 abs, position:bc, offset:0 10, fillcolor:0.2 0.2 0.2, rotation:0"
	return api.AddTextWatermarksFile(inFile, outFile, pages, true, "%p of %P", desc, conf())
}

// Extract collects the selected pages into a new PDF, preserving the order
// they are given (so it doubles as a reorder tool, e.g. "3,1,2").
func Extract(inFile, outFile string, pages []string) error {
	return api.CollectFile(inFile, outFile, pages, conf())
}

// Remove deletes the selected pages, keeping the rest.
func Remove(inFile, outFile string, pages []string) error {
	return api.RemovePagesFile(inFile, outFile, pages, conf())
}

// ImagesToPDF builds a PDF with one image per page, in order.
func ImagesToPDF(imgFiles []string, outFile string) error {
	return api.ImportImagesFile(imgFiles, outFile, nil, conf())
}

// ExtractImages writes every embedded image from the selected pages to outDir.
func ExtractImages(inFile, outDir string, pages []string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	return api.ExtractImagesFile(inFile, outDir, pages, conf())
}

// NUp places n source pages onto each output page (2, 4, 6, 9, 16...).
func NUp(inFile, outFile string, n int, pages []string) error {
	nup, err := api.PDFNUpConfig(n, "", conf())
	if err != nil {
		return err
	}
	return api.NUpFile([]string{inFile}, outFile, pages, nup, conf())
}

// Crop trims the selected pages to the given box (e.g. "10% 10% 80% 80%" or
// absolute "[0 0 200 300]"); pdfcpu's box mini-language.
func Crop(inFile, outFile, box string, pages []string) error {
	b, err := api.Box(box, types.POINTS)
	if err != nil {
		return err
	}
	return api.CropFile(inFile, outFile, pages, b, conf())
}

// PageCount returns the number of pages in inFile.
func PageCount(inFile string) (int, error) {
	return api.PageCountFile(inFile)
}

// Info returns parsed document metadata for inFile.
func Info(inFile string) (*pdfcpu.PDFInfo, error) {
	f, err := os.Open(inFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return api.PDFInfo(f, inFile, nil, false, conf())
}
