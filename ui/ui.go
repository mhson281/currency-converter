package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mhson281/currency-converter/api"
)

func BuildUI() fyne.CanvasObject {
	// Create a new application
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Enter amount")

	fromCurrency := widget.NewSelect([]string{"USD", "EUR", "GBP", "VND"}, nil)
	fromCurrency.PlaceHolder = "From"

	toCurrency := widget.NewSelect([]string{"USD", "EUR", "GBP", "VND"}, nil)
	toCurrency.PlaceHolder = "To"

	resultLabel := widget.NewLabel("Result: ")

	convertButton := widget.NewButton("Convert", func() {
		amount, err := strconv.ParseFloat(amountEntry.Text, 64)

		if err != nil {
			resultLabel.SetText("Invalid amount")
		}

		rates, err := api.FetchRates()
		if err != nil {
			resultLabel.SetText("Unable to fetch rates")
			return
		}

		rateFrom, ok1 := rates[fromCurrency.Selected]
		rateTo, ok2 := rates[toCurrency.Selected]
		if !ok1 || !ok2 {
			resultLabel.SetText("Invalid currency selection")
			return
		}

		converted := amount * (rateTo / rateFrom)
		resultLabel.SetText(fmt.Sprintf("Result: %.2f %s", converted, toCurrency.Selected))
	})

	// Layout
	form := container.NewVBox(
		widget.NewLabel("Currency Converter"),
		amountEntry,
		fromCurrency,
		toCurrency,
		convertButton,
		resultLabel,
	)

	return form
}
