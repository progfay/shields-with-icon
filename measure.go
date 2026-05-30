package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"math"
	"sort"
	"strings"
	"unicode/utf16"
)

// Verdana "bold 10px" character width table from the anafanafo package.
// Each entry is [lower, upper, width]: every code point in [lower, upper] has
// the given advance width (px). Entries are sorted ascending by lower bound.
//
// ref. https://github.com/metabolize/anafanafo/blob/main/packages/anafanafo/data/verdana-10px-bold.json
//
//go:embed data/verdana-10px-bold.json
var verdanaBoldData []byte

// for-the-badge renderer constants.
//
// ref. https://github.com/badges/shields/blob/ee31d6164c01c32cb9016cba49d858f569b82b29/badge-maker/lib/badge-renderers.js#L141-L161
const (
	badgeHeight   = 28
	textMargin    = 12
	logoTextGuter = 6
	letterSpacing = 1.25
	logoWidth     = 14 // simple-icons named logo default
)

type charWidthRange struct {
	lower int
	upper int
	width float64
}

var (
	verdanaBold []charWidthRange
	// emWidth is the fallback advance width used for code points missing from
	// the table, matching anafanafo's `emWidth = widthOf('m')`.
	emWidth float64
)

func init() {
	var raw [][3]float64
	if err := json.Unmarshal(verdanaBoldData, &raw); err != nil {
		log.Panicln(err)
	}

	verdanaBold = make([]charWidthRange, len(raw))
	for i, r := range raw {
		verdanaBold[i] = charWidthRange{lower: int(r[0]), upper: int(r[1]), width: r[2]}
	}

	emWidth = widthOf("m")
}

// isControlChar mirrors anafanafo's CharWidthTableConsumer.isControlChar.
func isControlChar(charCode int) bool {
	return charCode <= 31 || charCode == 127
}

// widthOfCharCode returns the advance width of a single code point and whether
// it was found in the table. Control characters have zero width.
//
// ref. CharWidthTableConsumer.widthOfCharCode
func widthOfCharCode(charCode int) (float64, bool) {
	if isControlChar(charCode) {
		return 0, true
	}

	i := sort.Search(len(verdanaBold), func(i int) bool {
		return verdanaBold[i].lower >= charCode
	})

	// Exact match on the beginning of a range.
	if i < len(verdanaBold) && verdanaBold[i].lower == charCode {
		return verdanaBold[i].width, true
	}

	// Otherwise the code point may fall inside the preceding range.
	if i > 0 {
		r := verdanaBold[i-1]
		if charCode >= r.lower && charCode <= r.upper {
			return r.width, true
		}
	}

	return 0, false
}

// widthOf sums the advance width of every code point in text using the bold
// Verdana table. Code points missing from the table fall back to emWidth.
//
// ref. CharWidthTableConsumer.widthOf
func widthOf(text string) float64 {
	total := 0.0
	for _, r := range text {
		if w, ok := widthOfCharCode(int(r)); ok {
			total += w
		} else {
			total += emWidth
		}
	}
	return total
}

// renderedMessage reproduces how Shields.io renders the badge text for a given
// title. Because generateShieldSrc escapes every "-" as "--", dashes survive
// Shields' escapeFormat unchanged; only its underscore rule remains observable
// here: a lone "_" becomes a space and "__" collapses to "_".
//
// ref. https://github.com/badges/shields/blob/ee31d6164c01c32cb9016cba49d858f569b82b29/core/badge-urls/path-helpers.js
func renderedMessage(title string) string {
	var b strings.Builder
	runes := []rune(title)
	for i := 0; i < len(runes); {
		if runes[i] != '_' {
			b.WriteRune(runes[i])
			i++
			continue
		}
		// Consume a maximal run of underscores.
		j := i
		for j < len(runes) && runes[j] == '_' {
			j++
		}
		n := j - i
		b.WriteString(strings.Repeat("_", n/2))
		if n%2 == 1 {
			b.WriteByte(' ')
		}
		i = j
	}
	return b.String()
}

// utf16Len returns the number of UTF-16 code units in s, matching JavaScript's
// String.prototype.length used for letter-spacing in the renderer.
func utf16Len(s string) int {
	return len(utf16.Encode([]rune(s)))
}

// forTheBadgeWidth computes the rendered pixel width of a for-the-badge style
// badge for the given title, assuming an empty label, a single message (the
// title) and a 14px simple-icons logo — the shape every badge in this repo
// takes after dash-escaping. Height is always badgeHeight (28).
//
// ref. badge-renderers.js forTheBadge (no-label + logo branch)
func forTheBadgeWidth(title string) int {
	message := strings.ToUpper(renderedMessage(title))

	var messageTextWidth float64
	if message != "" {
		messageTextWidth = math.Floor(widthOf(message)) + letterSpacing*float64(utf16Len(message))
	}

	// needsLabelRect == false, logo present:
	//   messageRectWidth = 2*textMargin + logoWidth + gutter + messageTextWidth
	width := 2*textMargin + logoWidth + logoTextGuter + messageTextWidth
	return int(math.Round(width))
}
