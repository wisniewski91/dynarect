# DynaRect

Work in progress

This package allow to simple create raylib.Rectangles relative to eachother.

Example of usage:
```
// Create DynaRect object
    rect := dynarect.DynaRect{
		X:          rect.X,
		Y:          rect.Y,
		Width:      rect.Width,
		Height:     rect.Height,
		Padding:    10,
		Spacing:    10,
		ChildCount: 1,
	}
    // Render panel with Rect position and size
    rl.Panel(rect.GetRect(), "Example panel")

    // Create rect for rows based on first rect. Set up different height and child count with smaller Padding and Spacing
	rowRect := localRect.GetChild(0).WithHeight(30).WithChildCount(2).WithPadding(2).WithSpacing(5)

// Create first button position, getting first child position including Padding inside container and space between elements in a row. 
	gui.Button(rowRect.GetChild(0).GetRect(), "Button 1")
	gui.Label(rowRect.GetChild(1).GetRect(), "Button 2")

// Move rowRect to the next row by offsetting Y position by height + Spacing 
	rowRect = rowRect.NexRow()

// Render next Buttons
	gui.Button(rowRect.GetChild(0).GetRect(), "Button 3")
	gui.Label(rowRect.GetChild(1).GetRect(), "Button 4")
```