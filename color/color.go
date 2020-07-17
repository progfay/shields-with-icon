package color

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	redMask   = 0xFF0000
	greenMask = 0x00FF00
	blueMask  = 0x0000FF
)

var (
	colorCodeRegExp = regexp.MustCompile(`^#?[0-9a-fA-F]{6}$`)
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
	Code  string
}

func NewColor(colorCode string) (*Color, error) {
	if !colorCodeRegExp.Match([]byte(colorCode)) {
		return nil, fmt.Errorf("invalid color code")
	}

	code := strings.TrimLeft(colorCode, "#")
	rgb, err := strconv.ParseUint(code, 16, 24)
	if err != nil {
		return nil, err
	}

	c := &Color{
		Red:   float64(rgb&redMask>>16) / 255,
		Green: float64(rgb&greenMask>>8) / 255,
		Blue:  float64(rgb&blueMask>>0) / 255,
		Code:  code,
	}

	return c, nil
}
