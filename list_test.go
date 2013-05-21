// list_test.go

// Copyright ©2013 David Rook except where otherwise noted. All rights
// reserved. Use of this file is governed by a BSD-style license that can be
// found in the LICENSE_MDR.md file.


// simplest way to use test is "go test"

package 

colorlist

import (
	"fmt"
	"image/color"
	"testing"
)

func Test_0001(t *testing.T) {
	fmt.Printf("Test_0001\n")

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
	
	ws := color.RGBA{245,245,245,255}	// whitesmoke 
	if ColorVal("WhiteSmoke") != ws {	// note CamelCase isn't how it was added
		t.Errorf("ColorVal(string) failed got %v, not {244,244,240,255} \n",ColorVal("WhiteSmoke"))
	}
}
