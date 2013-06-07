// mdr.go

// Copyright ©2013 David Rook except where otherwise noted. All rights
// reserved. Use of this file is governed by a BSD-style license that can be
// found in the LICENSE_MDR.md file.

// Features:
//	map from color name string to RGBA value
//	map from RGBA value to color name - or nearest color name
//	colornames can be added by user to supplement or override basic names
package colorlist

import (
	// below are go 1.1 standard lib packages only
	"fmt"
	"image/color"
	"strings"
)

type ColorPair struct {
	RGBAval color.RGBA
	Name    string
}

var ColorNameMap map[string]color.RGBA
var ColorValMap map[color.RGBA]string

func init() {
	var ColorNamesAry = []ColorPair{
		// first 16 are the HTML4.01 color list aka sRGB (en.wikipedia.org/wiki/Web_colors
		{color.RGBA{0, 255, 255, 255}, "Aqua"},
		{color.RGBA{0, 0, 255, 255}, "Blue"},
		{color.RGBA{0, 0, 0, 255}, "Black"},
		{color.RGBA{255, 0, 255, 255}, "Fuchsia"},
		{color.RGBA{0, 255, 0, 255}, "Lime"},
		{color.RGBA{255, 255, 0, 255}, "Yellow"},
		{color.RGBA{0, 128, 128, 255}, "Teal"},
		{color.RGBA{192, 192, 192, 255}, "Silver"},
		{color.RGBA{128, 128, 128, 255}, "Gray"},
		{color.RGBA{0, 128, 0, 255}, "Green"},
		{color.RGBA{0, 0, 128, 255}, "Navy"},
		{color.RGBA{128, 0, 0, 255}, "Maroon"},
		{color.RGBA{128, 128, 0, 255}, "Olive"},
		{color.RGBA{128, 0, 128, 255}, "Purple"},
		{color.RGBA{255, 0, 0, 255}, "Red"},
		{color.RGBA{255, 255, 255, 255}, "White"},
		// so-called Web-Safe colors
		{color.RGBA{240, 248, 255, 255}, "Aliceblue"},
		{color.RGBA{250, 235, 215, 255}, "Antiquewhite"},
		{color.RGBA{127, 255, 212, 255}, "Aquamarine"},
		{color.RGBA{240, 255, 255, 255}, "Azure"},
		{color.RGBA{245, 245, 220, 255}, "Beige"},
		{color.RGBA{255, 228, 196, 255}, "Bisque"},
		{color.RGBA{255, 235, 205, 255}, "Blanchedalmond"},
		{color.RGBA{138, 43, 226, 255}, "Blueviolet"},
		{color.RGBA{165, 42, 42, 255}, "Brown"},
		{color.RGBA{222, 184, 135, 255}, "Burlywood"},
		{color.RGBA{95, 158, 160, 255}, "Cadetblue"},
		{color.RGBA{127, 255, 0, 255}, "Chartreuse"},
		{color.RGBA{210, 105, 30, 255}, "Chocolate"},
		{color.RGBA{255, 127, 80, 255}, "Coral"},
		{color.RGBA{100, 149, 237, 255}, "Cornflowerblue"},
		{color.RGBA{255, 248, 220, 255}, "Cornsilk"},
		{color.RGBA{220, 20, 60, 255}, "Crimson"},
		{color.RGBA{0, 255, 255, 255}, "Cyan"},
		{color.RGBA{0, 0, 139, 255}, "Darkblue"},
		{color.RGBA{0, 139, 139, 255}, "Darkcyan"},
		{color.RGBA{184, 134, 11, 255}, "Darkgoldenrod"},
		{color.RGBA{169, 169, 169, 255}, "Darkgray"},
		{color.RGBA{0, 100, 0, 255}, "Darkgreen"},
		{color.RGBA{169, 169, 169, 255}, "Darkgrey"},
		{color.RGBA{189, 183, 107, 255}, "Darkkhaki"},
		{color.RGBA{139, 0, 139, 255}, "Darkmagenta"},
		{color.RGBA{85, 107, 47, 255}, "Darkolivegreen"},
		{color.RGBA{255, 140, 0, 255}, "Darkorange"},
		{color.RGBA{153, 50, 204, 255}, "Darkorchid"},
		{color.RGBA{139, 0, 0, 255}, "Darkred"},
		{color.RGBA{233, 150, 122, 255}, "Darksalmon"},
		{color.RGBA{143, 188, 143, 255}, "Darkseagreen"},
		{color.RGBA{72, 61, 139, 255}, "Darkslateblue"},
		{color.RGBA{47, 79, 79, 255}, "Darkslategray"},
		{color.RGBA{47, 79, 79, 255}, "Darkslategrey"},
		{color.RGBA{0, 206, 209, 255}, "Darkturquoise"},
		{color.RGBA{148, 0, 211, 255}, "Darkviolet"},
		{color.RGBA{255, 20, 147, 255}, "Deeppink"},
		{color.RGBA{0, 191, 255, 255}, "Deepskyblue"},
		{color.RGBA{105, 105, 105, 255}, "Dimgray"},
		{color.RGBA{105, 105, 105, 255}, "Dimgrey"},
		{color.RGBA{30, 144, 255, 255}, "Dodgerblue"},
		{color.RGBA{178, 34, 34, 255}, "Firebrick"},
		{color.RGBA{255, 250, 240, 255}, "Floralwhite"},
		{color.RGBA{34, 139, 34, 255}, "Forestgreen"},
		{color.RGBA{220, 220, 220, 255}, "Gainsboro"},
		{color.RGBA{248, 248, 255, 255}, "Ghostwhite"},
		{color.RGBA{255, 215, 0, 255}, "Gold"},
		{color.RGBA{218, 165, 32, 255}, "Goldenrod"},
		{color.RGBA{128, 128, 128, 255}, "Grey"},
		{color.RGBA{173, 255, 47, 255}, "Greenyellow"},
		{color.RGBA{240, 255, 240, 255}, "Honeydew"},
		{color.RGBA{255, 105, 180, 255}, "Hotpink"},
		{color.RGBA{205, 92, 92, 255}, "Indianred"},
		{color.RGBA{75, 0, 130, 255}, "Indigo"},
		{color.RGBA{255, 255, 240, 255}, "Ivory"},
		{color.RGBA{240, 230, 140, 255}, "Khaki"},
		{color.RGBA{230, 230, 250, 255}, "Lavender"},
		{color.RGBA{255, 240, 245, 255}, "Lavenderblush"},
		{color.RGBA{124, 252, 0, 255}, "Lawngreen"},
		{color.RGBA{255, 250, 205, 255}, "Lemonchiffon"},
		{color.RGBA{173, 216, 230, 255}, "Lightblue"},
		{color.RGBA{240, 128, 128, 255}, "Lightcoral"},
		{color.RGBA{224, 255, 255, 255}, "Lightcyan"},
		{color.RGBA{250, 250, 210, 255}, "Lightgoldenrodyellow"},
		{color.RGBA{211, 211, 211, 255}, "Lightgray"},
		{color.RGBA{144, 238, 144, 255}, "Lightgreen"},
		{color.RGBA{211, 211, 211, 255}, "Lightgrey"},
		{color.RGBA{255, 182, 193, 255}, "Lightpink"},
		{color.RGBA{255, 160, 122, 255}, "Lightsalmon"},
		{color.RGBA{32, 178, 170, 255}, "Lightseagreen"},
		{color.RGBA{135, 206, 250, 255}, "Lightskyblue"},
		{color.RGBA{119, 136, 153, 255}, "Lightslategray"},
		{color.RGBA{119, 136, 153, 255}, "Lightslategrey"},
		{color.RGBA{176, 196, 222, 255}, "Lightsteelblue"},
		{color.RGBA{255, 255, 224, 255}, "Lightyellow"},
		{color.RGBA{50, 205, 50, 255}, "Limegreen"},
		{color.RGBA{250, 240, 230, 255}, "Linen"},
		{color.RGBA{255, 0, 255, 255}, "Magenta"},
		{color.RGBA{102, 205, 170, 255}, "Mediumaquamarine"},
		{color.RGBA{0, 0, 205, 255}, "Mediumblue"},
		{color.RGBA{186, 85, 211, 255}, "Mediumorchid"},
		{color.RGBA{147, 112, 219, 255}, "Mediumpurple"},
		{color.RGBA{60, 179, 113, 255}, "Mediumseagreen"},
		{color.RGBA{123, 104, 238, 255}, "Mediumslateblue"},
		{color.RGBA{0, 250, 154, 255}, "Mediumspringgreen"},
		{color.RGBA{72, 209, 204, 255}, "Mediumturquoise"},
		{color.RGBA{199, 21, 133, 255}, "Mediumvioletred"},
		{color.RGBA{25, 25, 112, 255}, "Midnightblue"},
		{color.RGBA{245, 255, 250, 255}, "Mintcream"},
		{color.RGBA{255, 228, 225, 255}, "Mistyrose"},
		{color.RGBA{255, 228, 181, 255}, "Moccasin"},
		{color.RGBA{255, 222, 173, 255}, "Navajowhite"},
		{color.RGBA{253, 245, 230, 255}, "Oldlace"},
		{color.RGBA{107, 142, 35, 255}, "Olivedrab"},
		{color.RGBA{255, 165, 0, 255}, "Orange"},
		{color.RGBA{255, 69, 0, 255}, "Orangered"},
		{color.RGBA{218, 112, 214, 255}, "Orchid"},
		{color.RGBA{238, 232, 170, 255}, "Palegoldenrod"},
		{color.RGBA{152, 251, 152, 255}, "Palegreen"},
		{color.RGBA{175, 238, 238, 255}, "Paleturquoise"},
		{color.RGBA{219, 112, 147, 255}, "Palevioletred"},
		{color.RGBA{255, 239, 213, 255}, "Papayawhip"},
		{color.RGBA{255, 218, 185, 255}, "Peachpuff"},
		{color.RGBA{205, 133, 63, 255}, "Peru"},
		{color.RGBA{255, 192, 203, 255}, "Pink"},
		{color.RGBA{221, 160, 221, 255}, "Plum"},
		{color.RGBA{176, 224, 230, 255}, "Powderblue"},
		{color.RGBA{188, 143, 143, 255}, "Rosybrown"},
		{color.RGBA{65, 105, 225, 255}, "Royalblue"},
		{color.RGBA{139, 69, 19, 255}, "Saddlebrown"},
		{color.RGBA{250, 128, 114, 255}, "Salmon"},
		{color.RGBA{244, 164, 96, 255}, "Sandybrown"},
		{color.RGBA{46, 139, 87, 255}, "Seagreen"},
		{color.RGBA{255, 245, 238, 255}, "Seashell"},
		{color.RGBA{160, 82, 45, 255}, "Sienna"},
		{color.RGBA{135, 206, 235, 255}, "Skyblue"},
		{color.RGBA{106, 90, 205, 255}, "Slateblue"},
		{color.RGBA{112, 128, 144, 255}, "Slategray"},
		{color.RGBA{112, 128, 144, 255}, "Slategrey"},
		{color.RGBA{255, 250, 250, 255}, "Snow"},
		{color.RGBA{0, 255, 127, 255}, "Springgreen"},
		{color.RGBA{70, 130, 180, 255}, "Steelblue"},
		{color.RGBA{210, 180, 140, 255}, "Tan"},
		{color.RGBA{216, 191, 216, 255}, "Thistle"},
		{color.RGBA{255, 99, 71, 255}, "Tomato"},
		{color.RGBA{64, 224, 208, 255}, "Turquoise"},
		{color.RGBA{238, 130, 238, 255}, "Violet"},
		{color.RGBA{245, 222, 179, 255}, "Wheat"},
		{color.RGBA{245, 245, 245, 255}, "Whitesmoke"},
		{color.RGBA{154, 205, 50, 255}, "Yellowgreen"},
		// end of Web-Safe colors
		// more colors if needed can be added below
		{color.RGBA{255, 255, 224, 255}, "Paleyellow"}, // alias of Lightyellow
	}

	ColorNameMap = make(map[string]color.RGBA)
	ColorValMap = make(map[color.RGBA]string)

	for _, c := range ColorNamesAry {
		name := strings.ToLower(c.Name)
		ColorNameMap[name] = c.RGBAval
		ColorValMap[c.RGBAval] = name
	}

}

func AddColor(name string, r, g, b, a int) {
	colorval := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	ColorNameMap[name] = colorval
	ColorValMap[colorval] = strings.ToLower(name)
	//	fmt.Printf("colorlist.AddColor() added %s\n",name)
}

// returns empty string if no match, see also NearestColorName()
//	 all names are lowercase on return regardless of case when added
func ColorName(c color.RGBA) string {
	if name, ok := ColorValMap[c]; ok {
		return name
	} else {
		return ""
	}
}

// returns hex triplet as #abcdef string
func SVGColorStr(s string) string {
	c := ColorVal(s)
	rv := fmt.Sprintf("#%02x%02x%02x",c.R,c.G,c.B)
	return rv
}

// returns black {0,0,0,255} if no match
func ColorVal(s string) color.RGBA {
	rv := color.RGBA{0, 0, 0, 255} // Black is default
	if len(s) <= 1 {
		return rv
	}
	if s[0] == '#' {
		// TODO(mdr): should we memoize it or not?
		return HexToColorRGBA(s)
	}
	if rgba, ok := ColorNameMap[strings.ToLower(s)]; ok {
		return rgba
	} else {
		return rv
	}
}

// the map lengths need not be same, for example; if Lightblue and Paleblue map to same value
//  just a convenience function, useful if you're loading colormaps from files.
func mapLen() (names int, values int) {
	return len(ColorNameMap), len(ColorValMap)
}

// naive compare
//  is there a better method?
func colorDiff(a, b color.RGBA) int64 {
	var sumsq int64
	// note - extra promotions to wider signed int are required? convert inner pairs to int64
	sumsq = int64(int32(a.R)-int32(b.R)) * int64(int32(a.R)-int32(b.R))  // diff in red component sq'd
	sumsq += int64(int32(a.G)-int32(b.G)) * int64(int32(a.G)-int32(b.G)) // diff in green component sq'd
	sumsq += int64(int32(a.B)-int32(b.B)) * int64(int32(a.B)-int32(b.B)) // diff in blue component sq'd
	return sumsq
}

// always returns a name even if wildly off
func ColorNameNearest(c color.RGBA) string {
	// assume black is closest to start with {0,0,0,-}  any starting color would work
	bestDiff := colorDiff(ColorVal("black"), c)
	bestName := "black"

	for rgba, cName := range ColorValMap {
		if c == rgba {
			return cName
		}
		// not equal so see how good the match is
		diff := colorDiff(c, rgba)
		if false {
			fmt.Printf("%s diff value is %d\n", cName, diff)
		}
		if diff < bestDiff { // and save if its better than current best
			bestDiff = diff
			bestName = cName
		}
	}
	return bestName
}

// returns value of hex char
func validHexChar(c byte) int {
	var rv = -1
	var hexchars []byte = []byte("0123456789abcdefABCDEF")
	for ndx, h := range hexchars {
		if c == h {
			rv = int(ndx)
			break
		}
	}
	if rv < 0 {
		return -1
	}
	if rv > 15 {
		rv -= 6
	}
	return rv
}

func validHexString(s string) bool {
	for _, c := range s {
		if validHexChar(byte(c)) < 0 {
			return false
		}
	}
	return true
}

// convert #ccc -> {12,12,12,255}
func hex3ToColorRGBA(s string) color.RGBA {
	r := validHexChar(s[0])
	g := validHexChar(s[1])
	b := validHexChar(s[2])
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// convert #0c0c0c -> {12,12,12,255}
func hex6ToColorRGBA(s string) color.RGBA {
	r := validHexChar(s[0])*16 + validHexChar(s[1])
	g := validHexChar(s[2])*16 + validHexChar(s[3])
	b := validHexChar(s[4])*16 + validHexChar(s[5])
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// convert #ccc or #0c0c0c -> {12,12,12,255}
//   must be either 3 or 6 hex chars, preceeded by number symbol
//   returns black if bad format encountered
func HexToColorRGBA(s string) color.RGBA {
	//fmt.Printf("Converting hexToColorRGBA(%s)\n",s)
	rv := color.RGBA{0, 0, 0, 255}
	if len(s) <= 1 {
		return rv
	}
	if s[0] != '#' {
		return rv
	}
	s = strings.ToLower(s)
	if !validHexString(s[1:]) {
		return rv
	}
	if len(s) == 4 {
		return hex3ToColorRGBA(s[1:])
	}
	if len(s) == 7 {
		return hex6ToColorRGBA(s[1:])
	}
	return rv
}

