package parsingconfig

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

type ConfigReader struct{}

// loads the config.json File in the Current Working Directory and try to parse it into EntityFactoryConfig type
func (cr *ConfigReader) Read() (EntityFactoryConfig, error) {
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		return EntityFactoryConfig{
			ConnectionString: "",
		}, errors.New("error on get current working directory")
	}

	configFilePath := path.Join(cwd, "config.json")
	configFile, readFileErr := ioutil.ReadFile(configFilePath)
	if readFileErr != nil {
		return EntityFactoryConfig{
			ConnectionString: "",
		}, errors.New("error on read config file " + configFilePath)
	}

	var config EntityFactoryConfig
	parseConfigErr := json.Unmarshal(configFile, &config)
	if parseConfigErr != nil {
		return EntityFactoryConfig{
			ConnectionString: "",
		}, errors.New("error on parse config file " + configFilePath)
	}

	return config, nil
}
