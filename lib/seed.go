package autolunar

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Seed [][]uint8

func ReadSeed(name string) (Seed, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("[autolunar] cannot read empty seed"))
	}
	f, err := os.Open(fmt.Sprintf("./seeds/%s.csv", name))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var seed Seed
	for key, record := range records {
		if key == 0 {
			continue
		}
		seed = append(seed, []uint8{})
		for _, coordinate := range record {
			c, err := strconv.Atoi(coordinate)
			if err != nil {
				return nil, err
			}
			seed[key-1] = append(seed[key-1], uint8(c))
		}
	}

	return seed, nil
}
