package main

import (
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"github.com/firodj/perapian/app"
)

func main() {
	a := app.NewApp()

	err := a.Load("./samples/petstore.json")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		dialog.ShowError(err, a.MainWindow)
	}

	a.Run()
}
