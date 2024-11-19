package main

import (
	"fyne.io/fyne/app"
	"github.com/joho/godotenv"
	"github.com/mhson281/currency-converter/ui"
)

func main() {
	godotenv.Load()

	a := app.New()
	w := a.NewWindow("Currency Converter")

	// build the ui
	w.SetContent(ui.BuildUI())
}
