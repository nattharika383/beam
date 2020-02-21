package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("C:/ test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("c:/ test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
