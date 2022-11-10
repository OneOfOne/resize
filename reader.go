/*
Copyright (c) 2022, Ahmed W. <oneofone@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package resize

import (
	"image"
	"io"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// ResizeReader decodes an image from a reader and calls Resize with the given params.
// Supported formats are: PNG, JPEG, GIF by default.
// To support other formats, import the package that implements the format.
func ResizeReader(width, height uint, rd io.Reader, blur float64, interp InterpolationFunction) (img image.Image, format string, err error) {
	if img, format, err = image.Decode(rd); err == nil {
		img = ResizeWithBlur(width, height, img, blur, interp)
	}
	return
}
