package stringset

type StringSet struct {
	set map[string]struct{}
}

func New() StringSet {
	set := make(map[string]struct{})
	return StringSet{set: set}
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
