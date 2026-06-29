package cmd

import (
	"fmt"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var frompicsCmd = &cobra.Command{
	Use:   "frompics <img1> <img2...>",
	Short: "Build a PDF from images (JPG/PNG → PDF)",
	Long:  "Create a PDF with one image per page, in the order given. Accepts JPG, PNG, etc.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := outFlag
		if out == "" || outIsDir() {
			out = outputPathPlain("images.pdf")
		}
		if err := ensureParent(out); err != nil {
			return err
		}
		if err := pdfops.ImagesToPDF(args, out); err != nil {
			return err
		}
		after, _ := fileSize(out)
		fmt.Printf("✓ built PDF from %d image(s)\n  → %s  (%s)\n", len(args), out, humanBytes(after))
		return nil
	},
}
