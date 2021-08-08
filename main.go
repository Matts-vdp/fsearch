package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Matts-vdp/fsearch/fslib"
)

// returns all files in path with the name == filename
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: fsearch <path> <filename>")
		return
	}
	path := os.Args[1]
	name := os.Args[2]

	startTime := time.Now()
	val := fslib.CreateStringComp(name)
	out := fslib.SearchFor(path, val)
	for fpath := range out {
		fmt.Println(fpath)
	}
	fmt.Println("The program took", time.Since(startTime))
}
