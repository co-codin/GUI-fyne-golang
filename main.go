package main

import (
	"encoding/json"
	"net/http"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type randomFact struct {
	Text string `json:"text"`
}

func getRandomFact() (randomFact, error) {
	var fact randomFact
	resp, err : = client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return randomFact{}, err
	}
	defer resp.Body.close()

	err = json.NewDecoder(resp.Body).Decode(&fact)

	if err != nil {
		return randomFact{}, err
	}
	return fact, nil
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
