package cmd

import (
	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress <file.pdf | folder>",
	Short: "Reduce PDF file size (lossless)",
	Long: `Shrink a PDF by removing redundant data and optimizing its structure.

This is lossless: text stays sharp and nothing is re-rendered. Point it at a
folder to compress every PDF inside it.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTransform(args[0], "min", pdfops.Compress)
	},
}
