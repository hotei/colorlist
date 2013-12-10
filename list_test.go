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

func Test_0002(t *testing.T) {
	grey := color.RGBA{12, 12, 12, 255}
	black := color.RGBA{0, 0, 0, 255}
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
	} {
		{color.RGBA{255,102,204,255}, "radish"},
	} {
		if g, e := ColorVal(d.name), black; g != e {
			t.Errorf("ColorVal(%q) returned %v, expected %v", d.name, g, e)
		} else {
			t.Logf("ColorVal(%q) returned %v, as expected", d.name, g)
		}
		//AddColor("radish", 255, 102, 204, 255)
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
		{color.RGBA{255,102,204,255}, "radish", "radish", "#ff66cc"},
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
