package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	bs := make([]byte, 5)
	readFile(args[0], bs)
	fmt.Println(string(bs))
}

func readFile(filename string, bs []byte) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening: ", err)
		os.Exit(1)
	} 
	fp.Read(bs)
}