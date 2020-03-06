package parsing_config

import (
	parsingconfig "github.com/nodejayes/entity-factory/internal/parsing/config"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path"
)

func getConfigPath(fileName string) string {
	cwd, _ := os.Getwd()
	return path.Join(cwd, fileName)
}

func createValidConfigFile(fileName string) {
	_ = ioutil.WriteFile(getConfigPath(fileName), []byte("{\"connectionString\":\"testconnectionstring\"}"), 0644)
}

func createInvalidConfigFile(fileName string) {
	_ = ioutil.WriteFile(getConfigPath(fileName), []byte("\"connection_string\":\"testconnectionstring\"}"), 0644)
}

func deleteConfigFile(fileName string) {
	filePath := getConfigPath(fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return
	}
	_ = os.Remove(filePath)
}

var _ = ginkgo.Describe("ConfigReader Tests", func() {
	ginkgo.Context("[Method]: Read", func() {
		ginkgo.It("can read a config.json in the Working Directory", func() {
			createValidConfigFile("config.json")
			configReader := parsingconfig.ConfigReader{}
			config, err := configReader.Read()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(config.ConnectionString).To(gomega.Equal("testconnectionstring"))
			deleteConfigFile("config.json")
		})
		ginkgo.It("return a error when config.json not found", func() {
			configReader := parsingconfig.ConfigReader{}
			_, err := configReader.Read()
			gomega.Expect(err.Error()).To(gomega.ContainSubstring("error on read config file "))
		})
		ginkgo.It("return a error when config.json is invalid", func() {
			createInvalidConfigFile("config.json")
			configReader := parsingconfig.ConfigReader{}
			_, err := configReader.Read()
			gomega.Expect(err.Error()).To(gomega.ContainSubstring("error on parse config file "))
			deleteConfigFile("config.json")
		})
	})
})
