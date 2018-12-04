package manager

import (
	"encoding/json"
	"io/ioutil"

	"../basic"
)

func LoadConfiguration() (*basic.Configuration, error) {
	configurationFile, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return &basic.Configuration{}, err
	}

	configuration := &basic.Configuration{}

	err = json.Unmarshal(configurationFile, &configuration)

	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
