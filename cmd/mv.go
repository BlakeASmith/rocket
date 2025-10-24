package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move or rename a project directory using fuzzy search",
	Long: `Move or rename a project directory under rocket_root using fuzzy search.

The first argument is fuzzy-matched against subdirectory names. If multiple matches,
fzf is used for selection. The second argument is the destination path, which must be
relative to rocket_root. Intermediate directories are created as needed.`,
	Run: RunMv,
}

func RunMv(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: mv requires at least two arguments: source and destination\n")
		return
	}

	rocketRoot := GetRocketRoot()
	sourceQuery := args[0]
	dest := args[1]

	// Validate destination: must be relative to rocket_root
	if strings.HasPrefix(dest, "/") {
		fmt.Fprintf(os.Stderr, "Error: destination must be relative to rocket_root, not absolute\n")
		return
	}
	if strings.HasPrefix(dest, "~") {
		fmt.Fprintf(os.Stderr, "Error: destination must be relative to rocket_root, not home-relative\n")
		return
	}
	if strings.Contains(dest, "..") {
		fmt.Fprintf(os.Stderr, "Error: destination cannot contain '..' (must stay within rocket_root)\n")
		return
	}

	// Resolve source using fuzzy matching
	selectedSource := selectDir(rocketRoot, sourceQuery)
	if selectedSource == "" {
		return // No match, do nothing
	}

	// Build full destination path
	destFull := filepath.Join(rocketRoot, dest)
	destFull = filepath.Clean(destFull)

	// Ensure destination stays within rocket_root
	if !strings.HasPrefix(destFull, rocketRoot+"/") && destFull != rocketRoot {
		fmt.Fprintf(os.Stderr, "Error: destination must be within rocket_root\n")
		return
	}

	// If source and dest are the same, do nothing
	if selectedSource == destFull {
		return
	}

	// Create parent directories if needed
	parentDir := filepath.Dir(destFull)
	err := os.MkdirAll(parentDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directories: %v\n", err)
		return
	}

	// Perform the move/rename
	err = os.Rename(selectedSource, destFull)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error moving directory: %v\n", err)
	} else {
		relOld, _ := filepath.Rel(rocketRoot, selectedSource)
		relNew, _ := filepath.Rel(rocketRoot, destFull)
		fmt.Printf("Moved: %s -> %s\n", relOld, relNew)
	}
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
