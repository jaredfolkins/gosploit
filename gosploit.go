package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	tmp := os.TempDir()
	td, err := ioutil.ReadDir(tmp)
	if err != nil {
		panic(err)
	}

	location := "/"
	err = filepath.Walk(location, func(path string, _ os.FileInfo, _ error) error {
		numScanned++
		log.Println(path)
	})

	/*
		for _, t := range td {
			log.Println(t.Name())
			log.Println(t.ModTime())
		}
	*/

	/*
		log.Println(os.Hostname())
		g, err := os.Getgroups()
		if err != nil {
			panic(err)
		}

		log.Println(g)
		log.Println(os.Getpagesize())
		log.Println(os.Environ())
		log.Println(os.Environ())
	*/
}
