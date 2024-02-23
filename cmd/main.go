package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mua/deasciifier"
)

func deasciifyString(input string) string {
	return deasciifier.Deasciify(input)
}

func deasciifyFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	stat, _ := file.Stat()
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(bs)
	return deasciifyString(content)
}

func main() {
	filePtr := flag.String("f", "", "file to deasciify")
	stringPtr := flag.String("s", "", "string to deasciify")
	flag.Parse()

	if *filePtr != "" {
		fmt.Println(deasciifyFile(*filePtr))
		return
	}
	if *stringPtr != "" {
		fmt.Println(deasciifyString(*stringPtr))
		return
	}

	fmt.Println("Please provide a file or a string to deasciify")
}
