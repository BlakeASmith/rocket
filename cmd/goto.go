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

// This function needs to print a directory path to STDOUT
// Any other printing/logging must be done on stderr
func RunGoto(cmd *cobra.Command, args []string) {
	rocketRoot := GetRocketRoot()

	if len(args) == 0 {
		// Original behavior: list all and fzf
		findCmd := exec.Command("find", rocketRoot, "-type", "d", "-mindepth", "1", "-maxdepth", "1")
		fzfCmd := exec.Command("fzf")

		fzfCmd.Stdin, _ = findCmd.StdoutPipe()
		findCmd.Start()

		output, err := fzfCmd.Output()
		if err != nil {
			return
		}

		selected := strings.TrimSpace(string(output))
		if selected != "" {
			fmt.Println(selected)
		}
		return
	}

	// Fuzzy match mode
	query := strings.Join(args, " ")
	matches := getMatchingDirs(rocketRoot, query)

	if len(matches) == 1 {
		fmt.Println(matches[0])
	} else if len(matches) > 1 {
		// Pipe matches to fzf
		fzfCmd := exec.Command("fzf")
		fzfCmd.Stdin = strings.NewReader(strings.Join(matches, "\n"))

		output, err := fzfCmd.Output()
		if err != nil {
			return
		}

		selected := strings.TrimSpace(string(output))
		if selected != "" {
			fmt.Println(selected)
		}
	}
	// If no matches, do nothing
}

func init() {
	rootCmd.AddCommand(gotoCmd)
}
