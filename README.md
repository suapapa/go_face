# go_face : Face detector for Go

Package to find coordinates of faces from an image.
Docs are available at [godoc.org][1].

It's cgo binding of c library, neven.
The neven is backend of Android's [FaceDetector][2].


## Install

    $ go get github.com/suapapa/go_face

You should do following steps once to install neven.

You need `scons` to build and install the neven library.
Install it from package manager of your OS.

And, go to `go_face`'s source directory and run;

    $ cd $GOPATH/src/github.com/suapapa/go_face
    $ ./install_neven.sh

It will ask root password to install the library under `/usr/local`.


## Example

Check out `example` folder.

[1]:http://godoc.org/github.com/suapapa/go_face
[2]:http://developer.android.com/reference/android/media/FaceDetector.Face.html
