package parsing

import (
	"github.com/nodejayes/entity-factory/internal/parsing"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path"
)

func createDemoFile() {
	cwd, _ := os.Getwd()
	filePath := path.Join(cwd, "demo.txt")
	_ = ioutil.WriteFile(filePath, []byte("Hello World!"), 0644)
}

func deleteDemoFile() {
	cwd, _ := os.Getwd()
	filePath := path.Join(cwd, "demo.txt")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return
	}
	_ = os.Remove(filePath)
}

var _ = ginkgo.Describe("Pathfinder Tests", func() {
	ginkgo.Context("[Method]: GetFileInCurrentWorkingDirectory", func() {
		ginkgo.It("get the Path to a File that exists", func() {
			createDemoFile()
			finder := parsing.Pathfinder{}
			filePath, err := finder.GetFileInCurrentWorkingDirectory("demo.txt")
			gomega.Expect(err).To(gomega.BeNil())
			cwd, _ := os.Getwd()
			gomega.Expect(filePath).To(gomega.Equal(path.Join(cwd, "demo.txt")))
			deleteDemoFile()
		})
		ginkgo.It("error when file not exists", func() {
			finder := parsing.Pathfinder{}
			path, err := finder.GetFileInCurrentWorkingDirectory("demo.txt")
			gomega.Expect(err).To(gomega.Not(gomega.BeNil()))
			gomega.Expect(err.Error()).To(gomega.Equal("file does not exists"))
			gomega.Expect(path).To(gomega.Equal(""))
		})
	})
})
