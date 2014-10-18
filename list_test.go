// list_test.go

// Copyright ©2013 David Rook except where otherwise noted. All rights
// reserved. Use of this file is governed by a BSD-style license that can be
// found in the LICENSE_MDR.md file.

// simplest way to use test is "go test"

package colorlist

import (
	"image/color"
	"strings"
	"testing"
)

var black = color.RGBA{0, 0, 0, 255}

func Test_0002(t *testing.T) {
	grey := color.RGBA{12, 12, 12, 255}
	for _, d := range []struct {
		string
		color.RGBA
	}{
		{"#ccc", grey},
		{"#ccc", grey},
		{"#cCc", grey},
		{"#0c0c0c", grey},
		{"#0c0C0c", grey},
		{"#0c0Z0c", black},
	} {
		_ = d
		c := HexToColorRGBA(d.string)
		if c != d.RGBA {
			t.Errorf("HexToColorRGBA(%q) returned %v, expected %v",
				d.string, c, d.RGBA)
		} else {
			//t.Logf("HexToColorRGBA(%q) returned %v, as expected", d.string, c)
		}
	}
}

func Test_0001(t *testing.T) {
	whitesmoke := color.RGBA{245, 245, 245, 255} // whitesmoke
	for _, d := range []struct {
		color.RGBA
		name, nearest string
	}{
		{whitesmoke, "whitesmoke", "whitesmoke"},
		{color.RGBA{244, 244, 240, 255}, "", "whitesmoke"},
		{color.RGBA{0, 0, 0, 255}, "black", "black"},
		{color.RGBA{0, 0, 255, 255}, "blue", "blue"},
		{color.RGBA{0, 0, 231, 250}, "", "blue"},
		{color.RGBA{0, 0, 231, 5}, "", "blue"},
		{color.RGBA{0, 0, 230, 250}, "", "blue"},
		{color.RGBA{0, 0, 229, 250}, "", "mediumblue"},
		{color.RGBA{0, 0, 229, 5}, "", "mediumblue"},
	} {
		if g, e := ColorName(d.RGBA), d.name; g != e {
			t.Errorf("ColorName(%v) returned %q, expected %q", d.RGBA, g, e)
		} else {
			//t.Logf("ColorName(%v) returned %q, as expected", d.RGBA, g)
		}
		if g, e := ColorNameNearest(d.RGBA), d.nearest; g != e {
			t.Errorf("ColorNameNearest(%v) returned %q, expected %q", d.RGBA, g, e)
		} else {
			//t.Logf("ColorNameNearest(%v) returned %q, as expected", d.RGBA, g)
		}
	}

	black := color.RGBA{0, 0, 0, 255}
	for _, d := range []struct {
		color.RGBA
		name string
		dup  bool
	}{
		{color.RGBA{255, 102, 204, 255}, "radish", false},
		{color.RGBA{205, 175, 149, 255}, "PeachPuff3", false},
	} {
		if g, e := ColorVal(d.name), black; g != e {
			t.Errorf("ColorVal(%q) returned %v, expected %v", d.name, g, e)
		} else {
			//t.Logf("ColorVal(%q) returned %v, as expected", d.name, g)
		}
		r, g, b, a := d.RGBA.RGBA()
		AddColor(d.name, int(r), int(g), int(b), int(a))
		if g, e := ColorVal(d.name), d.RGBA; g != e {
			t.Errorf("ColorVal(%q) returned %v, expected %v", d.name, g, e)
		} else {
			//t.Logf("ColorVal(%q) returned %v, as expected", d.name, g)
		}
	}

	for _, d := range []struct {
		color.RGBA
		name, nearest, svg string
	}{
		{whitesmoke, "whitesmoke", "whitesmoke", "#f5f5f5"},
		{color.RGBA{0, 0, 0, 255}, "black", "black", "#000000"},
		{color.RGBA{0, 0, 255, 255}, "blue", "blue", "#0000ff"},
		{color.RGBA{255, 102, 204, 255}, "radish", "radish", "#ff66cc"},
	} {
		nm := strings.Title(d.name)
		if g, e := ColorVal(nm), d.RGBA; g != e {
			t.Errorf("ColorVal(%q) returned %v, expected %v", nm, g, e)
		} else {
			//t.Logf("ColorVal(%q) returned %v, as expected", nm, g)
		}
		if g, e := SVGColorStr(d.name), d.svg; g != e {
			t.Errorf("SVGColorStr(%q) returned %v, expected %v", d.name, g, e)
		} else {
			//t.Logf("SVGColorStr(%q) returned %v, as expected", d.name, g)
		}
		if g, e := ColorVal(d.svg), d.RGBA; g != e {
			t.Errorf("ColorVal(%q) returned %v, expected %v", d.svg, g, e)
		} else {
			//t.Logf("ColorVal(%q) returned %v, as expected", d.svg, g)
		}
	}
}

func TestBidirection(t *testing.T) {
	aliases := map[string]string{
		"cyan":           "aqua",
		"darkgrey":       "darkgray",
		"darkslategrey":  "darkslategray",
		"dimgrey":        "dimgray",
		"grey":           "gray",
		"lightgrey":      "lightgray",
		"lightslategrey": "lightslategray",
		"magenta":        "fuchsia",
		"paleyellow":     "lightyellow",
		"slategrey":      "slategray",
	}
	for name := range colorNameMap {
		if g := ColorName(ColorVal(name)); g != name {
			if aliases[g] == name {
				//t.Logf("Got alias %q for %q", g, name)
				continue
			}
			//fmt.Printf("\t\t%q: %q,\n", g, name)
			t.Errorf("ColorName(ColorVal(%q)) returned %q, expected %q",
				name, g, name)
		}
	}
	for c := range colorValMap {
		if g := ColorVal(ColorName(c)); g != c {
			t.Errorf("ColorVal(ColorName(%v)) returned %v, expected %v",
				c, g, c)
		}
	}

	for _, d := range []struct {
		color.RGBA
		name string
	}{
		{color.RGBA{225, 21, 61, 254}, "Crimson"},
		{color.RGBA{220, 20, 60, 255}, "Crimson2"},
	} {
		origRGBA := ColorVal(d.name)
		//origName := ColorName(d.RGBA)
		r, g, b, a := d.RGBA.RGBA()
		AddColor(d.name, int(r), int(g), int(b), int(a))
		for _, c := range []color.RGBA{d.RGBA, origRGBA} {
			if g := ColorVal(ColorName(c)); g != c {
				t.Errorf("After AddColor(%q): ColorVal(ColorName(%v)) returned %v, expected %v",
					d.name, c, g, c)
			}
		}
	}
}

func BenchmarkColorVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorVal("LightGoldenRodYellow")
	}
}

func BenchmarkSVGColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SVGColorStr("LightGoldenRodYellow")
	}
}

func BenchmarkNearMatch(b *testing.B) {
	c := ColorVal("LightGoldenRodYellow")
	for i := 0; i < b.N; i++ {
		ColorNameNearest(c)
	}
}

func BenchmarkNearNoMatch(b *testing.B) {
	c := ColorVal("LightGoldenRodYellow")
	c.R++
	c.A--
	for i := 0; i < b.N; i++ {
		ColorNameNearest(c)
	}
}

func BenchmarkHexToColorRGB_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexToColorRGBA("#abc")
	}
}

func BenchmarkHexToColorRGB_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexToColorRGBA("#5a5b5c")
	}
}

func BenchmarkHexToColorRGB_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexToColorRGBA("#ffffffffffffffffffff")
	}
}
