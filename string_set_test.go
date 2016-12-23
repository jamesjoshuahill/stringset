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

		It("is empty", func() {
			emptySet := stringset.New()
			Expect(emptySet.Empty()).To(BeTrue())
		})

		It("has an order of 0", func() {
			emptySet := stringset.New()
			Expect(emptySet.Order()).To(Equal(0))
		})

		It("is a subset of itself", func() {
			emptySet := stringset.New()
			Expect(emptySet.IsSubset(emptySet)).To(BeTrue())
		})
	})

	Context("when it has one member", func() {
		It("contains the member", func() {
			set := stringset.New("monkeys")
			Expect(set.Contains("monkeys")).To(BeTrue())
		})

		It("does not contain another member", func() {
			set := stringset.New("bananas")
			Expect(set.Contains("monkeys")).To(BeFalse())
		})

		It("prints neatly", func() {
			set := stringset.New("monkeys")
			Expect(fmt.Sprintf("%v", set)).To(Equal(`{monkeys}`))
		})

		It("is not empty", func() {
			set := stringset.New("monkeys")
			Expect(set.Empty()).To(BeFalse())
		})

		It("has an order of one", func() {
			set := stringset.New("monkeys")
			Expect(set.Order()).To(Equal(1))
		})

		It("is a subset of itself", func() {
			set := stringset.New("monkeys")
			Expect(set.IsSubset(set)).To(BeTrue())
		})
	})

	Context("when it has some members", func() {
		It("lists all members", func() {
			set := stringset.New("monkeys", "bananas", "trees")
			Expect(set.Members()).To(ContainElement("monkeys"))
			Expect(set.Members()).To(ContainElement("bananas"))
			Expect(set.Members()).To(ContainElement("trees"))
		})

		It("prints neatly", func() {
			set := stringset.New("monkeys", "bananas", "trees")
			Expect(fmt.Sprintf("%v", set)).To(Or(
				Equal(`{bananas monkeys trees}`),
				Equal(`{bananas trees monkeys}`),
				Equal(`{monkeys bananas trees}`),
				Equal(`{monkeys trees bananas}`),
				Equal(`{trees bananas monkeys}`),
				Equal(`{trees monkeys bananas}`),
			))
		})

		It("is not empty", func() {
			set := stringset.New("monkeys", "bananas", "trees")
			Expect(set.Empty()).To(BeFalse())
		})

		It("has an order of three", func() {
			set := stringset.New("monkeys", "bananas", "trees")
			Expect(set.Order()).To(Equal(3))
		})

		It("is a subset of itself", func() {
			set := stringset.New("monkeys", "bananas", "trees")
			Expect(set.IsSubset(set)).To(BeTrue())
		})
	})

	Context("when empty and the other set has a member", func() {
		It("subtracts nothing", func() {
			emptySet := stringset.New()
			other := stringset.New("trees")
			Expect(emptySet.Subtract(other)).To(Equal(emptySet))
		})

		It("has no members in common", func() {
			emptySet := stringset.New()
			other := stringset.New("trees")
			Expect(emptySet.Intersection(other)).To(Equal(emptySet))
		})

		It("includes all members in the union", func() {
			emptySet := stringset.New()
			other := stringset.New("trees")
			Expect(emptySet.Union(other)).To(Equal(other))

			By("not changing the set")
			Expect(emptySet).To(Equal(stringset.New()))
		})

		It("includes the member in the symmetric difference", func() {
			emptySet := stringset.New()
			other := stringset.New("trees")
			Expect(emptySet.SymmetricDifference(other)).To(Equal(other))

			By("not changing the set")
			Expect(emptySet).To(Equal(stringset.New()))
		})

		It("is a subset of the other", func() {
			emptySet := stringset.New()
			other := stringset.New("trees")
			Expect(emptySet.IsSubset(other)).To(BeTrue())
		})
	})

	Context("when it has members and the other set is a subset", func() {
		It("subtracts the member in common", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas")
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys")))
		})

		It("intersects", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas")
			Expect(set.Intersection(other)).To(Equal(stringset.New("bananas")))
		})

		It("includes all members in the union", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas")
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("subtracts the member in common from the symmetric difference", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas")
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas")
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})

	Context("when it has members and the other set intersects", func() {
		It("subtracts the member in common", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas", "trees")
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys")))
		})

		It("intersects", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas", "trees")
			Expect(set.Intersection(other)).To(Equal(stringset.New("bananas")))
		})

		It("includes all members in the union", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas", "trees")
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("subtracts the member in common from the symmetric difference", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas", "trees")
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("bananas", "trees")
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})

	Context("when it has members and the other set does not intersect", func() {
		It("does not subtract any members", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("trees", "sunshine")
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("does not intersect", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("trees")
			Expect(set.Intersection(other)).To(Equal(stringset.New()))
		})

		It("includes all members in the union", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("trees")
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("includes all members in the symmetric difference", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("trees")
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys", "bananas", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			set := stringset.New("monkeys", "bananas")
			other := stringset.New("trees")
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})
})
