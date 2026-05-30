package main

import "testing"

// Expected widths are round() of the exact width attribute returned by the live
// Shields.io for-the-badge SVG for each title's generated URL (dashes escaped
// as "--", so every badge is a single message with a 14px logo).
func TestForTheBadgeWidth(t *testing.T) {
	cases := []struct {
		title string
		want  int // round(live SVG width)
	}{
		{".ENV", 75},                // live 75
		{"A-Frame", 103},            // live 102.75 (single "A-FRAME")
		{"MobX-State-Tree", 169},    // live 168.75 (formerly "404: badge not found")
		{"Pop!_OS", 99},             // live 98.75 ("POP! OS": lone "_" -> space)
		{"30 seconds of code", 190}, // live 189.5
	}

	for _, c := range cases {
		if got := forTheBadgeWidth(c.title); got != c.want {
			t.Errorf("forTheBadgeWidth(%q) = %d, want %d", c.title, got, c.want)
		}
	}
}

func TestRenderedMessage(t *testing.T) {
	cases := []struct {
		title string
		want  string
	}{
		{"A-Frame", "A-Frame"}, // dashes are escaped upstream, so kept here
		{"Pop!_OS", "Pop! OS"}, // lone underscore -> space
		{"a__b", "a_b"},        // double underscore -> single
		{"a___b", "a_ b"},      // triple: pair + lone -> "_" + space
		{".ENV", ".ENV"},
	}

	for _, c := range cases {
		if got := renderedMessage(c.title); got != c.want {
			t.Errorf("renderedMessage(%q) = %q, want %q", c.title, got, c.want)
		}
	}
}
