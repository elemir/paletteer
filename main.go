package main

import (
    "os"
    "path/filepath"
    "fmt"
    "image"
    "image/png"
    "log"
    "strings"

	"github.com/urfave/cli/v2"

    _ "image/jpeg"
)

func Run(src, dst, palette string, blend float64) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    // TODO: save in same format
    srcImg, _, err := image.Decode(srcFile)
    if err != nil {
    }

    plt, err := LoadPaletteByName(palette)
    if err != nil {
        return err
    }

    dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    dstImg := Repalette(srcImg, plt, blend)

    return png.Encode(dstFile, dstImg)
}

func main() {
	app := &cli.App{
		Name:  "paletteer",
		Description: "make image more compatible with a specific color scheme",
        ArgsUsage: "filename",
        Authors: []*cli.Author{
            {Name: "Evgenii Omelchenko", Email: "elemir90@gmail.com"},
        },
        Flags: []cli.Flag{
            &cli.StringFlag{Name: "palette", Value: "gruvbox", Usage: "the palette to use"},
            &cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "output picture name (default <filename>-<palette>.png)"},
            &cli.Float64Flag{Name: "blend", Aliases: []string{"b"}, Usage: "coefficient of blending (default zero, so picture will be quantize)"},
        },
        Before: func(c *cli.Context) error {
            if c.Args().Len() != 1 {
                return fmt.Errorf("require exactly one argument")
            }

            return nil
        },
		Action: func (c *cli.Context) error {
            src := c.Args().First()
            palette := c.String("palette")
            dst := c.String("output")
            blend := c.Float64("blend")

            if dst == "" {
                basename := filepath.Base(src)
                name := strings.TrimSuffix(basename, filepath.Ext(basename))
                dst = fmt.Sprintf("%s-%s.png", name, palette)
            }

            return Run(src, dst, palette, blend)
        },
	}

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
