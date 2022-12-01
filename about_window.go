package main

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createAboutWindow(a *fyne.App) fyne.Window {
	u, err := url.Parse("https://massa.net")
	if err != nil {
		log.Fatal(err)
	}

	w := (*a).NewWindow("About")

	titleText := widget.NewLabel("Thyra Desktop")
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle = fyne.TextStyle{Bold: true}

	versionText := widget.NewLabel("Version 0.0.1")
	versionText.Alignment = fyne.TextAlignCenter

	thyraServerVersionText := widget.NewLabel("Thyra Server Version " + getThyraServerVersion())
	thyraServerVersionText.Alignment = fyne.TextAlignCenter

	separator := widget.NewSeparator()
	authorText := widget.NewLabel("by Massa")
	authorText.Alignment = fyne.TextAlignCenter

	massaNetHyperlink := widget.NewHyperlink("massa.net", u)
	massaNetHyperlink.Alignment = fyne.TextAlignCenter

	closeButton := widget.NewButton("Close", func() {
		w.Close()
	})
	closeButton.Alignment = widget.ButtonAlignCenter

	objects := []fyne.CanvasObject{
		titleText,
		versionText,
		separator,
		thyraServerVersionText,
		authorText,
		massaNetHyperlink,
		closeButton,
	}
	c := container.New(layout.NewVBoxLayout(), objects...)

	w.SetContent(c)
	return w
}
