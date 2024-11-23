package main

import (
	"log/slog"

	"fyne.io/fyne/app"
	"fyne.io/fyne"
	"github.com/joho/godotenv"
	"github.com/mhson281/currency-converter/ui"
)

func main() {
	godotenv.Load()

	slog.Info("Starting currency converter application")

	a := app.New()
	w := a.NewWindow("Currency Converter")

	// build the ui
	content := ui.BuildUI()
	if _, ok := content.(fyne.CanvasObject); !ok {
		panic("BuildUI() does not return fyne.CanvasObject")
	}

	w.SetContent(ui.BuildUI())
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

