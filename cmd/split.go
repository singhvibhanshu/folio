package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var splitSpan int

var splitCmd = &cobra.Command{
	Use:   "split <file.pdf>",
	Short: "Split a PDF into smaller files",
	Long:  "Split a PDF into chunks of --span pages each (default 1 = one file per page).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		dir := outFlag
		if dir == "" {
			base := strings.TrimSuffix(filepath.Base(in), filepath.Ext(in))
			dir = filepath.Join(filepath.Dir(in), base+"_split")
		}
		if err := pdfops.Split(in, dir, splitSpan); err != nil {
			return err
		}
		fmt.Printf("✓ split %s (%d page(s) per file)\n  → %s/\n", in, splitSpan, dir)
		return nil
	},
}

func init() {
	splitCmd.Flags().IntVar(&splitSpan, "span", 1, "pages per output file")
}
