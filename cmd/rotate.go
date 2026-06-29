package cmd

import (
	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var (
	rotateAngle int
	rotatePages string
)

var rotateCmd = &cobra.Command{
	Use:   "rotate <file.pdf | folder>",
	Short: "Rotate pages by a multiple of 90°",
	Long:  "Rotate pages clockwise by --angle (use a negative value for counter-clockwise).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pages := parsePages(rotatePages)
		return runTransform(args[0], "rotated", func(in, out string) error {
			return pdfops.Rotate(in, out, rotateAngle, pages)
		})
	},
}

func init() {
	rotateCmd.Flags().IntVar(&rotateAngle, "angle", 90, "rotation in degrees (90, 180, 270, -90)")
	rotateCmd.Flags().StringVar(&rotatePages, "pages", "", "pages to rotate, e.g. 1-3,5 (default: all)")
}
