package main

import (
    "bufio"
    "embed"
    "fmt"
    "io"
    "os"

    "github.com/lucasb-eyer/go-colorful"
)

//go:embed palettes/*
var palettes embed.FS

type Palette []colorful.Color

func LoadPaletteByName(name string) (Palette, error) {
    file, err := palettes.Open(fmt.Sprintf("palettes/%s", name))
    if err != nil {
        return nil, err
    }

    return LoadPaletteFromReader(file)
}

func LoadPaletteFromFile(filename string) (Palette, error) {
    file, err := os.Open(filename)
    if err != nil {
	    return nil, err
    }
    defer file.Close()

    return LoadPaletteFromReader(file)
}

func LoadPaletteFromReader(r io.Reader) (Palette, error) {
    var palette Palette

    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        c, err := colorful.Hex(scanner.Text())
        if err != nil {
            return nil, fmt.Errorf("unable to parse '%s' color: %w", err)
        }
        palette = append(palette, c)
    }
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("unable to read content: %w", err)
    }

    return palette, nil
}
