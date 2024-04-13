package main

import (
	"fmt"
	"os"
	"path"

	"github.com/roemer/gotaskr"
	"github.com/roemer/gotaskr/execr"
)

// Build variables
var version = "0.0.1"

// Internal variables
var outputDirectory = ".build-output"
var windowsBuildOutput = "perapian.exe"
var macosBuildOutput = "perapian.app"

func main() {
	os.Exit(gotaskr.Execute())
}

func init() {
	gotaskr.Task("Setup:Fyne-Cmd", func() error {
		version := "latest" // latest or develop
		return execr.Run(true, "go", "install", fmt.Sprintf("fyne.io/fyne/v2/cmd/fyne@%s", version))
	})

	gotaskr.Task("Run", func() error {
		return execr.Run(true, "go", "run", "-v", "main.go")
	})

	gotaskr.Task("Compile:Windows", func() error {
		os.Setenv("CGO_ENABLED", "1")
		os.Setenv("CC", "x86_64-w64-mingw32-gcc")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		fynePath := path.Join(homeDir, "go/bin/fyne")
		if err := execr.Run(true, fynePath, "package", "-os", "windows", "--appVersion", version, "-icon", "icon.png"); err != nil {
			return nil
		}
		os.Mkdir(outputDirectory, os.ModePerm)
		return os.Rename(windowsBuildOutput, path.Join(outputDirectory, windowsBuildOutput))
	}).DependsOn("Setup:Fyne-Cmd")

	gotaskr.Task("Compile:MacOS", func() error {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		fynePath := path.Join(homeDir, "go/bin/fyne")
		if err := execr.Run(true, fynePath, "package", "-os", "darwin", "--appVersion", version, "-icon", "icon.png"); err != nil {
			return nil
		}
		os.Mkdir(outputDirectory, os.ModePerm)
		return os.Rename(macosBuildOutput, path.Join(outputDirectory, macosBuildOutput))
	}).DependsOn("Setup:Fyne-Cmd")
}
