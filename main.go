package main

import (
	"chronocut/utils"
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
	views := viewSelectorTabs(app, window)

	UI := container.NewHBox(
		views,
	)

	window.SetContent(UI)
	window.Resize(fyne.NewSize(600, 600))
	window.ShowAndRun()
}

func videoPage(app fyne.App, window fyne.Window) *fyne.Container {

	fileNameLabel := widget.NewLabel("Selected file: ")

	selectButton := widget.NewButton("Select file", func() {
		dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			utils.HandleError("File open error: ", err)

			fileNameLabel.SetText("Selected file: " + reader.URI().Path())

		}, window).Show()
	})

	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	videoUI := container.NewVBox(
		fileNameLabel,
		selectButton,
		quitButton,
	)

	return videoUI
}

func viewSelectorTabs(app fyne.App, window fyne.Window) *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Video", theme.FileVideoIcon(), videoPage(app, window)),
		container.NewTabItemWithIcon("Download", theme.DownloadIcon(), widget.NewLabel("")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("")),
	)
	return tabs
}
