package cmd

import (
	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var (
	nupCount int
	nupPages string
)

var nupCmd = &cobra.Command{
	Use:   "nup <file.pdf | folder>",
	Short: "Place multiple pages onto each sheet",
	Long:  "Arrange N source pages per output page (great for handouts): 2, 4, 6, 9, 16.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pages := parsePages(nupPages)
		return runTransform(args[0], "nup", func(in, out string) error {
			return pdfops.NUp(in, out, nupCount, pages)
		})
	},
}

func init() {
	nupCmd.Flags().IntVarP(&nupCount, "count", "n", 4, "source pages per sheet (2, 4, 6, 9, 16)")
	nupCmd.Flags().StringVar(&nupPages, "pages", "", "pages to include (default: all)")
}
