// mdr.go

// Copyright ©2013 David Rook except where otherwise noted. All rights
// reserved. Use of this file is governed by a BSD-style license that can be
// found in the LICENSE_MDR.md file.

// Package colorlist implements a color name to/from RGBA value mapping.
//
// Features:
//	map from color name string to RGBA value
//	map from RGBA value to color name - or nearest color name
//	colornames can be added by user to supplement or override basic names
//	safe for concurrent use
package colorlist

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"strings"
	"sync"
)

type colorPair struct {
	RGBAval color.RGBA
	Name    string
}

var mapM sync.Mutex // Is this overkill?
var colorNameMap map[string]color.RGBA
var colorValMap map[color.RGBA]string

func init() {
	var colorNames = []colorPair{
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

	// no mutex required during init()
	colorNameMap = make(map[string]color.RGBA, len(colorNames))
	colorValMap = make(map[color.RGBA]string, len(colorNames))

	for _, c := range colorNames {
		name := strings.ToLower(c.Name)
		colorNameMap[name] = c.RGBAval
		colorValMap[c.RGBAval] = name
	}

}

// NotFoundError values describe a failure to lookup a color name.
type NotFoundError string

func (e NotFoundError) Error() string {
	return fmt.Sprintf("colorlist: color %q not found", string(e))
}

// AddColor adds a new color to the map.
// Just a convience wrapper to AddColorRGBA().
func AddColor(name string, r, g, b, a int) {
	AddColorRGBA(name, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
}

// AddColorRGBA adds a new color to the map.
//
// BUG(x): if the name matches an existing color then
// for the old color value x: ColorVal(ColorName(x)) != x
func AddColorRGBA(name string, colorval color.RGBA) {
	name = strings.ToLower(name)
	mapM.Lock()
	colorNameMap[name] = colorval
	colorValMap[colorval] = name
	mapM.Unlock()
	//	fmt.Printf("colorlist.AddColor() added %s\n",name)
}

// ColorName returns the name of the color or the empty string if there
// is no match.
// All names are returned lowercase regardless of case when added.
//
// See Also: ColorNameNearest()
func ColorName(c color.RGBA) string {
	mapM.Lock()
	defer mapM.Unlock()
	return colorValMap[c]
}

// SVGColorStr returns the hex triplet color value for the given color name.
//
// E.g. "#12a0f0".
func SVGColorStr(s string) string {
	c := ColorVal(s)
	var b [7]byte
	b[0] = '#'
	hex.Encode(b[1:], []byte{byte(c.R), byte(c.G), byte(c.B)})
	return string(b[:])
}

// ColorVal returns the RGBA color of the named color
// or black {0, 0, 0, 255} if no match is found.
//
// Also accepts three or six character hex values prefixed with '#'
// (e.g. #ccc or #0c0c0c).
func ColorVal(s string) color.RGBA {
	rgba, err := Color(s)
	if err != nil {
		return color.RGBA{0, 0, 0, 255}
	}
	return rgba
}

// Color returns the RGBA color of the named color, if found.
// If not found either an error from HexToColor() or NotFoundError is returned.
//
// Also accepts three or six character hex values prefixed with '#'
// (e.g. #ccc or #0c0c0c).
func Color(s string) (rgba color.RGBA, err error) {
	if len(s) >= 1 && s[0] == '#' {
		return HexToColor(s)
	}
	mapM.Lock()
	defer mapM.Unlock()
	if rgba, ok := colorNameMap[strings.ToLower(s)]; ok {
		return rgba, nil
	}
	return rgba, NotFoundError(s)
}

// the map lengths need not be same, for example; if Lightblue and Paleblue map to same value
//  just a convenience function, useful if you're loading colormaps from files.
// XXX loading/saving should probably be handled by this package if required.
func mapLen() (names int, values int) {
	mapM.Lock()
	defer mapM.Unlock()
	return len(colorNameMap), len(colorValMap)
}

// colorDiff returns the sum of the RGB differences squared.
// TODO(x): naive color comparison, is there a better method?
func colorDiff(a, b color.RGBA) (sumsq uint64) {
	// note - extra promotions to wider signed int are required? convert inner pairs to uint64
	sumsq = uint64(int32(a.R)-int32(b.R)) * uint64(int32(a.R)-int32(b.R))
	sumsq += uint64(int32(a.G)-int32(b.G)) * uint64(int32(a.G)-int32(b.G))
	sumsq += uint64(int32(a.B)-int32(b.B)) * uint64(int32(a.B)-int32(b.B))
	return sumsq
}

// ColorNameNearest is like ColorName but returns the closest match,
// even if wildly off.
func ColorNameNearest(c color.RGBA) string {
	// assume black is closest to start with {0,0,0,-}  any starting color would work
	bestDiff := colorDiff(color.RGBA{0, 0, 0, 255}, c)
	bestName := "black"

	mapM.Lock()
	defer mapM.Unlock()
	if cName, ok := colorValMap[c]; ok {
		return cName
	}
	for rgba, cName := range colorValMap {
		// not equal so see how good the match is
		diff := colorDiff(c, rgba)
		//fmt.Printf("%s diff value is %d\n", cName, diff)
		// and save if its better than current best
		if diff < bestDiff || (diff == bestDiff && cName < bestName) {
			bestDiff = diff
			bestName = cName
			if diff <= 0 {
				return bestName
			}
		}
	}
	return bestName
}

// fromHexChar converts a hex character into its value and a success flag.
// From encoding/hex/hex.go
func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}

// HexToColorRGBA converts hex colors to RGBA colors.
// Input must be either 3 or 6 hex chars, preceeded by '#'.
// Returns black if bad format encountered.
//
// E.g. convert #ccc or #0c0c0c -> {12,12,12,255}
func HexToColorRGBA(s string) color.RGBA {
	if c, err := HexToColor(s); err != nil {
		return color.RGBA{0, 0, 0, 255}
	} else {
		return c
	}
}

// HexToColor converts hex colors to RGBA colors.
// Input must be either 3 or 6 hex chars, preceeded by '#'.
//
// On failure, returns hex.ErrLength or hex.InvalidByteError (from encoding/hex).
//
// E.g. convert #ccc or #0c0c0c -> {12,12,12,255}
func HexToColor(s string) (c color.RGBA, err error) {
	//fmt.Printf("Converting hexToColorRGBA(%s)\n",s)
	if len(s) != 4 && len(s) != 7 {
		return c, hex.ErrLength
	}
	if s[0] != '#' {
		return c, hex.InvalidByteError(s[0])
	}
	s = s[1:]
	var result [3]uint8
	switch len(s) {
	case 3:
		for i := 0; i < 3; i++ {
			a, ok := fromHexChar(s[i])
			if !ok {
				return c, hex.InvalidByteError(s[i])
			}
			result[i] = a
		}
	case 6:
		for i := 0; i < 3; i++ {
			a, ok := fromHexChar(s[i*2])
			if !ok {
				return c, hex.InvalidByteError(s[i*2])
			}
			b, ok := fromHexChar(s[i*2+1])
			if !ok {
				return c, hex.InvalidByteError(s[i*2+1])
			}
			result[i] = (a << 4) | b
		}
	}
	return color.RGBA{result[0], result[1], result[2], 255}, nil
}
