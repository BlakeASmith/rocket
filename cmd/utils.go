package cmd

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// getMatchingDirs returns directories under rocketRoot that fuzzy match the query
// If query is empty, returns all directories
func getMatchingDirs(rocketRoot, query string) []string {
	findCmd := exec.Command("find", rocketRoot, "-type", "d", "-mindepth", "1")
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
		if query == "" {
			matches = append(matches, dir)
		} else {
			relative, err := filepath.Rel(rocketRoot, dir)
			if err != nil {
				continue
			}
			if fuzzy.Match(query, relative) {
				matches = append(matches, dir)
			}
		}
	}

	return matches
}

// selectDir performs fuzzy matching on directories under rocketRoot with the query,
// and if multiple matches, uses fzf to select one. Returns the selected absolute directory path or empty string.
// Shows relative paths in fzf for better UX.
func selectDir(rocketRoot, query string) string {
	matches := getMatchingDirs(rocketRoot, query)

	if len(matches) == 1 {
		return matches[0]
	} else if len(matches) > 1 {
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
			return ""
		}

		selectedRel := strings.TrimSpace(string(output))
		// Find the corresponding absolute
		for i, rel := range relatives {
			if rel == selectedRel {
				return matches[i]
			}
		}
	}
	return ""
}
