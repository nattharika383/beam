package file

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("output.txt")
	if err != nil {
		
}