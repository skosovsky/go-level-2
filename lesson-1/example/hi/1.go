package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	dirPath = flag.String("dirPath", "", "Path to dir with files")
	substrRequired = "Hello World!"
)

func main() {
	type HelloWorldStats struct {
		Count     int
		FileNames []string
	}

	flag.Parse()
	if *dirPath == "" {
		panic("dirPath shouldn't be empty")
	}

	stats := &HelloWorldStats{
		Count: 0,
		FileNames: make([]string, 0),
	}
	allFiles, err := ioutil.ReadDir(*dirPath)
	if err != nil{
		fmt.Printf("Error: %s", err)

	}
	for _, f := range(allFiles){
		file, err := os.Open(*dirPath + "/" + f.Name())
		defer file.Close()

		if err != nil{
			fmt.Printf("Error: %s", err)
		}

	}
}