package main

import (
	"github.com/smhouse/net/nm"
	"log"
)

func main() {
	devices, err := nm.GetDevices()
	if err != nil {
		log.Fatalln(err)
	}

	for _, d := range *devices {
		log.Printf("%+v\n", d)
	}

}
