package main

import (
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}