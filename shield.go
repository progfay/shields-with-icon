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

// Markdown renders the shield as a Markdown image: ![title](src).
func (s Shield) Markdown() string {
	return fmt.Sprintf("![%s](%s)", strings.ReplaceAll(strings.ReplaceAll(s.Title, "[", "\\["), "]", "\\]"), s.Src)
}

// HTML renders the shield as a single <img> with explicit width/height so the
// README reserves space before the badge loads, avoiding layout shift. The
// width is computed offline to match Shields.io's for-the-badge renderer.
//
// The caller wraps the whole gallery in one <picture> element (see main.go) so
// GitHub does not apply its default border-radius to each badge. We emit one
// <picture> for the entire document — instead of one per <img> — to keep the
// README under GitHub's 512KB front-page render limit. For the same reason src
// is written without HTML-escaping (it is already URL-encoded by
// generateShieldSrc, so raw "&" is valid and 1 byte instead of "&amp;") and the
// alt attribute is omitted: the badge text is the title, so it adds bytes
// without conveying anything new.
func (s Shield) HTML() string {
	return fmt.Sprintf(`<img src="%s" width="%d" height="%d">`,
		s.Src,
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
