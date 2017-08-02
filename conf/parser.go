package conf

import (
	"flag"
	"fmt"
	"os"
)

const (
	// Usage text
	usgIn      = "Path where to find the json files"
	usgExcl    = "Path you want to exclude from the json file search"
	usgOut     = "Path where the generated go file will be saved to"
	usgR       = "If json file search is going to be recursive (true or false)"
	usgPackage = "Specific package name the generate go file will be assigned to"
	usgFrom    = "Specific json file you want to generate go codes from"
	usgName    = "Specify what output file name is"
)

// Parse argument to config
func Parse() *Args {
	a := NewArgs()

	// Default print
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Cmd configurations
	flag.StringVar(&a.InPath, "in", a.InPath, usgIn)
	flag.Var(&a.ExclPath, "excl", usgExcl)
	flag.StringVar(&a.OutPath, "out", a.OutPath, usgOut)
	flag.BoolVar(&a.Recursive, "r", a.Recursive, usgR)
	flag.StringVar(&a.PkgName, "pkg", a.PkgName, usgPackage)
	flag.Var(&a.InFile, "from", usgFrom)
	flag.StringVar(&a.OutFileName, "name", a.OutFileName, usgName)

	flag.Parse()

	return a
}
