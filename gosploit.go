package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	var err error

	location := "/"
	err = filepath.Walk(location, func(path string, _ os.FileInfo, _ error) error {
		log.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	/*
		tmp := os.TempDir()
		td, err := ioutil.ReadDir(tmp)
		if err != nil {
			panic(err)
		}

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
