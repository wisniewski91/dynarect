package dynarect

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// DynaRect - Basic struct that holds all information that is needed to generate rl.Rectangle for raylib
// Padding - distance inside container
type DynaRect struct {
	X          float32
	Y          float32
	Width      float32
	Height     float32
	Padding    float32 // Inside distance of parent rect
	Spacing    float32 // Space between child elements
	ChildCount float32 // Used to calculate width of child rects generated from parent rect
}

// DynaRectIterator - struct that allow you to generate child rects inside parent rect that align childs next to
// eachother counting available width, and moving to the next row if not enought space is left
type DynaRectIterator struct {
	ParentRect  DynaRect
	ChildSize   rl.Vector2
	ActualChild int
}

// GetRect - Method to retrieve raylib.Rectangle from DynaRect
func (r DynaRect) GetRect() rl.Rectangle {
	return rl.NewRectangle(r.X, r.Y, r.Width, r.Height)
}

// NewDynaRect - constructor for DynaRect
func NewDynaRect(x, y, width, height, padding, spacing float32, childCount float32) DynaRect {
	return DynaRect{
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		Padding:    padding,
		Spacing:    spacing,
		ChildCount: childCount,
	}
}

// WithChildCount - create new DynaRect with new ChildCount value.
//
// ChildCount is used to calculate width of child DynaRect generated from parent DynaRect
func (r DynaRect) WithChildCount(newCount float32) DynaRect {
	return DynaRect{
		X:          r.X,
		Y:          r.Y,
		Width:      r.Width,
		Height:     r.Height,
		Padding:    r.Padding,
		Spacing:    r.Spacing,
		ChildCount: newCount,
	}
}

// NextRow - Calculate new DynaRect using actual Y position and Height of this element.
//
// Y: Y + Spacing + Height
func (r DynaRect) NexRow() DynaRect {
	return DynaRect{
		X:          r.X,
		Y:          r.Y + r.Spacing + r.Height,
		Width:      r.Width,
		Height:     r.Height,
		Padding:    r.Padding,
		Spacing:    r.Spacing,
		ChildCount: r.ChildCount,
	}
}

// WithHeight - returns new DynaRect with new Height value
func (r DynaRect) WithHeight(height float32) DynaRect {
	return DynaRect{
		X:          r.X,
		Y:          r.Y,
		Width:      r.Width,
		Height:     height,
		Padding:    r.Padding,
		Spacing:    r.Spacing,
		ChildCount: r.ChildCount,
	}
}

// GetChild - returns new DynaRect calculated based on child identifier
//
// It calculates Width and X of childRect. Other values are the same
func (r DynaRect) GetChild(child float32) DynaRect {
	spacingTotalSize := 2*r.Padding + ((r.ChildCount - 1) * r.Spacing)
	childWidth := (r.Width - spacingTotalSize) / r.ChildCount
	childXPosition := r.X + r.Padding + (child * childWidth) + (child * r.Spacing)
	return DynaRect{
		X:          childXPosition,
		Y:          r.Y + r.Padding,
		Width:      childWidth,
		Height:     r.Height - 2*r.Padding,
		Padding:    r.Padding,
		Spacing:    r.Spacing,
		ChildCount: r.ChildCount,
	}
}

// WithPadding - returns new DynaRect with updated Padding value
func (r DynaRect) WithPadding(padding float32) DynaRect {
	return DynaRect{
		X:          r.X,
		Y:          r.Y,
		Width:      r.Width,
		Height:     r.Height,
		Padding:    padding,
		Spacing:    r.Spacing,
		ChildCount: r.ChildCount,
	}
}

// WithSpacing - returns new DynaRect with updated Spacing
func (r DynaRect) WithSpacing(spacing float32) DynaRect {
	return DynaRect{
		X:          r.X,
		Y:          r.Y,
		Width:      r.Width,
		Height:     r.Height,
		Padding:    r.Padding,
		Spacing:    spacing,
		ChildCount: r.ChildCount,
	}
}

// WithScrollOffset - method allow us to generate Rectangle position for scrollable content
func (r DynaRect) WithScrollOffset(offset rl.Vector2) DynaRect {
	return DynaRect{
		X:          r.X + offset.X,
		Y:          r.Y + offset.Y,
		Width:      r.Width,
		Height:     r.Height,
		Padding:    r.Padding,
		Spacing:    r.Spacing,
		ChildCount: r.ChildCount,
	}
}

// NextChild - increase child count of DynaRectIterator instance by one
//
// It is used to calculate actual position of child element inside parent rect
func (dri *DynaRectIterator) NextChild() {
	dri.ActualChild++
}

// ChildRect - return raylib.Rectangle for actual child element of an DynaRectIterator instance
func (dri *DynaRectIterator) ChildRect() rl.Rectangle {
	childsInRow := int(dri.ParentRect.Width / dri.ChildSize.X)
	childPosition := dri.ActualChild % childsInRow
	childRow := dri.ActualChild / childsInRow
	return rl.Rectangle{
		X:      dri.ParentRect.X + float32(childPosition)*dri.ChildSize.X,
		Y:      dri.ParentRect.Y + float32(childRow)*dri.ChildSize.Y,
		Width:  dri.ChildSize.X,
		Height: dri.ChildSize.Y,
	}
}

// ChildDynaRect - returns DynaRect for actual child of element of an DynaRectIterator instance
func (dri *DynaRectIterator) ChildDynaRect() DynaRect {
	childsInRow := int(dri.ParentRect.Width / dri.ChildSize.X)
	childPosition := dri.ActualChild % childsInRow
	childRow := dri.ActualChild / childsInRow
	return DynaRect{
		X:          dri.ParentRect.X + float32(childPosition)*dri.ChildSize.X,
		Y:          dri.ParentRect.Y + float32(childRow)*dri.ChildSize.Y,
		Width:      dri.ChildSize.X,
		Height:     dri.ChildSize.Y,
		Padding:    dri.ParentRect.Padding,
		Spacing:    dri.ParentRect.Spacing,
		ChildCount: dri.ParentRect.ChildCount,
	}
}
