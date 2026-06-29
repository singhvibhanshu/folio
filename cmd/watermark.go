package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var (
	watermarkText  string
	watermarkPages string
)

var watermarkCmd = &cobra.Command{
	Use:   "watermark <file.pdf | folder>",
	Short: "Stamp text across pages",
	Long:  "Overlay semi-transparent diagonal text (e.g. CONFIDENTIAL, DRAFT) on the pages.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if watermarkText == "" {
			return errors.New("watermark text is required: --text \"CONFIDENTIAL\"")
		}
		pages := parsePages(watermarkPages)
		return runTransform(args[0], "watermarked", func(in, out string) error {
			return pdfops.Watermark(in, out, watermarkText, pages)
		})
	},
}

func init() {
	watermarkCmd.Flags().StringVarP(&watermarkText, "text", "t", "", "watermark text")
	watermarkCmd.Flags().StringVar(&watermarkPages, "pages", "", "pages to stamp, e.g. 1-3,5 (default: all)")
}
