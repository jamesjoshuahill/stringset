package stringset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/stringset"
)

var _ = Describe("StringSet", func() {
	Context("when it has one member", func() {
		It("knows it contains the member", func() {
			set := stringset.New().Add("monkeys")
			Expect(set.Contains("monkeys")).To(BeTrue())
		})

		It("knows it does not contain another member", func() {
			set := stringset.New().Add("bananas")
			Expect(set.Contains("monkeys")).To(BeFalse())
		})
	})

	Context("when it has some members", func() {
		It("produces a sorted list of members", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees", "sunshine"})
			sortedList := []string{"bananas", "monkeys", "sunshine", "trees"}
			Expect(set.Members()).To(Equal(sortedList))
		})
	})

	Context("when the other set has a member in common", func() {
		It("subtracts the member in common", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("bananas")
			Expect(set.Subtract(other)).To(Equal(stringset.New().Add("monkeys")))
		})

		It("knows which member is in common", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("bananas")
			Expect(set.Intersection(other)).To(Equal(stringset.New().Add("bananas")))
		})
	})

	Context("when the other set has nothing in common", func() {
		It("does not subtract any members", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().AddSlice([]string{"trees", "sunshine"})
			Expect(set.Subtract(other)).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))
		})

		It("knows no members are in common", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("trees")
			Expect(set.Intersection(other)).To(Equal(stringset.New()))
		})
	})
})
