package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rocket",
	Short: "Manage your project folders",
	Long: `Manage your project folders

	- TODO
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Set up config
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".config", "rocket")
	configPath := filepath.Join(configDir, "config.yml")

	// Ensure config dir exists
	os.MkdirAll(configDir, 0755)

	// If config doesn't exist, create with default
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, _ := os.Create(configPath)
		file.WriteString("rocket_root: ~/rocket\n")
		file.Close()
	}

	// Set up viper
	viper.SetConfigFile(configPath)
	viper.ReadInConfig()
	viper.SetDefault("rocket_root", "~/rocket")
}

// GetRocketRoot returns the expanded rocket_root path
func GetRocketRoot() string {
	root := viper.GetString("rocket_root")
	if strings.HasPrefix(root, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, root[2:])
	}
	return root
}
