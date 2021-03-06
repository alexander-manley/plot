// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image"
	"math"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

// Image is a plotter that draws a scaled, raster image.
type Image struct {
	img            image.Image
	cols           int
	rows           int
	xmin, xmax, dx float64
	ymin, ymax, dy float64
}

// NewImage creates a new image plotter.
// Image will plot img inside the rectangle defined by the
// (xmin, ymin) and (xmax, ymax) points given in the data space.
// The img will be scaled to fit inside the rectangle.
func NewImage(img image.Image, xmin, ymin, xmax, ymax float64) *Image {
	bounds := img.Bounds()
	cols := bounds.Dx()
	rows := bounds.Dy()
	dx := math.Abs(xmax-xmin) / float64(cols)
	dy := math.Abs(ymax-ymin) / float64(rows)
	return &Image{
		img:  img,
		cols: cols,
		rows: rows,
		xmin: xmin,
		xmax: xmax,
		dx:   dx,
		ymin: ymin,
		ymax: ymax,
		dy:   dy,
	}
}

// Plot implements the Plot method of the plot.Plotter interface.
func (img *Image) Plot(c draw.Canvas, p *plot.Plot) {
	trX, trY := p.Transforms(&c)
	xmin := trX(img.xmin)
	ymin := trY(img.ymin)
	xmax := trX(img.xmax)
	ymax := trY(img.ymax)
	rect := vg.Rectangle{
		Min: vg.Point{xmin, ymin},
		Max: vg.Point{xmax, ymax},
	}
	c.DrawImage(rect, img.img)
}

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (img *Image) DataRange() (xmin, xmax, ymin, ymax float64) {
	return img.xmin, img.xmax, img.ymin, img.ymax
}

// GlyphBoxes implements the GlyphBoxes method
// of the plot.GlyphBoxer interface.
func (img *Image) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	return nil
}

func (img *Image) x(c int) float64 {
	if c >= img.cols || c < 0 {
		panic("plotter/image: illegal range")
	}
	return img.xmin + float64(c)*img.dx
}

func (img *Image) y(r int) float64 {
	if r >= img.rows || r < 0 {
		panic("plotter/image: illegal range")
	}
	return img.ymin + float64(r)*img.dy
}
