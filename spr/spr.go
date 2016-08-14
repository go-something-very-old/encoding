package spr

import (
	"fmt"
	"image"

	bin "github.com/go-otserv/encoding/binary"
)

// File is wrapper for reading .spr file
type File struct {
	bin.File
	Signature       uint32
	SpritesCount    uint32
	SpriteOffset    int
	SpriteDataSize  int
	HasAplhaChannel bool
}

// Open opens given file for reading
func Open(path string, hasAplhaChannel bool) (*File, error) {
	fh, err := bin.OpenFile(path)
	if err != nil {
		return nil, err
	}
	sprfh := &File{*fh, 0, 0, 0, 0, hasAplhaChannel}

	sprfh.Signature, err = sprfh.UInt32()
	if err != nil {
		return sprfh, err
	}
	sprfh.SpritesCount, err = sprfh.UInt32()
	if err != nil {
		return sprfh, err
	}
	sprfh.SpriteOffset = 8
	sprfh.SpriteDataSize = 32 * 32 * 4

	return sprfh, nil
}

// GetSprite parses .spr file to extract sprite of given id
func (sprfh *File) GetSprite(id int) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	if id == 0 {
		return img, nil
	}
	sprfh.Seek(int64(((id-1)*4)+sprfh.SpriteOffset), 0)

	sprAddress, err := sprfh.UInt32()
	if err != nil {
		return nil, err
	}
	if sprAddress == 0 {
		return img, nil
	}

	sprfh.Seek(int64(sprAddress), 0)

	// skip color key - 3 * uint8
	sprfh.Discard(3)

	pixelDataSize, err := sprfh.UInt16()
	if err != nil {
		return nil, err
	}

	var writePos int
	var readPos uint16
	for readPos < pixelDataSize && writePos < sprfh.SpriteDataSize {
		transparentPixels, err := sprfh.UInt16()
		if err != nil {
			return nil, err
		}
		coloredPixels, err := sprfh.UInt16()
		if err != nil {
			return nil, err
		}

		var i uint16
		for i < transparentPixels && writePos < sprfh.SpriteDataSize {
			img.Pix[writePos+0] = 0
			img.Pix[writePos+1] = 0
			img.Pix[writePos+2] = 0
			img.Pix[writePos+3] = 0
			writePos += 4
			i++
		}

		i = 0
		for i < coloredPixels && writePos < sprfh.SpriteDataSize {
			var r, g, b, a uint8
			r, err = sprfh.UInt8()
			if err != nil {
				return nil, err
			}
			g, err = sprfh.UInt8()
			if err != nil {
				return nil, err
			}
			b, err = sprfh.UInt8()
			if err != nil {
				return nil, err
			}
			if sprfh.HasAplhaChannel {
				a, err = sprfh.UInt8()
				if err != nil {
					return nil, err
				}
			} else {
				a = 255
			}
			img.Pix[writePos+0] = r
			img.Pix[writePos+1] = g
			img.Pix[writePos+2] = b
			img.Pix[writePos+3] = a
			writePos += 4
			i++
		}

		readPos += 4 + (3 * coloredPixels)
	}

	// fill remaining pixels with alpha
	for writePos < sprfh.SpriteDataSize {
		img.Pix[writePos+0] = 0
		img.Pix[writePos+1] = 0
		img.Pix[writePos+2] = 0
		img.Pix[writePos+3] = 0
		writePos += 4
	}

	return img, nil
}

// CombineSprites combines given images into one bigger image. For example one
// 64x64 px image from four 32x32 px images.
// Order of images in slice is important, first sprite will be placed in bottom
// right tile, next - one tile to the left and so on, always starting new row
// from rightmost tile.
// +---+---+
// | 4 | 3 | <- Layout produced by passing slice of images
// +---+---+    [1, 2, 3, 4].
// | 2 | 1 |
// +---+---+
func CombineSprites(width, height int, sprites []*image.RGBA) (*image.RGBA, error) {
	spritesNum := len(sprites)
	combinedTiles := width * height
	if spritesNum != combinedTiles {
		return nil, fmt.Errorf(
			"Expected slice of size %d (%d[width] * %d[height] = %d) but slice of size %d given",
			combinedTiles, width, height, combinedTiles, spritesNum,
		)
	}
	xMax, yMax := width-1, height-1
	combined := image.NewRGBA(image.Rect(0, 0, 32*width, 32*height))
	x, y := xMax, yMax

	nextCords := func(x, y int) (int, int) {
		if x > 0 {
			x--
			return x, y
		}
		if y > 0 {
			x = xMax
			y--
			return x, y
		}
		return xMax, yMax
	}

	for _, sprite := range sprites {
		for xp := 0; xp < 32; xp++ {
			for yp := 0; yp < 32; yp++ {
				combined.Set((32*x)+xp, (32*y)+yp, sprite.At(xp, yp))
			}
		}
		x, y = nextCords(x, y)
	}

	return combined, nil
}
