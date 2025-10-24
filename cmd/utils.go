package cmd

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// getMatchingDirs returns directories under rocketRoot that fuzzy match the query
func getMatchingDirs(rocketRoot, query string) []string {
	findCmd := exec.Command("find", rocketRoot, "-type", "d", "-mindepth", "1", "-maxdepth", "1")
	output, err := findCmd.Output()
	if err != nil {
		return nil
	}

	dirs := strings.Split(strings.TrimSpace(string(output)), "\n")
	var matches []string

	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		base := filepath.Base(dir)
		if fuzzy.Match(query, base) {
			matches = append(matches, dir)
		}
	}

	return matches
}
