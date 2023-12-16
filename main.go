package main

import (
	"chronocut/ffmpeg"
	"chronocut/utils"
	"fmt"
	"os/exec"

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
	filepath := "/home/dft/Yakuza 0 - Friday Night.mp4"
	ffcmd := exec.Command("ffmpeg", "-i", filepath, "-f", "null", "-")

	out, err := ffcmd.CombinedOutput()
	utils.HandleError("error starting ffmpeg", err)

	duration, err := ffmpeg.GetDuration(string(out))
	utils.HandleError("Couldnt get duration", err)
	fmt.Println(duration[:8])

	window.SetContent(UI)
	window.Resize(fyne.NewSize(600, 600))
	window.ShowAndRun()
}

func videoPage(app fyne.App, window fyne.Window) *fyne.Container {
	fileNameLabel := widget.NewLabel("Selected file: ")

	selectButton := widget.NewButton("Select file", func() {
		dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			utils.HandleError("File open error", err)

			filepath := reader.URI().Path()
			fileNameLabel.SetText("Selected file: " + reader.URI().Path())
			ffcmd := exec.Command("ffmpeg", "-i", filepath, "-f", "null", "-")

			out, err := ffcmd.CombinedOutput()
			utils.HandleError("error starting ffmpeg", err)

			duration, err := ffmpeg.GetDuration(string(out))
			utils.HandleError("Couldnt get duration", err)
			fmt.Println(duration[:8])

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
