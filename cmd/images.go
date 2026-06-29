package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var imagesPages string

var imagesCmd = &cobra.Command{
	Use:   "images <file.pdf>",
	Short: "Extract embedded images from a PDF",
	Long: `Save the images embedded in a PDF as separate files.

Note: this pulls out the bitmaps actually stored inside the document. It does
not rasterize/screenshot pages (that would need a rendering engine and break
folio's single-binary, fully-offline design).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		in := args[0]
		dir := outFlag
		if dir == "" {
			base := strings.TrimSuffix(filepath.Base(in), filepath.Ext(in))
			dir = filepath.Join(filepath.Dir(in), base+"_images")
		}
		if err := pdfops.ExtractImages(in, dir, parsePages(imagesPages)); err != nil {
			return err
		}
		fmt.Printf("✓ extracted images from %s\n  → %s/\n", in, dir)
		return nil
	},
}

func init() {
	imagesCmd.Flags().StringVar(&imagesPages, "pages", "", "pages to pull images from (default: all)")
}
