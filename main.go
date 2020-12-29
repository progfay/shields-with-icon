package main

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/progfay/colorcontrast"
)

var (
	white = color.White
	black = color.Gray{Y: 34}
)

func formatShield(icon Icon) (string, error) {
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

	return fmt.Sprintf("![%v](https://img.shields.io/static/v1?style=for-the-badge&message=%s&color=%v&logo=%s&logoColor=%s&label=)",
		strings.ReplaceAll(strings.ReplaceAll(icon.Title, "[", "\\["), "]", "\\]"),
		url.QueryEscape(icon.Title),
		url.QueryEscape(background),
		url.QueryEscape(icon.Title),
		url.QueryEscape(foreground),
	), nil
}

func main() {
	icons, err := getIcons()
	if err != nil {
		log.Panicln(err)
	}

	for _, icon := range icons {
		shield, err := formatShield(icon)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Fprintln(os.Stdout, shield)
		fmt.Fprintf(os.Stderr, "## %[1]s\n```markdown\n%[1]s\n```\n", shield)
	}
}
