package cmd

import (
	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var pagenumPages string

var pagenumCmd = &cobra.Command{
	Use:   "pagenum <file.pdf | folder>",
	Short: "Add page numbers",
	Long:  `Stamp "<n> of <total>" centered at the bottom of each page.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pages := parsePages(pagenumPages)
		return runTransform(args[0], "numbered", func(in, out string) error {
			return pdfops.PageNumbers(in, out, pages)
		})
	},
}

func init() {
	pagenumCmd.Flags().StringVar(&pagenumPages, "pages", "", "pages to number, e.g. 2- (default: all)")
}
