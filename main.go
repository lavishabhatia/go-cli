package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

var help = flag.Bool("help", false, "Show help")
var isFileFlag = false
var nameFlag = "Hello There!"

func main() {
	flag.BoolVar(&isFileFlag, "file", false, "A boolean flag")
	flag.StringVar(&nameFlag, "name", "Hello There!", "A string flag")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if isFileFlag {
		_, err := os.Create(nameFlag)
		if err != nil {
			fmt.Println("Error creating file")
			os.Exit(0)
		} else {
			fmt.Printf("%v file created", nameFlag)
			os.Exit(0)
		}
	}

	isErr := os.Mkdir(nameFlag, fs.ModePerm)
	if isErr != nil {
		fmt.Println("error Creating directory")
		os.Exit(0)
	} else {
		fmt.Printf("%v Directory created successfully", nameFlag)
	}

}
