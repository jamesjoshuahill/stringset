package stringset_test

import (
	"fmt"

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

		It("prints neatly", func() {
			emptySet := stringset.New()
			Expect(fmt.Sprintf("%v", emptySet)).To(Equal(`{}`))
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

			It("includes all members in the union", func() {
				emptySet := stringset.New()
				other := stringset.New().Add("trees")
				Expect(emptySet.Union(other)).To(Equal(other))

				By("not changing the set")
				Expect(emptySet).To(Equal(stringset.New()))
			})

			It("includes the member in the symmetric difference", func() {
				emptySet := stringset.New()
				other := stringset.New().Add("trees")
				Expect(emptySet.SymmetricDifference(other)).To(Equal(other))

				By("not changing the set")
				Expect(emptySet).To(Equal(stringset.New()))
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

		It("prints neatly", func() {
			set := stringset.New().Add("monkeys")
			Expect(fmt.Sprintf("%v", set)).To(Equal(`{monkeys}`))
		})
	})

	Context("when it has some members", func() {
		It("lists all members alphabetically", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees", "sunshine"})
			sortedList := []string{"bananas", "monkeys", "sunshine", "trees"}
			Expect(set.Members()).To(Equal(sortedList))
		})

		It("prints neatly", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees", "sunshine"})
			Expect(fmt.Sprintf("%v", set)).To(Equal(`{bananas monkeys sunshine trees}`))
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

		It("includes all members in the union", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("bananas")
			Expect(set.Union(other)).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))
		})

		It("subtracts the member in common from the symmetric difference", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("bananas")
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New().AddSlice([]string{"monkeys"})))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))
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

		It("includes all members in the union", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("trees")
			Expect(set.Union(other)).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))
		})

		It("includes all members in the symmetric difference", func() {
			set := stringset.New().AddSlice([]string{"monkeys", "bananas"})
			other := stringset.New().Add("trees")
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New().AddSlice([]string{"monkeys", "bananas"})))
		})
	})
})
