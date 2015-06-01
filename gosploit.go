package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Hostname())
	g, err := os.Getgroups()
	if err != nil {
		panic(err)
	}
	log.Println(g)
	log.Println(os.Getpagesize())
	log.Println(os.Environ())
	os.Clearenv()
	log.Println(os.Environ())
}
