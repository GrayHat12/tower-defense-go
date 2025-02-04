package sprites

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type CropOptions struct {
	Rectangle image.Rectangle
}

func cropSprites(sprite *ebiten.Image, crop *CropOptions) *ebiten.Image {
	// return sprite
	if crop == nil {
		crop = &CropOptions{
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 100 - 64,
				},
				Max: image.Point{
					X: 64,
					Y: 100,
				},
			},
		}
	}
	return ebiten.NewImageFromImage(sprite.SubImage(crop.Rectangle))
}
