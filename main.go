package main

import (
	"fmt"
	"log"
	"net/url"

	i "github.com/progfay/shields-with-icon/icon"
)

func FormatShield(icon i.Icon) string {
	return fmt.Sprintf("[![%v](http://img.shields.io/badge/%s-%s?style=for-the-badge&logo=%s&logoColor=FFFFFF)](%s)",
	icon.Title,
		url.QueryEscape(icon.Title),
		url.QueryEscape(icon.Hex),
		url.QueryEscape(icon.Title),
		icon.Source,
	)
}

func main() {
	icons, err := i.GetIcons()
	if err != nil {
		log.Panicln(err)
	}

	for _, icon := range icons {
		fmt.Println(FormatShield(icon))
	}
}
