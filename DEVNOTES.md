##  Drag

```go
type ContextMenuButton struct {
    ...
    startDragOff   *fyne.Position
	currentDragPos fyne.Position

    OnDragEnd func(pos fyne.Position) `json:"-"`

    parent fyne.CanvasObject
}

var _ fyne.Draggable = (*ContextMenuButton)(nil)

func (b *ContextMenuButton) DragEnd() {
	b.startDragOff = nil
	if b.OnDragEnd != nil {
		b.OnDragEnd(b.currentDragPos)
	}
}

func (b *ContextMenuButton) Dragged(e *fyne.DragEvent) {
	if b.startDragOff == nil {
		b.currentDragPos = b.Position().Add(e.Position)
		start := e.Position.Subtract(e.Dragged)
		b.startDragOff = &start
	} else {
		b.currentDragPos = b.currentDragPos.Add(e.Dragged)

	}

	if b.parent != nil {
		b.parent.Move(b.currentDragPos.Subtract(b.startDragOff))
	}
}

func (b *ContextMenuButton) SetParent(p fyne.CanvasObject) {
	b.parent = p
}
```