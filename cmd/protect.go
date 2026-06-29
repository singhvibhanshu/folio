package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var (
	protectPassword string
	protectOwnerPW  string
)

var protectCmd = &cobra.Command{
	Use:   "protect <file.pdf | folder>",
	Short: "Password-protect a PDF (AES-256)",
	Long:  "Encrypt a PDF so it can only be opened with the given password.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if protectPassword == "" {
			return errors.New("a password is required: --password <pw>")
		}
		return runTransform(args[0], "protected", func(in, out string) error {
			return pdfops.Protect(in, out, protectPassword, protectOwnerPW)
		})
	},
}

func init() {
	protectCmd.Flags().StringVarP(&protectPassword, "password", "p", "", "password required to open the PDF")
	protectCmd.Flags().StringVar(&protectOwnerPW, "owner", "", "separate owner password (default: same as --password)")
}
