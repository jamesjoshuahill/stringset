package stringset_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/stringset"
)

var _ = Describe("StringSet", func() {
	var set stringset.StringSet
	var other stringset.StringSet

	Context("when is has no members", func() {
		BeforeEach(func() {
			set = stringset.New()
		})

		It("does not contain a member", func() {
			Expect(set.Contains("monkeys")).To(BeFalse())
		})

		It("lists no members", func() {
			Expect(set.Members()).To(BeEmpty())
		})

		It("prints neatly", func() {
			Expect(fmt.Sprintf("%v", set)).To(Equal(`{}`))
		})

		It("is empty", func() {
			Expect(set.Empty()).To(BeTrue())
		})

		It("has an order of 0", func() {
			Expect(set.Order()).To(Equal(0))
		})

		It("is a subset of itself", func() {
			Expect(set.IsSubset(set)).To(BeTrue())
		})
	})

	Context("when it has one member", func() {
		BeforeEach(func() {
			set = stringset.New("monkeys")
		})

		It("contains the member", func() {
			Expect(set.Contains("monkeys")).To(BeTrue())
		})

		It("does not contain another member", func() {
			Expect(set.Contains("bananas")).To(BeFalse())
		})

		It("prints neatly", func() {
			Expect(fmt.Sprintf("%v", set)).To(Equal(`{monkeys}`))
		})

		It("is not empty", func() {
			Expect(set.Empty()).To(BeFalse())
		})

		It("has an order of one", func() {
			Expect(set.Order()).To(Equal(1))
		})

		It("is a subset of itself", func() {
			Expect(set.IsSubset(set)).To(BeTrue())
		})
	})

	Context("when it has some members", func() {
		BeforeEach(func() {
			set = stringset.New("monkeys", "bananas", "trees")
		})

		It("lists all members", func() {
			Expect(set.Members()).To(ContainElement("monkeys"))
			Expect(set.Members()).To(ContainElement("bananas"))
			Expect(set.Members()).To(ContainElement("trees"))
		})

		It("prints neatly", func() {
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
			Expect(set.Empty()).To(BeFalse())
		})

		It("has an order of three", func() {
			Expect(set.Order()).To(Equal(3))
		})

		It("is a subset of itself", func() {
			Expect(set.IsSubset(set)).To(BeTrue())
		})
	})

	Context("when empty and the other set has a member", func() {
		BeforeEach(func() {
			set = stringset.New()
			other = stringset.New("trees")
		})

		It("subtracts nothing", func() {
			Expect(set.Subtract(other)).To(Equal(stringset.New()))
		})

		It("has no members in common", func() {
			Expect(set.Intersection(other)).To(Equal(stringset.New()))
		})

		It("includes all members in the union", func() {
			Expect(set.Union(other)).To(Equal(other))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New()))
		})

		It("includes the member in the symmetric difference", func() {
			Expect(set.SymmetricDifference(other)).To(Equal(other))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New()))
		})

		It("is a subset of the other", func() {
			Expect(set.IsSubset(other)).To(BeTrue())
		})
	})

	Context("when it has members and the other set is a subset", func() {
		BeforeEach(func() {
			set = stringset.New("monkeys", "bananas")
			other = stringset.New("bananas")
		})

		It("subtracts the member in common", func() {
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys")))
		})

		It("intersects", func() {
			Expect(set.Intersection(other)).To(Equal(stringset.New("bananas")))
		})

		It("includes all members in the union", func() {
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("subtracts the member in common from the symmetric difference", func() {
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})

	Context("when it has members and the other set intersects", func() {
		BeforeEach(func() {
			set = stringset.New("monkeys", "bananas")
			other = stringset.New("bananas", "trees")
		})

		It("subtracts the member in common", func() {
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys")))
		})

		It("intersects", func() {
			Expect(set.Intersection(other)).To(Equal(stringset.New("bananas")))
		})

		It("includes all members in the union", func() {
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("subtracts the member in common from the symmetric difference", func() {
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys", "trees")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})

	Context("when it has members and the other set does not intersect", func() {
		BeforeEach(func() {
			set = stringset.New("monkeys", "bananas")
			other = stringset.New("trees", "sunshine")
		})

		It("does not subtract any members", func() {
			Expect(set.Subtract(other)).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("does not intersect", func() {
			Expect(set.Intersection(other)).To(Equal(stringset.New()))
		})

		It("includes all members in the union", func() {
			Expect(set.Union(other)).To(Equal(stringset.New("monkeys", "bananas", "trees", "sunshine")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("includes all members in the symmetric difference", func() {
			Expect(set.SymmetricDifference(other)).To(Equal(stringset.New("monkeys", "bananas", "trees", "sunshine")))

			By("not changing the set")
			Expect(set).To(Equal(stringset.New("monkeys", "bananas")))
		})

		It("is not a subset of the other", func() {
			Expect(set.IsSubset(other)).To(BeFalse())
		})
	})
})
