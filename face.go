// Copyright 2013, Homin Lee All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package face

import (
	"image"
)

type Point struct {
	X, Y float32
}

// type Face represents a face which is detected in an image
type Face struct {
	LeftEye, RightEye Point
	Confidence        float32
}

func (f *Face) Rect() image.Rectangle {
	eyeDist := f.RightEye.X - f.LeftEye.X
	midX := (f.RightEye.X + f.LeftEye.X) / 2
	midY := (f.RightEye.Y + f.LeftEye.Y) / 2

	return image.Rectangle{
		Min: image.Point{
			X: int(midX - eyeDist*1.5 + 0.5),
			Y: int(midY - eyeDist*2 + 0.5),
		},
		Max: image.Point{
			X: int(midX + eyeDist*1.5 + 0.5),
			Y: int(midY + eyeDist*2 + 0.5),
		},
	}
}

// Find faces from given image
func Detect(img interface{}) []*Face {
	var bwBuff []byte
	var bounds image.Rectangle
	switch i := img.(type) {
	case *image.Gray:
		// dbg("image.Gray...")
		bounds = i.Bounds()
		bwBuff = i.Pix
	case *image.YCbCr:
		// dbg("image.YCbCr...")
		bounds = i.Bounds()
		bwBuff = i.Y
	case image.Image:
		// dbg("Other image.Image, %T...", i)
		bounds = i.Bounds()
		grayImg := image.NewGray(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{bounds.Dx(), bounds.Dy()},
		})
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				grayImg.Set(x-bounds.Min.X, y-bounds.Min.Y, i.At(x, y))
			}
		}
		bwBuff = grayImg.Pix
	default:
		return nil
	}

	// dbg("Creating Context...")
	w, h := bounds.Dx(), bounds.Dy()
	n := newNevenContext(w, h, 5)
	if n == nil {
		return nil
	}
	defer n.destroy()

	// dbg("Detecting...")
	fn := n.detect(bwBuff)
	fs := make([]*Face, fn)
	for i := 0; i < fn; i++ {
		fs[i] = n.getFace(i)
	}

	return fs
}
