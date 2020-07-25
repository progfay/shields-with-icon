package main

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
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

	B := uint8(rgb & 0xFF)
	rgb >>= 8
	G := uint8(rgb & 0xFF)
	rgb >>= 8
	R := uint8(rgb & 0xFF)

	return &color.RGBA{R, G, B, 0xFF}, nil
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%02X%02X%02X", r>>8, g>>8, b>>8)
}
