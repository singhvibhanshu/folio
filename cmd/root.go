package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// outFlag is the shared --out destination. It can be a file (used directly) or
// a directory (outputs are named automatically inside it). Empty means write
// alongside each input.
var outFlag string

var rootCmd = &cobra.Command{
	Use:   "folio",
	Short: "folio — a privacy-first, fully offline PDF toolkit",
	Long: `folio performs common PDF operations entirely on your machine.

No uploads, no servers, no internet — your documents never leave your computer.
Merge, split, compress, rotate, protect, watermark, reorder pages, and more.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the CLI and is the single entry point from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outFlag, "out", "o", "",
		"output file or directory (default: alongside the input)")

	rootCmd.AddCommand(
		mergeCmd, splitCmd, compressCmd, rotateCmd,
		protectCmd, unlockCmd, watermarkCmd, pagenumCmd,
		extractCmd, removeCmd, frompicsCmd, imagesCmd,
		nupCmd, cropCmd, infoCmd,
	)
}

// ---- shared helpers --------------------------------------------------------

// outputPath computes the destination for a single-file op that produces a PDF.
// The suffix (e.g. "min") is inserted before the extension when folio names the
// file itself: report.pdf -> report.min.pdf.
func outputPath(inFile, suffix string) string {
	if outFlag != "" && !outIsDir() {
		return outFlag
	}
	base := strings.TrimSuffix(filepath.Base(inFile), filepath.Ext(inFile))
	name := base + "." + suffix + ".pdf"
	if outFlag != "" {
		return filepath.Join(outFlag, name)
	}
	return filepath.Join(filepath.Dir(inFile), name)
}

// outIsDir reports whether --out should be treated as a directory: either it
// already exists as one, or it has no file extension.
func outIsDir() bool {
	if outFlag == "" {
		return false
	}
	if fi, err := os.Stat(outFlag); err == nil {
		return fi.IsDir()
	}
	return filepath.Ext(outFlag) == ""
}

// ensureParent makes sure the directory holding p exists.
func ensureParent(p string) error {
	return os.MkdirAll(filepath.Dir(p), 0o755)
}

// outputPathPlain places a fixed-named output (e.g. "merged.pdf") inside --out
// when it is a directory, otherwise in the current directory.
func outputPathPlain(name string) string {
	if outFlag != "" && outIsDir() {
		return filepath.Join(outFlag, name)
	}
	return name
}

// pdfInputs resolves an argument to a list of PDF files. A directory expands to
// every *.pdf inside it (enabling batch mode); a file resolves to itself.
func pdfInputs(arg string) ([]string, error) {
	fi, err := os.Stat(arg)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []string{arg}, nil
	}
	matches, _ := filepath.Glob(filepath.Join(arg, "*.pdf"))
	if len(matches) == 0 {
		return nil, fmt.Errorf("no PDF files found in %s", arg)
	}
	return matches, nil
}

// runTransform applies a single-input PDF operation to arg, supporting batch
// mode when arg is a directory. suffix names auto-generated outputs.
func runTransform(arg, suffix string, op func(in, out string) error) error {
	inputs, err := pdfInputs(arg)
	if err != nil {
		return err
	}
	if len(inputs) > 1 && outFlag != "" && !outIsDir() {
		return fmt.Errorf("--out is a single file but %d inputs were given; point --out at a directory", len(inputs))
	}
	for _, in := range inputs {
		out := outputPath(in, suffix)
		if err := ensureParent(out); err != nil {
			return err
		}
		before, _ := fileSize(in)
		if err := op(in, out); err != nil {
			return fmt.Errorf("%s: %w", in, err)
		}
		after, _ := fileSize(out)
		reportWrite(in, out, before, after)
	}
	return nil
}

// parsePages turns "1-3,5,8" into pdfcpu's []string page selection. An empty
// or blank spec returns nil, which pdfcpu treats as "all pages".
func parsePages(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// ---- reporting -------------------------------------------------------------

func fileSize(p string) (int64, error) {
	fi, err := os.Stat(p)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func humanBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for x := n / unit; x >= unit; x /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(n)/float64(div), "KMGT"[exp])
}

// reportWrite prints a single ✓ line. When the size meaningfully changed
// (compress/optimize), it shows the before→after delta.
func reportWrite(in, out string, before, after int64) {
	if before > 0 && after > 0 && after != before {
		pct := 100 * (float64(before) - float64(after)) / float64(before)
		fmt.Printf("✓ %s\n  → %s  (%s → %s, %+.0f%%)\n",
			in, out, humanBytes(before), humanBytes(after), -pct)
		return
	}
	fmt.Printf("✓ %s\n  → %s\n", in, out)
}
