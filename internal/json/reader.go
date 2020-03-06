package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// read JSON Files
type JsonReader struct{}

// loads the config.json File in the Current Working Directory and try to parse it into EntityFactoryConfig type
func (cr *JsonReader) Read(file string, target interface{}) error {
	configFile, readFileErr := ioutil.ReadFile(file)
	if readFileErr != nil {
		return errors.New("error on read config file " + file)
	}

	parseConfigErr := json.Unmarshal(configFile, target)
	if parseConfigErr != nil {
		return errors.New("error on parse config file " + file)
	}

	return nil
}
