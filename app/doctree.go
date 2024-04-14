package app

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type docTree struct {
	widget.Tree
}

func NewDocTree() *docTree {
	t := &docTree{}
	t.ChildUIDs = t.OnChildUIDs
	t.IsBranch = t.OnIsBranch
	t.CreateNode = t.OnCreateNode
	t.UpdateNode = t.OnUpdateNode
	t.ExtendBaseWidget(t)
	return t
}

func SampleOnChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	switch id {
	case "paths":
		return []widget.TreeNodeID{"/pets", "/foods"}
	case "/pets":
		return []widget.TreeNodeID{"GET/pets", "POST/pets", "/pets/{petId}"}
	case "/foods":
		return []widget.TreeNodeID{"GET/foods", "POST/foods"}
	case "/pets/{petId}":
		return []widget.TreeNodeID{"GET/pets/{petId}"}
	}

	return []string{}
}

func (t *docTree) OnChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	if id == "" { // ROOT
		return []widget.TreeNodeID{"info", "paths"}
	}

	if GetMainApp().Doc == nil {
		return SampleOnChildUIDs(id)
	}

	return []string{}
}

func (t *docTree) OnIsBranch(id widget.TreeNodeID) bool {
	if id == "" || id == "paths" {
		return true
	}
	if strings.HasPrefix(id, "/") {
		return true
	}
	return false
}

func (t *docTree) OnCreateNode(branch bool) fyne.CanvasObject {
	icon := widget.NewIcon(theme.CancelIcon())

	// if branch
	return container.NewBorder(
		nil,
		nil,
		container.NewHBox(
			icon,
			widget.NewLabel("Left"),
		),
		nil,
	)
}

func (t *docTree) OnUpdateNode(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	rootContainer := o.(*fyne.Container)
	leftContainer := rootContainer.Objects[0].(*fyne.Container)
	//rightContainer := rootContainer.Objects[1].(*fyne.Container)
	fileIcon := leftContainer.Objects[0].(*widget.Icon)

	nameLabel := leftContainer.Objects[1].(*widget.Label)

	text := id
	if branch {
		if t.IsBranchOpen(id) {
			fileIcon.SetResource(theme.FolderOpenIcon())
		} else {
			fileIcon.SetResource(theme.FolderIcon())
		}
	} else {
		fileIcon.SetResource(theme.CancelIcon())
	}
	nameLabel.SetText(text)
}
