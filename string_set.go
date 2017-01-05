// Package stringset provides a value type for a set of strings.
package stringset

import (
	"fmt"
	"strings"
)

// StringSet represents a set of strings.
type StringSet struct {
	set map[string]struct{}
}

// New returns a StringSet with zero or more members.
func New(members ...string) StringSet {
	set := make(map[string]struct{})
	for _, member := range members {
		set[member] = struct{}{}
	}
	return StringSet{set: set}
}

// String returns a string representation of s.
func (s StringSet) String() string {
	return fmt.Sprintf("{%s}", strings.Join(s.Members(), " "))
}

// IsEmpty returns true when s has no members.
func (s StringSet) IsEmpty() bool {
	return len(s.set) == 0
}

// Length returns the number of members in s.
func (s StringSet) Length() int {
	return len(s.set)
}

// Contains returns true if member exists in s.
func (s StringSet) Contains(member string) bool {
	_, ok := s.set[member]
	return ok
}

// Members returns a slice of the members of s, order is not guaranteed.
func (s StringSet) Members() []string {
	members := []string{}
	for member := range s.set {
		members = append(members, member)
	}
	return members
}

// IsSubset returns true if s is a subset of other.
func (s StringSet) IsSubset(other StringSet) bool {
	for member := range s.set {
		if !other.Contains(member) {
			return false
		}
	}
	return true
}

// IsProperSubset returns true if s is a proper subset of other.
func (s StringSet) IsProperSubset(other StringSet) bool {
	return other.Length() > s.Length() && s.IsSubset(other)
}

// IsSuperset returns true if s is a superset of other.
func (s StringSet) IsSuperset(other StringSet) bool {
	return other.IsSubset(s)
}

// Subtract returns a new StringSet of members of s not in other.
func (s StringSet) Subtract(other StringSet) StringSet {
	set := make(map[string]struct{})
	for member := range s.set {
		if !other.Contains(member) {
			set[member] = struct{}{}
		}
	}
	return StringSet{set: set}
}

// Intersection returns a new StringSet of members in both s and other.
func (s StringSet) Intersection(other StringSet) StringSet {
	set := make(map[string]struct{})
	for member := range s.set {
		if other.Contains(member) {
			set[member] = struct{}{}
		}
	}
	return StringSet{set: set}
}

// Union returns a new StringSet of all the members of s and other.
func (s StringSet) Union(other StringSet) StringSet {
	set := make(map[string]struct{})
	for member := range s.set {
		set[member] = struct{}{}
	}
	for member := range other.set {
		set[member] = struct{}{}
	}
	return StringSet{set: set}
}

// SymmetricDifference returns a new StringSet of the members of s and other
// that are not in both sets.
func (s StringSet) SymmetricDifference(other StringSet) StringSet {
	set := make(map[string]struct{})
	for member := range s.set {
		if !other.Contains(member) {
			set[member] = struct{}{}
		}
	}
	for member := range other.set {
		if !s.Contains(member) {
			set[member] = struct{}{}
		}
	}
	return StringSet{set: set}
}
