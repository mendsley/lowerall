package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

func processDir(p string, w *sync.WaitGroup) {
	defer w.Done()

	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal("Failed to read directory: ", err)
	}

	for _, f := range files {
		name := path.Join(p, f.Name())
		lname := strings.ToLower(name)
		if name != lname {
			err := os.Rename(name, lname)
			if err != nil {
				log.Fatalf("Failed to rename %s to %s: %v", name, lname, err)
			}
		}
		if f.IsDir() {
			w.Add(1)
			go processDir(lname, w)
		}
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)
	processDir(".", &w)
	w.Wait()
}
