// +build !appengine

// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/suapapa/go_face"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s jpeg_file", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()

	// Change Image to Gray
	grayImg := image.NewGray(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{bounds.Dx(), bounds.Dy()},
	})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x, y))
		}
	}

	faces := face.Detect(grayImg)

	fmt.Println("Found", len(faces))
	for i, f := range faces {
		fmt.Println(i, f)
	}
}
