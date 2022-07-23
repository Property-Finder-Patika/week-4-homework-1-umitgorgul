package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func visit(path string, di fs.DirEntry, err error) error {
	a := len(path)
	if a > 4 {
		b := path[a-3:]
		if b == ".go" {
			fmt.Println(path[:])
		}
	}
	return nil
}

func main() {
	flag.Parse()
	//root := flag.Arg(0)
	err := filepath.WalkDir(".", visit)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
