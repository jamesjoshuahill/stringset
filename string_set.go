package stringset

import (
	"fmt"
	"strings"
)

type StringSet struct {
	set map[string]struct{}
}

func New() StringSet {
	set := make(map[string]struct{})
	return StringSet{set: set}
}

func (s StringSet) String() string {
	return fmt.Sprintf("{%s}", strings.Join(s.Members(), " "))
}

func (s StringSet) Add(member string) StringSet {
	s.set[member] = struct{}{}
	return s
}

func (s StringSet) AddSlice(slice []string) StringSet {
	for _, element := range slice {
		s.Add(element)
	}
	return s
}

func (s StringSet) AddSet(other StringSet) StringSet {
	for member, _ := range other.set {
		s.Add(member)
	}
	return s
}

func (s StringSet) Members() []string {
	members := []string{}
	for member, _ := range s.set {
		members = append(members, member)
	}
	return members
}

func (s StringSet) Contains(member string) bool {
	_, ok := s.set[member]
	return ok
}

func (s StringSet) Subtract(other StringSet) StringSet {
	difference := New()
	for member, _ := range s.set {
		if !other.Contains(member) {
			difference.Add(member)
		}
	}
	return difference
}

func (s StringSet) Intersection(other StringSet) StringSet {
	intersection := New()
	for member, _ := range s.set {
		if other.Contains(member) {
			intersection.Add(member)
		}
	}
	return intersection
}

func (s StringSet) Union(other StringSet) StringSet {
	union := New()
	union.AddSet(s)
	union.AddSet(other)
	return union
}

func (s StringSet) SymmetricDifference(other StringSet) StringSet {
	union := s.Union(other)
	intersection := s.Intersection(other)
	return union.Subtract(intersection)
}
