package data

import (
	"embed"
	"encoding/gob"
	"github.com/RoaringBitmap/roaring"
)

// data [][]byte "github.com/RoaringBitmap/roaring" Normal NormalEx Tsupai TsupaiEx
//go:embed bitmap
var FS embed.FS

var (
	Normal   roaring.Bitmap
	NormalEx roaring.Bitmap
	Tsupai   roaring.Bitmap
	TsupaiEx roaring.Bitmap
)

func init() {
	var bitmaps [][]byte
	fs, err := FS.Open("bitmap")
	if err != nil {
		panic(err)
	}
	if err := gob.NewDecoder(fs).Decode(&bitmaps); err != nil {
		panic(err)
	}
	if _, err := Normal.FromBuffer(bitmaps[0]); err != nil {
		panic(err)
	}
	if _, err := NormalEx.FromBuffer(bitmaps[1]); err != nil {
		panic(err)
	}
	if _, err := Tsupai.FromBuffer(bitmaps[2]); err != nil {
		panic(err)
	}
	if _, err := TsupaiEx.FromBuffer(bitmaps[3]); err != nil {
		panic(err)
	}
}
