package cmd

import (
	"errors"

	"github.com/singhvibhanshu/folio/internal/pdfops"
	"github.com/spf13/cobra"
)

var unlockPassword string

var unlockCmd = &cobra.Command{
	Use:   "unlock <file.pdf | folder>",
	Short: "Remove password protection from a PDF",
	Long:  "Decrypt a password-protected PDF you have the password for, producing an open copy.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if unlockPassword == "" {
			return errors.New("the current password is required: --password <pw>")
		}
		return runTransform(args[0], "unlocked", func(in, out string) error {
			return pdfops.Unlock(in, out, unlockPassword)
		})
	},
}

func init() {
	unlockCmd.Flags().StringVarP(&unlockPassword, "password", "p", "", "the PDF's current password")
}
