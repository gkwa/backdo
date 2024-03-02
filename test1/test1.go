package test1

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Match struct {
	Path     string
	RHS      []string
	NotFound bool
}

func RunTest(dirPath1, dirPath2 string, excludeExisting []string) error {
	// Generate lists of files from directory paths
	list1 := generateFileList(dirPath1)
	list2 := generateFileList(dirPath2)

	// Apply substring filters to existing directory paths
	list2 = applySubstringFilters(list2, excludeExisting)

	// Create a map for faster lookup
	list2Map := make(map[string]bool)
	for _, path := range list2 {
		list2Map[path] = true
	}

	// Create a list of matches
	var matches []Match

	// Loop over list 1
	for _, path1 := range list1 {
		baseName1 := filepath.Base(path1)
		var rhs []string

		// Loop over list 2 to find matches
		for _, path2 := range list2 {
			baseName2 := filepath.Base(path2)
			if strings.EqualFold(baseName1, baseName2) {
				rhs = append(rhs, path2)
			}
		}

		// Append to matches if there are any matches
		if len(rhs) > 0 {
			matches = append(matches, Match{
				Path: path1,
				RHS:  rhs,
			})
		} else {
			// If file not found in existing directory, mark it as not found
			matches = append(matches, Match{
				Path:     path1,
				NotFound: true,
			})
		}
	}

	// Print matches
	for _, match := range matches {
		fmt.Printf("Path: %s\n", match.Path)
		if match.NotFound {
			fmt.Println("  Not found in existing directory")
		} else {
			for _, rhs := range match.RHS {
				fmt.Printf("  RHS: %s\n", rhs)
			}
		}
	}

	return nil
}

func GenerateScript(dirPath1, dirPath2 string, excludeExisting []string) error {
	// Generate lists of files from directory paths
	list1 := generateFileList(dirPath1)
	list2 := generateFileList(dirPath2)

	// Apply substring filters to existing directory paths
	list2 = applySubstringFilters(list2, excludeExisting)

	// Create a map for faster lookup
	list2Map := make(map[string]bool)
	for _, path := range list2 {
		list2Map[path] = true
	}

	// Create a list of matches
	var matches []Match

	// Loop over list 1
	for _, path1 := range list1 {
		baseName1 := filepath.Base(path1)
		var rhs []string

		// Loop over list 2 to find matches
		for _, path2 := range list2 {
			baseName2 := filepath.Base(path2)
			if strings.EqualFold(baseName1, baseName2) {
				rhs = append(rhs, path2)
			}
		}

		// Append to matches if there are any matches
		if len(rhs) > 0 {
			matches = append(matches, Match{
				Path: path1,
				RHS:  rhs,
			})
		}
	}

	// Define the bash CLI template
	const scriptTemplate = `#!/bin/bash
set -x
set -e

{{range .}}
mv "{{.Path}}" "{{index .RHS 0}}"{{end}}
`

	// Parse the template
	tmpl, err := template.New("script").Parse(scriptTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Execute the template
	err = tmpl.Execute(os.Stdout, matches)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}

func generateFileList(dirPath string) []string {
	var fileList []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

	return fileList
}

func applySubstringFilters(paths []string, filters []string) []string {
	filteredPaths := paths[:0] // Create a new slice with zero length

Outer:
	for _, path := range paths {
		for _, filter := range filters {
			if strings.Contains(path, filter) {
				continue Outer // Skip this path if it contains any filter substring
			}
		}
		filteredPaths = append(filteredPaths, path)
	}

	return filteredPaths
}
