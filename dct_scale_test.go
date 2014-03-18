package main

import (
	"testing"
        "fmt"
        "github.com/dsachitano/go_image/jpeg"
        "os"
)

// utility for opening a file
func openImage(path string, t *testing.T) (file *os.File) {

        f, err := os.Open(path)
        if err != nil {
                fmt.Printf("%v\n", err)
        }
	if f == nil {
		t.Fail()
	}

	return f
}

func TestDctScale(t *testing.T) {

	f := openImage("testdata/redwoods.jpg", t)

        dctScaledimg, _ := jpeg.DCTScale(f)

	if dctScaledimg == nil {
		t.FailNow()
	}

	// get size
        x := dctScaledimg.Bounds().Max.X
        y := dctScaledimg.Bounds().Max.Y

	// make sure it's non-zero sized
	if x <= 0 || y <= 0 {
		t.Fail()
	}

        f.Close()
}

func noTestProgScale(t *testing.T) {
	f := openImage("testdata/redwoods.progressive.jpg", t)

	numScans := 1

        progScaledimg, _ := jpeg.ProgScale(f, numScans)

	if progScaledimg == nil {
		t.FailNow()
	}

	// get size
        x := progScaledimg.Bounds().Max.X
        y := progScaledimg.Bounds().Max.Y

	// make sure it's non-zero sized
	if x <= 0 || y <= 0 {
		t.Fail()
	}

        f.Close()
}
