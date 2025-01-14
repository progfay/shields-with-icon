package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/progfay/colorcontrast"
)

type Shield struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}

func (s Shield) String() string {
	return fmt.Sprintf("![%s](%s)", strings.ReplaceAll(strings.ReplaceAll(s.Title, "[", "\\["), "]", "\\]"), s.Src)
}

func generateShieldSrc(icon Icon) (string, error) {
	color, err := hexToColor(icon.Hex)
	if err != nil {
		return "", err
	}

	var foreground, background string

	if colorcontrast.CalcContrastRatio(white, *color) >= 2.5 {
		foreground = colorToHex(white)
		background = colorToHex(*color)
	} else {
		foreground = colorToHex(*color)
		background = colorToHex(black)
	}

	return fmt.Sprintf("https://img.shields.io/badge/%s-%s?style=for-the-badge&logo=%s&logoColor=%s",
		url.PathEscape(icon.Title),
		url.PathEscape(background),
		url.QueryEscape(icon.Title),
		url.QueryEscape(foreground),
	), nil
}

func IconToShield(icon Icon) (*Shield, error) {
	src, err := generateShieldSrc(icon)
	if err != nil {
		return nil, err
	}

	return &Shield{
		Title: icon.Title,
		Src:   src,
	}, nil
}
