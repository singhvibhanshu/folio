package cmd

import (
	"fmt"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <file.pdf>",
	Short: "Show document details",
	Long:  "Print page count, dimensions, metadata, and security flags for a PDF.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		nfo, err := pdfops.Info(in)
		if err != nil {
			return err
		}
		size, _ := fileSize(in)

		fmt.Printf("%s\n", in)
		row := func(k, v string) {
			if v != "" {
				fmt.Printf("  %-14s %s\n", k+":", v)
			}
		}
		row("Pages", fmt.Sprintf("%d", nfo.PageCount))
		row("Size", humanBytes(size))
		row("PDF version", nfo.Version)
		if len(nfo.Dimensions) > 0 {
			d := nfo.Dimensions[0]
			row("Page size", fmt.Sprintf("%.0f × %.0f pt", d.Width, d.Height))
		}
		row("Title", nfo.Title)
		row("Author", nfo.Author)
		row("Subject", nfo.Subject)
		row("Creator", nfo.Creator)
		row("Producer", nfo.Producer)
		row("Created", nfo.CreationDate)
		row("Modified", nfo.ModificationDate)
		row("Tagged", yesno(nfo.Tagged))
		row("Watermarked", yesno(nfo.Watermarked))
		row("Has form", yesno(nfo.Form))
		row("Signed", yesno(nfo.Signatures))
		row("Bookmarks", yesno(nfo.Outlines))
		return nil
	},
}

func yesno(b bool) string {
	if b {
		return "yes"
	}
	return ""
}
