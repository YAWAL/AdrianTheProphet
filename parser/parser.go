package parser

import (
	"os"
	"log"
	"encoding/csv"
)

func ParseDataset(path string) (records [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Error during opening file: ", err)
	}
	defer file.Close()

	data := csv.NewReader(file)
	//data.FieldsPerRecord = 8

	records, err = data.ReadAll()
	if err != nil {
		log.Println("Error during reading file: ", err)
		return records, err
	}
	if len(records) < 1 {
		log.Println("Empty file")
	}
	return records, nil
}
