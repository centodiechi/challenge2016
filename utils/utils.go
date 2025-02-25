package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

var AreaMap = make(map[string]bool)

func LoadCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		cityCode := record[0]
		provinceCode := record[1]
		countryCode := record[2]

		AreaMap[countryCode] = true
		AreaMap[fmt.Sprintf("%s-%s", provinceCode, countryCode)] = true
		AreaMap[fmt.Sprintf("%s-%s-%s", cityCode, provinceCode, countryCode)] = true
	}

	return nil
}
