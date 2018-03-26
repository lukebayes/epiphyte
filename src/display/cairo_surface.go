package display

import (
	"fmt"
	"github.com/golang-ui/cairo"
	"math"
)

type cairoSurfaceAdapter struct {
	context *cairo.Cairo
}

func (c *cairoSurfaceAdapter) MoveTo(x float64, y float64) {
	cairo.MoveTo(c.context, x, y)
}

func (c *cairoSurfaceAdapter) SetRgba(r, g, b, a float64) {
	cairo.SetSourceRgba(c.context, r, g, b, a)
}

func (c *cairoSurfaceAdapter) SetLineWidth(width float64) {
	cairo.SetLineWidth(c.context, width)
}

func (c *cairoSurfaceAdapter) Stroke() {
	cairo.Stroke(c.context)
}

func (c *cairoSurfaceAdapter) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	cairo.Arc(c.context, xc, yc, radius, angle1, angle2)
}

func (c *cairoSurfaceAdapter) DrawRectangle(x float64, y float64, width float64, height float64) {
	fmt.Println("DRAW RECT:", x, y, width, height)
	cairo.MakeRectangle(c.context, x, y, width, height)
}

func (c *cairoSurfaceAdapter) Fill() {
	cairo.Fill(c.context)
}

func (c *cairoSurfaceAdapter) FillPreserve() {
	cairo.FillPreserve(c.context)
}

func (c *cairoSurfaceAdapter) getYOffsetFor(d Displayable) float64 {
	current := d
	offset := 0.0
	for current != nil {
		offset += math.Max(0, current.GetY())
		current = d.GetParent()
	}
	return offset
}

func (s *cairoSurfaceAdapter) getXOffsetFor(d Displayable) float64 {
	current := d
	offset := 0.0
	for current != nil {
		offset += math.Max(0, current.GetX())
		current = d.GetParent()
	}
	return offset
}

func (c *cairoSurfaceAdapter) GetOffsetSurfaceFor(d Displayable) Surface {
	fmt.Println("Cairo.GetOffSurfaceFor:", d.GetId())
	x := c.getXOffsetFor(d)
	y := c.getYOffsetFor(d)
	return &SurfaceDelegate{
		delegateTo: c,
		offsetX:    x,
		offsetY:    y,
	}
}

func NewCairoSurfaceAdapter(cairo *cairo.Cairo) Surface {
	return &cairoSurfaceAdapter{context: cairo}
}
