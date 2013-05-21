<center>
# colorlist
</center>

[This package] [4] started as a fork of github.com/btracey/colorlist.  Since then
it has diverged enough that retaining it as a fork is probably not useful.  Brendan
decided to change his repository name to colors so I've un-hitched this from
his original repositor which is now marked 'deprecated'.  This is now being
re-issued as a new repository : github.com/hotei/colorlist.

The main difference between the packages is that Brendans provides pre-named
variables representing colors.  I chose to go the route of color lookup tables.
There are advantages to both approaches.  

### Features

This package contains a list of colors in image.color GGBA format.
The first set of colors is specified by the [HTML 4 standard][1]
The second set of  colors are specified by the [CSS3 standard][3]

[Color-book.org][2] is a web site with many colors displayed and listed in
both rgb and cmyk forms.

### Have a Quick Look at Interface

Inspect the [GoDocs][5]

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
[5]: http://godoc.org/hotei/colorlist "GoDoc.org"
[7]: http://en.wikipedia.org/wiki/X11_color_names "X-11 color names"

Copyright ©2013 David Rook except where otherwise noted. All rights
reserved. Use of this file is governed by a BSD-style license that can be
found in the LICENSE_MDR.md file.
