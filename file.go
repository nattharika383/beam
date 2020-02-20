package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walk(path string, info os.FileInfo, err error) error {
	fmt.Println(pant)
	return nil
}

func main(){
	filepath.Walk(".", walkFn)
}