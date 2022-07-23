package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func visit(path string, di fs.DirEntry, err error) error {
	a := len(path) // get length of path abc.txt = 7 , abcdf.go = 8 vb vb vb 
	if a > 3 { // look if path is length bigger than 3 because we need the show paths with .go  
		b := path[a-3:] // we slicing path for searching ".go:" 
		if b == ".go" {  // if it find a path which has ".go" print it
			fmt.Println(path[:])
		}
	}
	return nil
}

func main() {
	flag.Parse()
	//root := flag.Arg(0) // for searching in specific location
	root := "." // for searching in code is location
	err := filepath.WalkDir(root, visit)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
