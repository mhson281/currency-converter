package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	apiURL = "https://openexchangerates.org/api/latest.json"
)

type RatesResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Unable to load .env file due to the following error", "error", err)
		return err
	}
	return nil
}

func FetchRates() (map[string]float64, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		err := errors.New("API_KEY is not found in .env")
		slog.Error("Missing API_KEY", "error", err)
		return nil, err
	}
	url := fmt.Sprintf("%s?app_id=%s", apiURL, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to fetch rates from API", "url", url, "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data RatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		slog.Error("Failed to decode API response", "error", err)
		return nil, err
	}

	slog.Info("Succesfully fetched rates from API", "base", data.Base, "num_rates", len(data.Rates))

	return data.Rates, nil
}
