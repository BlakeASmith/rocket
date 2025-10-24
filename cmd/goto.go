package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// gotoCmd represents the goto command
var gotoCmd = &cobra.Command{
	Use:   "goto",
	Short: "Navigate to a project directory using fuzzy search",
	Long: `Navigate to a project directory under rocket_root using fuzzy search with fzf.

This command lists subdirectories of rocket_root and allows you to select one using fzf.
Once selected, it prints "cd <path>" to stdout, which your shell can execute to change directory.`,
	Run: RunGoto,
}

// This function needs to print a directory path to STDOUT
// Any other printing/logging must be done on stderr
func RunGoto(cmd *cobra.Command, args []string) {
	rocketRoot := GetRocketRoot()

	if len(args) == 0 {
		// Original behavior: list all and fzf
		matches := getMatchingDirs(rocketRoot, "")
		if len(matches) == 0 {
			return
		}

		// Compute relatives for fzf display
		relatives := make([]string, len(matches))
		for i, match := range matches {
			rel, _ := filepath.Rel(rocketRoot, match)
			relatives[i] = rel
		}

		fzfCmd := exec.Command("fzf")
		fzfCmd.Stdin = strings.NewReader(strings.Join(relatives, "\n"))

		output, err := fzfCmd.Output()
		if err != nil {
			return
		}

		selectedRel := strings.TrimSpace(string(output))
		// Find the corresponding absolute and print it
		for i, rel := range relatives {
			if rel == selectedRel {
				fmt.Println(matches[i])
				return
			}
		}
		return
	}

	// Fuzzy match mode
	query := strings.Join(args, " ")
	selected := selectDir(rocketRoot, query)
	if selected != "" {
		fmt.Println(selected)
	}
}

func init() {
	rootCmd.AddCommand(gotoCmd)
}
