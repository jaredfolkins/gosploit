package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Hostname())
	log.Println(os.Environ())
	log.Println(os.Clearenv())
	log.Println(os.Environ())
}
