package necessity

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func MustDecodeImage(r io.Reader) (image.Image, string) {
	x, s, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	return x, s
}

func MustDecodePNG(r io.Reader) image.Image {
	x, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	return x
}

func MustDecodeJPEG(r io.Reader) image.Image {
	x, err := jpeg.Decode(r)
	if err != nil {
		panic(err)
	}

	return x
}
