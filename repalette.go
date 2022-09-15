package main

import (
    "image"
    "image/color"

    "github.com/lucasb-eyer/go-colorful"

    _ "image/jpeg"
    _ "image/png"
)

func (plt Palette) Blend(c color.Color, coeff float64) color.Color {
    cf, ok := colorful.MakeColor(c)
    if !ok {
        return c
    }

    minDistance := 1.0
    nearest := colorful.Color{}
    for _, pc := range plt {
        distance := pc.DistanceCIEDE2000(cf)
        if distance < minDistance {
            minDistance = distance
            nearest = pc
        }
    }

    return nearest.BlendLuv(cf, coeff)
}

func Repalette(img image.Image, plt Palette, coeff float64) image.Image {
    bounds := img.Bounds()
    dst := image.NewRGBA(bounds)

    for x := bounds.Min.X; x < bounds.Max.X; x++ {
        go func(x int) {
            for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
                nearest := plt.Blend(img.At(x, y), coeff)
                dst.Set(x, y, nearest)
            }
        }(x)
    }

    return dst
}
