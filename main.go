package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Tree(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get filename from file info
		filename := info.Name()

		// Ignore hidden files
		if filename[0] == '.' {
			// Special error to stop recursing down the directory
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			fmt.Errorf("Rel(%s, %s): %v", root, path, err)
			return err
		}

		depth := len(strings.Split(rel, string(filepath.Separator)))

		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), filename)
		return nil
	})
	return err
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := Tree(arg)
		if err != nil {
			log.Printf("Tree %s: %v\n", arg, err)
		}
	}
}
