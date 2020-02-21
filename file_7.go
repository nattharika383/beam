package file

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("output.txt")
	if err != nil {
		return
}

defer file.Close()

file.WriteString("Hello")
file.WriteString("output.txt")