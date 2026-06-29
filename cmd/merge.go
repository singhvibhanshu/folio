package cmd

import (
	"fmt"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <a.pdf> <b.pdf> [more.pdf...]",
	Short: "Combine several PDFs into one",
	Long:  "Combine two or more PDFs into a single document, in the order given.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := outFlag
		if out == "" || outIsDir() {
			out = outputPathPlain("merged.pdf")
		}
		if err := ensureParent(out); err != nil {
			return err
		}
		if err := pdfops.Merge(args, out); err != nil {
			return err
		}
		after, _ := fileSize(out)
		fmt.Printf("✓ merged %d files\n  → %s  (%s)\n", len(args), out, humanBytes(after))
		return nil
	},
}
