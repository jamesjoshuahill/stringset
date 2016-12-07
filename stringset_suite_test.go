package stringset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStringset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringset Suite")
}
