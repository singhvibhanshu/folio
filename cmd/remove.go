package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var removePages string

var removeCmd = &cobra.Command{
	Use:   "remove <file.pdf>",
	Short: "Delete selected pages",
	Long:  "Produce a copy of the PDF with the given pages removed.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pages := parsePages(removePages)
		if len(pages) == 0 {
			return errors.New("specify which pages to remove: --pages 2,4")
		}
		return runTransform(args[0], "removed", func(in, out string) error {
			return pdfops.Remove(in, out, pages)
		})
	},
}

func init() {
	removeCmd.Flags().StringVar(&removePages, "pages", "", "pages to delete, e.g. 2,4,6-8")
}
