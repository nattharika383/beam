package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

