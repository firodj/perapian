package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("PerAPIan")

	hello := widget.NewLabel("Pet Store")

	tree := widget.NewTree(
		// ID
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "": // ROOT
				return []widget.TreeNodeID{"info", "paths"}
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
		},
		// BRANCH
		func(id widget.TreeNodeID) bool {
			if id == "" || id == "paths" {
				return true
			}
			if strings.HasPrefix(id, "/") {
				return true
			}
			return false
		},
		// CREATE
		func(branch bool) fyne.CanvasObject {
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
		},
		nil)
	tree.UpdateNode = // UPDATE
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			rootContainer := o.(*fyne.Container)
			leftContainer := rootContainer.Objects[0].(*fyne.Container)
			//rightContainer := rootContainer.Objects[1].(*fyne.Container)
			fileIcon := leftContainer.Objects[0].(*widget.Icon)

			nameLabel := leftContainer.Objects[1].(*widget.Label)

			text := id
			if branch {
				if tree.IsBranchOpen(id) {
					fileIcon.SetResource(theme.FolderOpenIcon())
				} else {
					fileIcon.SetResource(theme.FolderIcon())
				}
			} else {
				fileIcon.SetResource(theme.CancelIcon())
			}
			nameLabel.SetText(text)
		}

	sideMenu := container.NewBorder(
		//  Top
		container.NewVBox(hello),
		nil, nil, nil,
		// Fill
		tree,
	)

	rightContent := container.NewStack()

	infoPage := container.New(layout.NewFormLayout(),
		widget.NewLabel("openapi"),
		widget.NewEntry(),
		widget.NewLabel("version"),
		widget.NewEntry(),
		widget.NewLabel("title"),
		widget.NewEntry(),
		widget.NewLabel("license"),
		widget.NewSelect([]string{"MIT"}, nil),
	)

	magicButton := widget.NewButton("Hi!", nil)
	magicButton.OnTapped = func() {
		magicButton.SetText("Hi Welcome :)")
	}

	tree.OnSelected = func(id widget.TreeNodeID) {
		fmt.Println("id", id)
		if id == "info" {
			rightContent.Objects = nil
			rightContent.Add(infoPage)
			// rightContent.Refresh()
		} else {
			rightContent.Objects = nil
			rightContent.Refresh()
		}
	}

	w.SetContent(container.NewHSplit(
		sideMenu,
		rightContent,
	))
	w.Resize(fyne.Size{Width: 1024, Height: 768})

	w.ShowAndRun()
}
