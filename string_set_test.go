package stringset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/stringset"
)

var _ = Describe("StringSet", func() {
	Context("when it contains a string", func() {
		It("returns true", func() {
			set := stringset.New([]string{"monkey", "bananas", "trees"})
			Expect(set.Contains("monkey")).To(BeTrue())
		})
	})

	Context("when it does not contain a string", func() {
		It("returns false", func() {
			set := stringset.New([]string{"bananas", "trees"})
			Expect(set.Contains("monkey")).To(BeFalse())
		})
	})
})
