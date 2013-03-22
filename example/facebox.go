// +build !appengine

// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"github.com/suapapa/go_face"
	"image"
	_ "image/jpeg"
	"image/png"
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

	fmt.Println("Detecting faces...")
	faces := face.Detect(img)
	if len(faces) == 0 {
		fmt.Println("Failed to find face")
		os.Exit(0)
	}

	fmt.Println("Found", len(faces))
	bounds := img.Bounds()
	eyeMaskImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			eyeMaskImg.Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x, y))
		}
	}

	gc := draw2d.NewGraphicContext(eyeMaskImg)
	gc.SetLineWidth(5)
	for i, f := range faces {
		fmt.Println(i, f)
		r := f.Rect()
		gc.MoveTo(float64(r.Min.X), float64(r.Min.Y))
		gc.LineTo(float64(r.Min.X), float64(r.Max.Y))
		gc.LineTo(float64(r.Max.X), float64(r.Max.Y))
		gc.LineTo(float64(r.Max.X), float64(r.Min.Y))
		gc.LineTo(float64(r.Min.X), float64(r.Min.Y))
		gc.Stroke()
	}

	saveToPngFile("result.png", eyeMaskImg)
}

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Fatal(err)
	}
	err = b.Flush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}
