package main

import (
	"fmt"
	"github.com/liangrog/go-j2s/gen"
	"os"
)

func main() {
	if err := gen.Proc(); err != nil {
		fmt.Fprintf(os.Stderr, "go-j2s: %v\n", err)
		os.Exit(1)
	}
}
