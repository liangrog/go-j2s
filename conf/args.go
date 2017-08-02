package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultOutFileName = "j2s"
)

// Allowing the same argument multiple time
type MultiStrVar []string

// Presenting a comma-separated string
func (v *MultiStrVar) String() string {
	return strings.Join(*v, ",")
}

// Append multiple string vars to the slice
func (v *MultiStrVar) Set(s string) error {
	if *v == nil {
		*v = make([]string, 0, 1)
	}

	*v = append(*v, s)
	return nil
}

// Allowed user input parameters
type Args struct {
	// Path to find json file
	// All json files found will be output to
	// one file
	InPath string

	// Paths that should be excluded
	ExclPath MultiStrVar

	// Path to output go file
	OutPath string

	// If it's recursive
	Recursive bool

	// Package name, if ignored, it'll be using directory name
	PkgName string

	// Input json file names with relative path
	// If specified, take over InPath
	InFile MultiStrVar

	// Output go file name
	OutFileName string
}

// Validating the user input  parameters
func (a *Args) Validate() error {
	if len(a.PkgName) == 0 {
		return fmt.Errorf("Missing package name")
	}

	// Input path must exit
	_, err := os.Lstat(a.InPath)
	if err != nil {
		return fmt.Errorf("Given input path doesn't exist or empty")
	}

	// OutPath can not be empty
	if len(a.OutPath) == 0 {
		return fmt.Errorf("Given output path doesn't exist or empty")
	}

	// Create OutPath if doesn't exist
	_, err = os.Lstat(a.OutPath)
	if err != nil {
		if err = os.MkdirAll(a.OutPath, 0744); err != nil {
			return fmt.Errorf("Cannot create given output path")
		}
	}

	// Check if specified Json file exist
	if len(a.InFile) > 0 {
		for _, p := range a.InFile {
			if _, err := os.Stat(p); os.IsNotExist(err) {
				return fmt.Errorf("Cannot find specified json file", p)
			}
		}
	}

	return nil
}

// Get output file full path and name
func (a *Args) GetOutFile() string {
	s := []string{filepath.Join(a.OutPath, a.OutFileName), "go"}
	return strings.Join(s, ".")
}

// Initiating the cmd args with default values
func NewArgs() (a *Args) {
	// Current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get current working directory stats")
	}

	a = new(Args)
	a.InPath = cwd
	a.ExclPath = MultiStrVar{}
	a.PkgName = "main"
	a.Recursive = false
	a.InFile = MultiStrVar{}
	a.OutPath = cwd
	a.OutFileName = defaultOutFileName
	return
}
