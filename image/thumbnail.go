package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"golang.org/x/image/draw"
)

/*
生成缩略图
输入: 宽或高为0时自动按比例缩放
*/
func Thumbnail(r io.Reader, w io.Writer, mimetype string, width int, height int) error {
	var src image.Image
	var err error

	switch mimetype {
	case "image/jpeg":
		src, err = jpeg.Decode(r)
	case "image/png":
		src, err = png.Decode(r)
	}

	if err != nil {
		return err
	}

	ratio := (float64)(src.Bounds().Max.Y) / (float64)(src.Bounds().Max.X)
	if width <= 0 {
		if height <= 0 {
			return fmt.Errorf("both width and height must be positive")
		}
		width = int(math.Round(float64(height) / ratio))
	} else if height <= 0 {
		height = int(math.Round(float64(width) * ratio))
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = jpeg.Encode(w, dst, nil)
	if err != nil {
		return err
	}

	return nil
}
