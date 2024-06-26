package app

import (
	"fmt"
	"slices"
	"strings"

	"github.com/firodj/perapian/common"
	"github.com/getkin/kin-openapi/openapi3"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	FileName   string
	Doc        *openapi3.T
	GUI        fyne.App
	MainWindow fyne.Window

	DocTree      *docTree
	RightContent *fyne.Container
}

var mainApp *App = nil

func NewApp() *App {
	a := &App{}
	mainApp = a
	a.Init()
	return a
}

func GetMainApp() *App { return mainApp }

func (a *App) Init() {
	gui := fyneApp.New()
	a.GUI = gui
	w := gui.NewWindow("PerAPIan")
	a.MainWindow = w

	a.DocTree = NewDocTree()
	a.DocTree.OnSelected = a.onDocTreeSelected
	a.RightContent = container.NewStack()

	a.MainWindow.SetContent(container.NewHSplit(
		a.DocTree,
		a.RightContent,
	))
	a.MainWindow.Resize(fyne.Size{Width: 1024, Height: 768})
}

func (a *App) Run() {
	a.MainWindow.ShowAndRun()
}

func (a *App) Load(filename string) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(filename)
	if err != nil {
		return err
	}
	a.Doc = doc
	a.DocTree.Refresh()
	a.DocTree.OpenAllBranches()
	return nil
}

func (a *App) onDocTreeSelected(id widget.TreeNodeID) {
	a.RightContent.Objects = nil
	pathparts := strings.SplitN(id, "/", 2)
	if id == "info" {
		a.RightContent.Add(CreateInfoPage())
	} else if strings.HasPrefix(id, "/") {
		fmt.Printf("select path: %s\n", id)
	} else {
		if len(pathparts) == 2 && slices.Contains(common.HttpMethods, pathparts[0]) {
			meth := pathparts[0]
			path := "/" + pathparts[1]
			fmt.Printf("ops: %s, path: %s\n", meth, path)
		}
	}
}
