package cmd

import (
	"fmt"
	"os/exec"
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

func RunGoto(cmd *cobra.Command, args []string) {
	rocketRoot := GetRocketRoot()

	// Find subdirectories
	findCmd := exec.Command("find", rocketRoot, "-type", "d", "-mindepth", "1", "-maxdepth", "1")
	fzfCmd := exec.Command("fzf")

	// Pipe find output to fzf
	fzfCmd.Stdin, _ = findCmd.StdoutPipe()

	// Run find in background
	findCmd.Start()

	// Get fzf output
	output, err := fzfCmd.Output()
	if err != nil {
		// fzf was cancelled or no selection
		return
	}

	selected := strings.TrimSpace(string(output))
	if selected != "" {
		fmt.Println("cd " + selected)
	}
}

func init() {
	rootCmd.AddCommand(gotoCmd)
}
