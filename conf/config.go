package conf

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

const FileName = "large_data.json"

type Config struct {
	NumberSets []NumberSet
}

type NumberSet struct {
	A int `json:"a" validate:"min=-10,max=10"`
	B int `json:"b" validate:"min=-10,max=10"`
}

func New() (Config, error) {
	config := Config{}

	file, err := os.Open(FileName)
	if err != nil {
		return config, errors.Wrap(err, "open")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bytes, &config.NumberSets)
	if err != nil {
		return config, errors.Wrap(err, "unmarshal config")
	}

	v := validator.New()

	for _, numberSet := range config.NumberSets {
		err = v.Struct(numberSet)
		if err != nil {
			return config, errors.Wrap(err, "validate config")
		}
	}

	return config, nil
}
