package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("C:/ test")

	if os.IsNotExist(err)
}