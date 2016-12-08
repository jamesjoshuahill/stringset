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

	Context("when the other set has a member in common", func() {
		It("subtracts the member in common", func() {
			set := stringset.New([]string{"monkeys", "bananas"})
			other := stringset.New([]string{"bananas"})
			Expect(set.Subtract(other)).To(Equal(stringset.New([]string{"monkeys"})))
		})
	})

	Context("when the other set has nothing in common", func() {
		It("does not subtract any members", func() {
			set := stringset.New([]string{"monkeys", "bananas"})
			other := stringset.New([]string{"trees", "sunshine"})
			Expect(set.Subtract(other)).To(Equal(stringset.New([]string{"monkeys", "bananas"})))
		})
	})
})
