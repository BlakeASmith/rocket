package cmd

import (
	"fmt"
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

	query := strings.Join(args, " ")
	selected := selectDir(rocketRoot, query)
	if selected != "" {
		fmt.Println(selected)
	}
}

func init() {
	rootCmd.AddCommand(gotoCmd)
}
