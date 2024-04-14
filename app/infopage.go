package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/getkin/kin-openapi/openapi3"
)

// https://opensource.org/license
var SPDXIds = []string{
	"MIT",
	"Apache-2.0",
	"MPL-2.0",
	"LGPL-3.0-only",
	"GPL-3.0-only",
	"BSD-3-Clause",
	"CDDL-1.0",
	"EPL-2.0",
	"Unlicense",
}

func CreateInfoPage() *fyne.Container {
	doc := GetMainApp().Doc
	if doc == nil {
		return container.NewStack()
	}

	if doc.Info.License == nil {
		doc.Info.License = &openapi3.License{}
	}

	licenseBind := binding.BindString(&doc.Info.License.Name)
	licenseSelect := widget.NewSelectEntry(SPDXIds)
	licenseSelect.OnChanged = func(text string) {
		err := licenseBind.Set(text)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	}
	text, err := licenseBind.Get()
	if err == nil {
		licenseSelect.SetText(text)
	}

	infoPage := container.New(layout.NewFormLayout(),
		widget.NewLabel("openapi"),
		widget.NewEntryWithData(binding.BindString(&doc.OpenAPI)),
		widget.NewLabel("version"),
		widget.NewEntryWithData(binding.BindString(&doc.Info.Version)),
		widget.NewLabel("title"),
		widget.NewEntryWithData(binding.BindString(&doc.Info.Title)),
		widget.NewLabel("license"),
		licenseSelect,
	)

	return infoPage
}
