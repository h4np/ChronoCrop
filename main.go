package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Chronocut")

	window.SetContent(renderUI(app, window))
	window.Resize(fyne.NewSize(600, 600))
	window.ShowAndRun()
}

func renderUI(app fyne.App, window fyne.Window) *fyne.Container {
	tabs := topTabs()

	fileNameLabel := widget.NewLabel("Selected file: ")

	selectButton := widget.NewButton("Select file", func() {
		dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if reader != nil {
				fileNameLabel.SetText("Selected file: " + reader.URI().Path())
				fmt.Println("reader nil")
				reader.Close()
			}
			fileNameLabel.SetText("Selected file: " + reader.URI().Path())
		}, window).Show()
	})
	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	UI := container.NewVBox(
		tabs,
		fileNameLabel,
		selectButton,
		quitButton,
	)

	return UI
}

func topTabs() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Video", theme.FileVideoIcon(), widget.NewLabel("")),
		container.NewTabItemWithIcon("Download", theme.DownloadIcon(), widget.NewLabel("")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("")),
	)
	return tabs
}
