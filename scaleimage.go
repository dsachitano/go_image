package main

import (
	"fmt"
	"github.com/dsachitano/go_image/jpeg"
	"os"
	"flag"
	"math/rand"
)

func main() {

	// parse input
	inputFileName := flag.String("input", "ERROR", "path of file to scale")
	flag.Parse()

	f, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	dctScaledimg, _ := jpeg.DCTScale(f)

	// DS: scaledImage should be full of pixels.  write it out to disk
        // as a JPEG
	x := dctScaledimg.Bounds().Max.X
	y := dctScaledimg.Bounds().Max.Y

        outputFileName := fmt.Sprintf("scaled_%d_%d_rand_%d.jpg", x, y, rand.Int())
        outputFile, _ := os.Create(outputFileName)
        jpeg.Encode(outputFile, dctScaledimg, &jpeg.Options{jpeg.DefaultQuality})

	f.Close()
}
