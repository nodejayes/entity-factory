package parsing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParsingConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParsingConfig Suite")
}
