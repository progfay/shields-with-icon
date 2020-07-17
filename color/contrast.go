package color

import "math"

var White = Color{
	Red:   1,
	Green: 1,
	Blue:  1,
	Code:  "FFFFFF",
}

var Black = Color{
	Red:   0,
	Green: 0,
	Blue:  0,
	Code:  "000000",
}

func GetContrastRatio(foreground, background Color) float64 {
	// b := getRelativeLuminance(background)
	// f := getRelativeLuminance(foreground)
	return getContrastRatioOpaque(foreground, background)
}

func getContrastRatioOpaque(foreground, background Color) float64 {
	b := getRelativeLuminance(background)
	f := getRelativeLuminance(foreground)

	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
	if b < f {
		return (f + 0.05) / (b + 0.05)
	} else {
		return (b + 0.05) / (f + 0.05)
	}
}

func getRelativeLuminance(c Color) float64 {
	// https://www.w3.org/TR/2008/REC-WCAG20-20081211/#relativeluminancedef
	var r, g, b float64

	if c.Red <= 0.03928 {
		r += c.Red / 12.92
	} else {
		r += math.Pow((c.Red+0.055)/1.055, 2.4)
	}

	if c.Green <= 0.03928 {
		g += c.Green / 12.92
	} else {
		g += math.Pow((c.Green+0.055)/1.055, 2.4)
	}

	if c.Blue <= 0.03928 {
		b += c.Blue / 12.92
	} else {
		b += math.Pow((c.Blue+0.055)/1.055, 2.4)
	}

	return 0.2126*r + 0.7152*g + 0.0722*b
}
