package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell integration code for installation",
	Long:  "Print the shell integration code that should be added to your shell profile",
	Run: func(cmd *cobra.Command, args []string) {
		// Read and print init.sh content
		initPath := filepath.Join("init.sh")
		content, err := os.ReadFile(initPath)
		if err != nil {
			fmt.Printf("Error reading init.sh: %v\n", err)
			return
		}
		fmt.Print(string(content))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
