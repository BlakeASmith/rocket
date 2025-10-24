package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a project directory using fuzzy search",
	Long: `Remove a project directory under rocket_root using fuzzy search.

This command lists subdirectories of rocket_root and allows you to select one using fzf.
Once selected, it deletes the directory.`,
	Run: RunRm,
}

func RunRm(cmd *cobra.Command, args []string) {
	rocketRoot := GetRocketRoot()

	query := strings.Join(args, " ")
	selected := selectDir(rocketRoot, query)

	if selected == "" {
		return
	}

	// Safety check: ensure selected is under rocketRoot
	if !strings.HasPrefix(selected, rocketRoot) {
		fmt.Fprintf(os.Stderr, "Error: selected directory is not under rocket_root\n")
		return
	}

	// Confirm deletion? For now, just delete
	err := os.RemoveAll(selected)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting directory: %v\n", err)
	} else {
		rel, _ := filepath.Rel(rocketRoot, selected)
		fmt.Printf("Deleted: %s\n", rel)
	}
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
