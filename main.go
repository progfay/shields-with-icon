package main

import (
	"fmt"
	"log"
	"net/url"

	c "github.com/progfay/shields-with-icon/color"
	i "github.com/progfay/shields-with-icon/icon"
)

var (
	white = c.Color{
		Red:   1,
		Green: 1,
		Blue:  1,
		Code:  "FFFFFF",
	}
	black = c.Color{
		Red:   34.0 / 255.0,
		Green: 34.0 / 255.0,
		Blue:  34.0 / 255.0,
		Code:  "222222",
	}
)

func FormatShield(icon i.Icon) (string, error) {
	color, err := c.NewColor(icon.Hex)
	if err != nil {
		return "", nil
	}

	var foreground, background c.Color

	if c.GetContrastRatio(white, *color) >= 7.0 {
		foreground = white
		background = *color
	} else {
		foreground = *color
		background = black
	}

	return fmt.Sprintf("[![%v](http://img.shields.io/badge/%s-%s?style=for-the-badge&logo=%s&logoColor=%s)](%s)",
		icon.Title,
		url.QueryEscape(icon.Title),
		url.QueryEscape(background.Code),
		url.QueryEscape(icon.Title),
		url.QueryEscape(foreground.Code),
		icon.Source,
	), nil
}

func main() {
	icons, err := i.GetIcons()
	if err != nil {
		log.Panicln(err)
	}

	for _, icon := range icons {
		shield, err := FormatShield(icon)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(shield)
	}
}
