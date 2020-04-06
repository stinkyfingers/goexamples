package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

// data from: https://github.com/fivethirtyeight/data/blob/master/candy-power-ranking/candy-data.csv

// Candy represents a single candy object
type Candy struct {
	Competitor   string  `json:"name"`
	Sugarpercent float64 `json:"sugarPercent"`
	Pricepercent float64 `json:"pricePercent"`
}

func main() {
	candies, err := parseCSV("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = writeJSON(candies)
	if err != nil {
		log.Fatal(err)
	}
}

// parseCSV opens and reads the parameterized filename.
// It then creates a Candy instance and appends it to the
// output Candy slice
func parseCSV(filename string) ([]Candy, error) {
	var candies []Candy
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	header := true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if header {
			header = false
			continue
		}

		sugarpercent, err := strconv.ParseFloat(line[10], 64)
		if err != nil {
			return nil, err
		}
		pricepercent, err := strconv.ParseFloat(line[11], 64)
		if err != nil {
			return nil, err
		}

		candy := Candy{
			Competitor:   line[0],
			Sugarpercent: sugarpercent,
			Pricepercent: pricepercent,
		}
		candies = append(candies, candy)
	}
	return candies, nil
}

// writeJSON creates a file and encodes a slice of
// Candy to it.
func writeJSON(candies []Candy) error {
	f, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(candies)
}
