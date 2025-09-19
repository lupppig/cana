package main

import (
	"log"

	"github.com/lupppig/cana"
)

func main() {
	c := cana.Canabis("localhost:8080")

	if err := c.ServeCana(); err != nil {
		log.Fatal(err)
	}

}
