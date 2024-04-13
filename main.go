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
	"github.com/firodj/perapian/custom"
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

	pathPageCreate := func() *fyne.Container {
		treeParams := widget.NewList(
			func() int {
				return 2
			},
			func() fyne.CanvasObject {
				return container.New(layout.NewFormLayout(),
					widget.NewLabel("name"),
					container.NewBorder(nil, nil, nil, widget.NewCheck("required", nil), widget.NewEntry()),
					widget.NewLabel("in"), widget.NewSelect([]string{"query", "path", "header", "cookie"}, nil),
					widget.NewLabel("description"), widget.NewEntry(),
				)
			},
			func(id int, o fyne.CanvasObject) {
				rootContainer := o.(*fyne.Container)
				nameContainer := rootContainer.Objects[1].(*fyne.Container)

				if nameEntry, ok := nameContainer.Objects[0].(*widget.Entry); ok {
					text := ""
					if id == 0 {
						text = "limit"
					} else if id == 1 {
						text = "page"
					}
					nameEntry.SetText(text)
				}
			},
		)
		pathPage := container.NewBorder(
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("summary"), widget.NewEntry(),
				widget.NewLabel("operationId"), widget.NewEntry(),
				widget.NewLabel("tags"), widget.NewSelectEntry([]string{}),
			),
			nil, nil, nil,
			container.NewBorder(
				container.NewHBox(widget.NewLabel("Parameters"),
					layout.NewSpacer(),
					custom.NewContextMenuButton(":", fyne.NewMenu("A", fyne.NewMenuItem("B", nil))),
				),
				nil, nil, nil,
				treeParams,
			),
		)
		return pathPage
	}

	magicButton := widget.NewButton("Hi!", nil)
	magicButton.OnTapped = func() {
		magicButton.SetText("Hi Welcome :)")
	}

	menuItem1 := fyne.NewMenuItem("A", nil)
	menuItem2 := fyne.NewMenuItem("B", nil)
	menuItem3 := fyne.NewMenuItem("C", nil)
	menu := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)

	schemaPageCreate := func() *fyne.Container {
		return container.NewBorder(
			nil, nil, nil, nil,
			widget.NewTree(
				func(id widget.TreeNodeID) []widget.TreeNodeID {
					switch id {
					case "":
						return []widget.TreeNodeID{"root"}
					case "root":
						return []widget.TreeNodeID{"id", "name", "tag"}
					}
					return []widget.TreeNodeID{}
				},
				func(id widget.TreeNodeID) bool {
					switch id {
					case "":
						return true
					case "root":
						return true
					}
					return false
				},
				func(branch bool) fyne.CanvasObject {
					return container.NewHBox(
						widget.NewLabel("Name"),
						layout.NewSpacer(),
						widget.NewSelect([]string{"boolean", "integer", "integer/int32", "integer/int64", "number", "number/float", "number/double", "string", "string/password", "string/uuid", "object", "array"}, nil),
						custom.NewContextMenuButton(":", menu),
					)
				},
				func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
					rootContainer := o.(*fyne.Container)

					labelName := rootContainer.Objects[0].(*widget.Label)
					labelName.SetText(id)
					buttonType := rootContainer.Objects[2].(*widget.Select)
					// buttonContext := rootContainer.Objects[3].(*custom.ContextMenuButton)

					switch id {
					case "root":
						buttonType.SetSelected("object")
					case "id":
						buttonType.SetSelected("integer/int64")
					default:
						buttonType.SetSelected("string")
					}
				},
			),
		)

	}

	tree.OnSelected = func(id widget.TreeNodeID) {
		fmt.Println("id", id)
		switch id {
		case "info":
			rightContent.Objects = nil
			rightContent.Add(infoPage)
			// rightContent.Refresh()

		case "GET/pets":
			rightContent.Objects = nil
			rightContent.Add(pathPageCreate())

		case "/pets":
			rightContent.Objects = nil
			rightContent.Add(schemaPageCreate())

		default:
			rightContent.Objects = nil
			rightContent.Refresh()
		}
	}
	tree.OpenAllBranches()

	w.SetContent(container.NewHSplit(
		sideMenu,
		rightContent,
	))
	w.Resize(fyne.Size{Width: 1024, Height: 768})

	w.ShowAndRun()
}
