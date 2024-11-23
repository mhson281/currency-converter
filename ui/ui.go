package ui

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/mhson281/currency-converter/api"
)

func addCommas(value float64) string {
	// Format the float to two decimal places
	plain := strconv.FormatFloat(value, 'f', 2, 64)
	parts := strings.Split(plain, ".") // Split into whole number and decimal part

	// Add commas to the whole number part
	whole := parts[0]
	n := len(whole)
	withCommas := ""

	for i, digit := range whole {
		if i > 0 && (n-i)%3 == 0 {
			withCommas += ","
		}
		withCommas += string(digit)
	}

	// Reattach the decimal part, if present
	if len(parts) > 1 {
		withCommas += "." + parts[1]
	}

	return withCommas
}


func BuildUI() fyne.CanvasObject {
	// Create input fields for amount and currency selection
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Enter amount")

	fromCurrency := widget.NewSelect([]string{"USD", "EUR", "GBP", "CAD", "VND"}, nil)
	fromCurrency.PlaceHolder = "From"

	toCurrency := widget.NewSelect([]string{"USD", "EUR", "GBP", "CAD", "VND"}, nil)
	toCurrency.PlaceHolder = "To"

	resultLabel := widget.NewLabel("Result: ")

	// Create a button to perform the conversion
	convertButton := widget.NewButton("Convert", func() {
		// Validate the amount input
		if amountEntry.Text == "" {
			resultLabel.SetText("Amount cannot be empty")
			return
		}

		// Parse the amount entered by the user
		amount, err := strconv.ParseFloat(amountEntry.Text, 64)
		if err != nil {
			resultLabel.SetText("Invalid amount")
			return
		}

		if amount < 0 {
			resultLabel.SetText("Please enter a value greater than 0")
			return
		}

		// Validate currency selections
		if fromCurrency.Selected == "" || toCurrency.Selected == "" {
			resultLabel.SetText("Please select both currencies")
			return
		}

		// Fetch the latest currency rates
		rates, err := api.FetchRates()
		if err != nil {
			resultLabel.SetText("Unable to fetch rates")
			return
		}

		// Perform the currency conversion
		rateFrom, ok1 := rates[fromCurrency.Selected]
		rateTo, ok2 := rates[toCurrency.Selected]
		if !ok1 || !ok2 {
			resultLabel.SetText("Invalid currency selection")
			return
		}

		converted := amount * (rateTo / rateFrom)

		var formattedResult string
		if converted > 1000 {
			formattedResult = addCommas(converted)
		} else {
			formattedResult = fmt.Sprintf("%.2f", converted)
		}

		resultLabel.SetText(fmt.Sprintf("Result: %s %s", formattedResult, toCurrency.Selected))
	})

	// Arrange UI components in a vertical box layout
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

