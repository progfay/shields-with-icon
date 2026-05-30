package main

import (
	"fmt"
	"html"
	"net/url"
	"strings"

	"github.com/progfay/colorcontrast"
)

type Shield struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}

// Markdown renders the shield as a Markdown image: ![title](src).
func (s Shield) Markdown() string {
	return fmt.Sprintf("![%s](%s)", strings.ReplaceAll(strings.ReplaceAll(s.Title, "[", "\\["), "]", "\\]"), s.Src)
}

// HTML renders the shield as a <picture> wrapping an <img> with explicit
// width/height so the README reserves space before the badge loads, avoiding
// layout shift. The width is computed offline to match Shields.io's
// for-the-badge renderer.
func (s Shield) HTML() string {
	return fmt.Sprintf(`<picture><img alt="%s" src="%s" width="%d" height="%d"></picture>`,
		html.EscapeString(s.Title),
		html.EscapeString(s.Src),
		forTheBadgeWidth(s.Title),
		badgeHeight,
	)
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

	// Escape every "-" as "--" so Shields.io treats the title as a single
	// message instead of splitting it into label/message on dashes.
	return fmt.Sprintf("https://img.shields.io/badge/%s-%s?style=for-the-badge&logo=%s&logoColor=%s",
		url.PathEscape(strings.ReplaceAll(icon.Title, "-", "--")),
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
