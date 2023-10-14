package db

import (
	"encoding/json"
	"io/ioutil"
)

type Genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (Genesis, error) {
	var (
		loadedGenesis Genesis
	)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}
