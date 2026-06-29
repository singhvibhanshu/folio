package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var (
	cropBox   string
	cropPages string
)

var cropCmd = &cobra.Command{
	Use:   "crop <file.pdf | folder>",
	Short: "Crop page margins",
	Long: `Trim pages to a region using pdfcpu's box syntax, for example:

  --box "10% 10% 80% 80%"     keep the centered 80% region
  --box "[0 0 200 300]"        absolute crop box in points`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cropBox == "" {
			return errors.New("a crop region is required: --box \"10% 10% 80% 80%\"")
		}
		pages := parsePages(cropPages)
		return runTransform(args[0], "cropped", func(in, out string) error {
			return pdfops.Crop(in, out, cropBox, pages)
		})
	},
}

func init() {
	cropCmd.Flags().StringVar(&cropBox, "box", "", "crop region (percent or absolute points)")
	cropCmd.Flags().StringVar(&cropPages, "pages", "", "pages to crop (default: all)")
}
