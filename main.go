package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
    app := app.New()
    window := app.NewWindow("Chronocut")

    fileNameLabel := widget.NewLabel("Selected file: ")

    selectButton := widget.NewButton("Select file", func() {
        dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
            if err != nil {
                fmt.Println(err)
                return
            }
            fileNameLabel.SetText("Selected file: " + reader.URI().Path())
        }, window).Show()

    })


    content := container.NewVBox(
        widget.NewLabel("Chronocut"),
        fileNameLabel,
        selectButton,
        widget.NewButton("Quit", func() {
            app.Quit()
        }),
    )

    window.SetContent(content)
    window.Resize(fyne.NewSize(600, 600))
    window.ShowAndRun()
}
