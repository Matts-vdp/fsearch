package fslib

import (
	"fmt"
	"os"
)

type DirEntry interface {
	Name() string
	IsDir() bool
}

// Creates a validating function that is used to know wich files are valid
// Using this will only validate files who's name == str
func CreateStringComp(str string) func(DirEntry) bool {
	return func(file DirEntry) bool {
		return file.Name() == str
	}
}

// Blocks a goroutine until done has received numStreams messages
// Used to block untill all incoming streams are done
func waitTillDone(numStreams int, done chan bool) {
	for i := 0; i < numStreams; i++ {
		<-done
	}
}

// Creates a new goroutine for every directory and waits utill all routines are done
func handleDirs(dirs []string, out chan string, validator func(DirEntry) bool) {
	done := make(chan bool)
	for _, dir := range dirs {
		go searchForFile(dir, out, done, validator)
	}
	waitTillDone(len(dirs), done)
}

// Returns a channel on wich valid results are returned when found
// The channel is closed when all files are checked
func SearchFor(path string, validator func(DirEntry) bool) chan string {
	out := make(chan string)
	go func() {
		done := make(chan bool)
		go searchForFile(path, out, done, validator)
		for {
			select {
			case fpath := <-out:
				fmt.Println(fpath)
			case <-done:
				close(out)
				return
			}
		}
	}()
	return out
}

// Searches for a file in a directory and all its sub directories and returns the valid results on the out channel
// The validator will be used to check if a file is valid
// When al underlying files and folders are checked a response is send on the done channel
func searchForFile(path string, out chan string, done chan bool, validator func(DirEntry) bool) {
	files, err := searchFolder(path)
	if err != nil {
		done <- true
		return
	}
	dirs := make([]string, 0)
	for _, file := range files {
		fpath := path + "/" + file.Name()
		if file.IsDir() {
			dirs = append(dirs, fpath)
			continue
		} else {
			if validator(file) {
				out <- fpath
			}
		}
	}
	handleDirs(dirs, out, validator)
	done <- true
}

// Gets all files and folders in a folder
func searchFolder(path string) ([]os.DirEntry, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := f.ReadDir(-1)
	if err != nil {
		return nil, err
	}
	return files, nil
}
