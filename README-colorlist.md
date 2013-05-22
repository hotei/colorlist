<center>
# colorlist
</center>

[This go package] [4] started as a fork of github.com/btracey/colorlist.  Since then
it has diverged enough that retaining it as a fork is probably not useful.  Brendan
decided to change his repository name to colors so I've un-hitched this from
his original repository which is now marked 'deprecated'.  This is now being
re-issued as a new repository : github.com/hotei/colorlist.

The main difference between the packages is that Brendans provides pre-named
variables representing colors.  For example colors.Black is an RGBA value in
his package. I chose to go the route of color lookup tables.  The same value in
my package would be colorlist.ColorVal("black").  You can also specify color values
with rgb hex codes.  Evaluating colorlist.ColorVal("#f5F5f5") returns the same
result as colorlist.ColorVal("WhiteSmoke")  This can be useful if you're using
a color picker from PhotoShop or Gimp to grab specific colors.

There are advantages and disadvantages to both approaches.  In his system a bad 
color name results in a compile time error - which is not always a bad thing.  
In mine it results in a default color (black) being substituted but the program 
continues to run - though of course it might not produce the expected colors (unless
you spelled it "blackk" :-) .

With my system configuration files can easily map colors from strings to
RGBA values.  You can also add your own colors to the maps on the fly and even
change the existing color values if you wish.

### Features

This package contains a list of about 150 standard colors in image.color RGBA format.

The first set (16) colors is specified by the [HTML 4 standard][1]

The second set of  colors are specified by the [CSS3 standard][3]

[Color-book.org][2] is a web site with many colors displayed and listed in
both rgb and cmyk forms.

There's a test file with examples to show usage.

### Take a Quick Look at the Exported Interface

Inspect the [GoDocs][5] at godoc.org

### Installation and import

```
go get github.com/hotei/colorlist

import "github.com/hotei/colorlist"
```


### Resources

* [Color-Book.org][2] a web site with many colors displayed and listed in
both rgb and cmyk forms.
* [W3 HTML 4 spec][1]
* [W3 svg color spec][3]
* [X-11 Color names][7]

[1]: http://www.w3.org/TR/REC-html40/types.html#h-6.5	"HTML 4 color info"
[2]: http://color-book.org/color-index,a "color-book.org"
[3]: http://www.w3.org/TR/css3-color/#svg-color "www.W3.org svg color"
[4]: http://www.github.com/hotei/colorlist "github/hotei/colorlist"
[5]: http://godoc.org/github.com/hotei/colorlist "GoDoc.org"
[7]: http://en.wikipedia.org/wiki/X11_color_names "X-11 color names"

Copyright ©2013 David Rook except where otherwise noted. All rights
reserved. Use of this file is governed by a BSD-style license that can be
found in the LICENSE_MDR.md file.

