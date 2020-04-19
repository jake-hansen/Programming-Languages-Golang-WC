package utilities

import (
	"fmt"
	"io/ioutil"
)

func Readfile(filename string) []byte{
	data, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error reading file", error)
		return nil
	}
	return data
}
 