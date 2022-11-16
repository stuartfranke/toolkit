package main

import (
	"github.com/stuartfranke/toolkit"
	"log"
)

func main() {
	toSlug := "hello world 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}