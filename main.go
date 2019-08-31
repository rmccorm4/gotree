package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		fmt.Errorf("Could not use stat %s: %v", root, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Errorf("Could not read dir %s: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		// Skip hidden files/directories
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		childIndent := "│  "
		// If we're at the last child of this directory, print nice corner
		if i == len(names)-1 {
			fmt.Printf(indent + "└──")
			childIndent = "   "
		} else {
			fmt.Printf(indent + "├──")
		}

		if err := tree(filepath.Join(root, name), indent+childIndent); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}
