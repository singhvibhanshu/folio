package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var extractPages string

var extractCmd = &cobra.Command{
	Use:   "extract <file.pdf>",
	Short: "Keep (and optionally reorder) selected pages",
	Long: `Build a new PDF from just the pages you choose.

The output follows the order you give, so it also reorders:
  folio extract in.pdf --pages 1-3      keep pages 1 to 3
  folio extract in.pdf --pages 3,1,2    reorder the first three pages`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pages := parsePages(extractPages)
		if len(pages) == 0 {
			return errors.New("specify which pages to keep: --pages 1-3,5")
		}
		return runTransform(args[0], "pages", func(in, out string) error {
			return pdfops.Extract(in, out, pages)
		})
	},
}

func init() {
	extractCmd.Flags().StringVar(&extractPages, "pages", "", "pages to keep, e.g. 1-3,5 (order is preserved)")
}
