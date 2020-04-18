package utilities

import (
	"fmt"
	"io/ioutil"
)

func Readfile(filename string) {
	fmt.Println("filepath is", filename)

	data, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error reading file", error)
		return
	}
	fmt.Println(string(data))
}
 