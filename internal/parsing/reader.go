package parsing

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type JsonReader struct{}

// loads the config.json File in the Current Working Directory and try to parse it into EntityFactoryConfig type
func (cr *JsonReader) Read(file string) (EntityFactoryConfig, error) {
	configFile, readFileErr := ioutil.ReadFile(file)
	if readFileErr != nil {
		return EntityFactoryConfig{
			ConnectionString: "",
		}, errors.New("error on read config file " + file)
	}

	var config EntityFactoryConfig
	parseConfigErr := json.Unmarshal(configFile, &config)
	if parseConfigErr != nil {
		return EntityFactoryConfig{
			ConnectionString: "",
		}, errors.New("error on parse config file " + file)
	}

	return config, nil
}
