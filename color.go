package main

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
)

const (
	rMask = 0xFF0000
	gMask = 0x00FF00
	bMask = 0x0000FF
)

var (
	colorCodeRegExp = regexp.MustCompile(`^#?[0-9a-fA-F]{6}$`)
)

func hexToColor(colorCode string) (*color.RGBA, error) {
	if !colorCodeRegExp.Match([]byte(colorCode)) {
		return nil, fmt.Errorf("invalid color code")
	}

	code := strings.TrimLeft(colorCode, "#")
	rgb, err := strconv.ParseUint(code, 16, 24)
	if err != nil {
		return nil, err
	}

	c := color.RGBA{
		R: uint8(rgb & rMask >> 16),
		G: uint8(rgb & gMask >> 8),
		B: uint8(rgb & bMask >> 0),
		A: 0xFF,
	}

	return &c, nil
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%02X%02X%02X", r>>8, g>>8, b>>8)
}
