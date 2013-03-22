package face

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"testing"
)

var (
	img    image.Image
	bounds image.Rectangle

	dbg func(format string, v ...interface{})
)

func init() {
	file, err := os.Open(os.Getenv("TEST_JPG"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	bounds = img.Bounds()

	dbg = log.Printf
}

func TestDetectYCbCrImage(t *testing.T) {
	t.Logf("Detecting faces...")
	faces := Detect(img)
	if faces == nil {
		t.Fatal("Error on detecting face!")
	}

	t.Logf("Found %d", len(faces))
	for i, f := range faces {
		t.Logf("%d, %v", i, f)
	}
}

func TestDetectGrayImage(t *testing.T) {
	t.Logf("Converting to Gray...")
	grayImg := image.NewGray(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{bounds.Dx(), bounds.Dy()},
	})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x, y))
		}
	}

	t.Logf("Detecting faces...")
	faces := Detect(grayImg)
	if faces == nil {
		t.Fatal("Error on detecting face!")
	}

	t.Logf("Found %d", len(faces))
	for i, f := range faces {
		t.Logf("%d, %v", i, f)
	}
}

func TestDetectRGBAImage(t *testing.T) {
	t.Logf("Converting to RGBA...")
	rgbaImg := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{bounds.Dx(), bounds.Dy()},
	})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x, y))
		}
	}

	t.Logf("Detecting faces...")
	faces := Detect(rgbaImg)
	if faces == nil {
		t.Fatal("Error on detecting face!")
	}

	t.Logf("Found %d", len(faces))
	for i, f := range faces {
		t.Logf("%d, %v", i, f)
	}
}
