package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new project",
	Long: `Create a new project with the specified name.

The name parameter is required.
By default, this command will change to the newly created directory.
Use --no-go to prevent changing directory.`,
	Args: cobra.ExactArgs(1),
	Run:  RunNew,
}

var noGo bool

func RunNew(cmd *cobra.Command, args []string) {
	name := args[0]

	// Get rocket root
	rocketRoot := GetRocketRoot()

	// Create the project directory under rocket root
	projectPath := filepath.Join(rocketRoot, name)
	err := os.MkdirAll(projectPath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", projectPath, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Created new project directory: %s\n", projectPath)

	// Change to the new directory unless --no-go is specified
	if !noGo {
		err = os.Chdir(projectPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error changing to directory %s: %v\n", projectPath, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Changed to directory: %s\n", projectPath)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolVar(&noGo, "no-go", false, "Don't change to the newly created directory")
}
