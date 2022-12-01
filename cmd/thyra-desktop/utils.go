package main

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2"
)

func openURL(a *fyne.App, urlToOpen string) {
	u, err := url.Parse(urlToOpen)
	if err != nil {
		log.Fatal(err)
	}
	(*a).OpenURL(u)
}
