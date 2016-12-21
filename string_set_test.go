package stringset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/stringset"
)

var _ = Describe("StringSet", func() {
	Context("when empty", func() {
		It("does not contain a member", func() {
			emptySet := stringset.New()
			Expect(emptySet.Contains("monkeys")).To(BeFalse())
		})

		It("lists no members", func() {
			emptySet := stringset.New()
			Expect(emptySet.Members()).To(BeEmpty())
		})

		Context("when the other set has a member", func() {
			It("subtracts nothing", func() {
				emptySet := stringset.New()
				other := stringset.New().Add("trees")
				Expect(emptySet.Subtract(other)).To(Equal(emptySet))
			})

			It("has no members in common", func() {
				emptySet := stringset.New()
				other := stringset.New().Add("trees")
				Expect(emptySet.Intersection(other)).To(Equal(emptySet))
			})
		})
	})

	Context("when it has one member", func() {
		It("contains the member", func() {
			set := stringset.New().Add("monkeys")
			Expect(set.Contains("monkeys")).To(BeTrue())
		})

		It("does not contain another member", func() {
			set := stringset.New().Add("bananas")
			Expect(set.Contains("monkeys")).To(BeFalse())
		})
	})

	Context("when it has some members", func() {
		It("lists all members alphabetically", func() {
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

		It("intersects with the member", func() {
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

		It("does not intersect", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("trees")
			Expect(set.Intersection(other)).To(Equal(stringset.New()))
		})
	})
})
