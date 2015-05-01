// colorlist_test.go

// Copyright ©2013 David Rook except where otherwise noted. All rights
// reserved. Use of this file is governed by a BSD-style license that can be
// found in the LICENSE_MDR.md file.

// simplest way to use test is "go test"

package colorlist

import (
	"fmt"
	"image/color"
	"testing"
)

func Test_0002(t *testing.T) {
	fmt.Printf("\nTest_0002\n")

	s1 := "#ccc"
	s1b := "#cCc"

	s2 := "#0c0c0c"
	s2b := "#0c0C0c"

	s3 := "#0c0Z0c"

	// convert to rgba {12,12,12,255}
	// TODO(mdr): make test use t.Errorf if results don't work out right
	fmt.Printf("Everything should convert to { 12, 12, 12, 255 }\n")
	fmt.Printf(" %s to color.RGBA = %v\n", s1, HexToColorRGBA(s1))
	fmt.Printf(" %s to color.RGBA = %v\n", s1b, HexToColorRGBA(s1b))
	fmt.Printf(" %s to color.RGBA = %v\n", s2, HexToColorRGBA(s2))
	fmt.Printf(" %s to color.RGBA = %v\n", s2b, HexToColorRGBA(s2b))
	fmt.Printf("Next one should fail and return default of black {0,0,0,255}\n")
	fmt.Printf(" %s to color.RGBA = %v\n", s3, HexToColorRGBA(s3))
}

func Test_0001(t *testing.T) {
	fmt.Printf("\nTest_0001\n")

	if ColorName(color.RGBA{245, 245, 245, 255}) != "whitesmoke" {
		t.Errorf("ColorName failed for Whitesmoke\n")
	}

	fmt.Printf("ColorName({245,245,245,255}) = %s (should be whitesmoke)\n",
		ColorName(color.RGBA{245, 245, 245, 255}))

	fmt.Printf("NearestColorName({244,244,240,255}) = %s (should be whitesmoke)\n",
		ColorNameNearest(color.RGBA{244, 244, 240, 255}))

	if ColorNameNearest(color.RGBA{244, 244, 240, 255}) != "whitesmoke" {
		t.Errorf("NearestColorName failed\n")
	}

	ws := color.RGBA{245, 245, 245, 255} // whitesmoke
	if ColorVal("WhiteSmoke") != ws {    // note CamelCase isn't how it was added
		t.Errorf("ColorVal(string) failed got %v, not {244,244,240,255} \n", ColorVal("WhiteSmoke"))
	}

	AddColor("radish", 255, 102, 204, 255)

	//ws = ColorVal("#ccc")
	newSmoke := ColorVal("#f5f5F5")
	fmt.Printf("Testing newSmoke = #f5f5F5\n")
	if newSmoke != ws {
		t.Errorf("ColorVal(#string) failed got %v, not {244,244,240,255} \n", newSmoke)
	}

	svgSmoke := SVGColorStr("whitesmoke")
	fmt.Printf("svg name of whitesmoke should be #f5f5f5, got %s\n", svgSmoke)
	if svgSmoke != "#f5f5f5" {
		t.Errorf("svg name of whitesmoke should be #f5f5f5, not %s\n", svgSmoke)
	}
	svgSmoke = SVGColorStr("black")
	fmt.Printf("svg name of black should be #000000, got %s\n", svgSmoke)
	if svgSmoke != "#000000" {
		t.Errorf("svg name of black should be #000000, not %s\n", svgSmoke)
	}
	svgSmoke = SVGColorStr("radish")
	fmt.Printf("svg name of radish should be #ff66cc, got %s\n", svgSmoke)
	if svgSmoke != "#ff66cc" {
		t.Errorf("svg name of radish should be #ff66cc, not %s\n", svgSmoke)
	}
}
