package excel

import (
	"encoding/csv"
	"fmt"
	"os"
	"scrapeQuotes/model"
	"strings"
)

func ExportStructToCSV(data []model.Quote) error {
	// Create the CSV file
	file, err := os.Create("scrapedQuotes.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	header := []string{"Quote", "Author", "Tags"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write the data
	for _, item := range data {
		fmt.Println(item)
		tags := strings.Join(item.Tags, ", ")
		row := []string{item.Text, item.Author, tags}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
