package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List project directories",
	Long: `List project directories under rocket_root.

If no arguments, lists all subdirectories. If a query is provided,
lists only directories that fuzzy match the query.`,
	Run: RunLs,
}

func RunLs(cmd *cobra.Command, args []string) {
	rocketRoot := GetRocketRoot()
	query := strings.Join(args, " ")
	matches := getMatchingDirs(rocketRoot, query)

	for _, dir := range matches {
		rel, _ := filepath.Rel(rocketRoot, dir)
		fmt.Println(rel)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
