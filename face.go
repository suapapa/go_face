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

// Find faces from given image
// TODO: Current it only receive image in image.Gray type
func Detect(image *image.Gray) []*Face {
	w, h := image.Rect.Dx(), image.Rect.Dy()
	n := newNevenContext(w, h, 5)
	if n == nil {
		return nil
	}
	defer n.destroy()

	fn := n.detect(image.Pix)
	fs := make([]*Face, fn)
	for i := 0; i < fn; i++ {
		fs[i] = n.getFace(i)
	}

	return fs
}
