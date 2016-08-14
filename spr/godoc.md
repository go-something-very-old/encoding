# Package `spr`
## Overview

## Index

* Functions
  * [func CombineSprites(width, height int, sprites []*image.RGBA) (*image.RGBA, error)](#CombineSprites)

* Types
  * [File](#File)
	 * [func Open(path string, hasAplhaChannel bool) (*File, error)](#Open)
	 * [func (sprfh *File) GetSprite(id int) (*image.RGBA, error)](#File-GetSprite)

## Functions

### func <a href="https://github.com/go-otserv/encoding/blob/master/spr/spr.go#L147" name="CombineSprites">CombineSprites</a> [¶](#CombineSprites)
```go
func CombineSprites(width, height int, sprites []*image.RGBA) (*image.RGBA, error)
```
CombineSprites combines given images into one bigger image. For example one
64x64 px image from four 32x32 px images.

Order of images in slice is important, first sprite will be placed in bottom
right tile, next - one tile to the left and so on, always starting new row
from rightmost tile.
<pre>
+---+---+
| 4 | 3 | <- Layout produced by passing slice of images
+---+---+    [1, 2, 3, 4].
| 2 | 1 |
+---+---+
</pre>

## Types

### type <a href="https://github.com/go-otserv/encoding/blob/master/spr/spr.go#L11" name="File">File</a> [¶](#File)
```go
type File struct {
	bin.BinaryFile
	Signature       uint32
	SpritesCount    uint32
	SpriteOffset    int
	SpriteDataSize  int
	HasAplhaChannel bool
}
```
File is wrapper for reading .spr file  

#### func <a href="https://github.com/go-otserv/encoding/blob/master/spr/spr.go#L21" name="Open">Open</a> [¶](#Open)
```go
func Open(path string, hasAplhaChannel bool) (*File, error)
```
Open opens given file for reading

#### func <a href="https://github.com/go-otserv/encoding/blob/master/spr/spr.go#L43" name="File-GetSprite">GetSprite</a> [¶](#File-GetSprite)
```go
func (sprfh *File) GetSprite(id int) (*image.RGBA, error)
```
GetSprite parses .spr file to extract sprite of given id

***
_Last updated 12 Aug 2016_
