// Copyright 2013, Homin Lee All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package face

// #cgo LDFLAGS: -lneven -lstdc++
// #include <neven.h>
// neven_env_t *neven_create_prec(int w, int h, int max_faces)
// {
//	return neven_create(w, h, max_faces, NEVEN_INITDATA_PREC);
// }
import "C"
import "unsafe"

type nevenContext struct {
	ctx  *C.neven_env_t
	w, h int
}

func newNevenContext(w, h int, maxFaces int) *nevenContext {
	ctx := C.neven_create_prec(C.int(w), C.int(h), C.int(maxFaces))
	if ctx == nil {
		return nil
	}

	return &nevenContext{
		ctx: ctx,
		w:   w,
		h:   h,
	}
}

func (n *nevenContext) detect(buffer []byte) int {
	return int(C.neven_detect(n.ctx, unsafe.Pointer(&buffer[0])))
}

func (n *nevenContext) getFace(i int) *Face {
	var face C.neven_face_t
	C.neven_get_face(n.ctx, &face, C.int(i))

	return &Face{
		LeftEye: Point{
			X: float32(face.lefteye.x),
			Y: float32(face.lefteye.y),
		},
		RightEye: Point{
			X: float32(face.righteye.x),
			Y: float32(face.righteye.y),
		},
		Confidence: float32(face.confidence),
	}
}

func (n *nevenContext) destroy() {
	C.neven_destroy(n.ctx)
}
