package main

import (
	"fmt"
	"github.com/dsachitano/go_image/jpeg"
	"os"
	"flag"
)

func main() {

	// parse input
	inputFileName := flag.String("input", "ERROR", "path of file to scale")
	flag.Parse()

	f, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	jpeg.Decode(f)

	f.Close()
}
