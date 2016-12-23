package stringset

import (
	"fmt"
	"strings"
)

type StringSet struct {
	set map[string]struct{}
}

func New(members ...string) StringSet {
	set := make(map[string]struct{})
	for _, member := range members {
		set[member] = struct{}{}
	}
	return StringSet{set: set}
}

func (s StringSet) String() string {
	return fmt.Sprintf("{%s}", strings.Join(s.Members(), " "))
}

func (s StringSet) Empty() bool {
	return len(s.set) == 0
}

func (s StringSet) Order() int {
	return len(s.set)
}

func (s StringSet) Contains(member string) bool {
	_, ok := s.set[member]
	return ok
}

func (s StringSet) Members() []string {
	members := []string{}
	for member, _ := range s.set {
		members = append(members, member)
	}
	return members
}

func (s StringSet) IsSubset(other StringSet) bool {
	for member, _ := range s.set {
		if !other.Contains(member) {
			return false
		}
	}
	return true
}

func (s StringSet) Subtract(other StringSet) StringSet {
	var difference []string
	for member, _ := range s.set {
		if !other.Contains(member) {
			difference = append(difference, member)
		}
	}
	return New(difference...)
}

func (s StringSet) Intersection(other StringSet) StringSet {
	var intersection []string
	for member, _ := range s.set {
		if other.Contains(member) {
			intersection = append(intersection, member)
		}
	}
	return New(intersection...)
}

func (s StringSet) Union(other StringSet) StringSet {
	var union []string
	union = append(union, s.Members()...)
	union = append(union, other.Members()...)
	return New(union...)
}

func (s StringSet) SymmetricDifference(other StringSet) StringSet {
	union := s.Union(other)
	intersection := s.Intersection(other)
	return union.Subtract(intersection)
}
