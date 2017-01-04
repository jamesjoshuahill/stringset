package stringset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"

	"testing"
)

func TestStringset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringset Suite")
}

func sliceOfStrings(length int) []string {
	var slice []string
	for i := 0; i < length; i++ {
		slice = append(slice, uuid.New())
	}
	return slice
}

func sliceOfStringsIncludingMonkeys(length, monkeysIndex int) []string {
	var slice []string
	for i := 0; i < length; i++ {
		if i == monkeysIndex {
			slice = append(slice, "monkeys")
		} else {
			slice = append(slice, uuid.New())
		}
	}
	return slice
}
