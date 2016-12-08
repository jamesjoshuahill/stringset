package stringset

type StringSet struct {
	set map[string]struct{}
}

func New(slice []string) StringSet {
	set := make(map[string]struct{})
	for _, element := range slice {
		set[element] = struct{}{}
	}
	return StringSet{set: set}
}

func (s StringSet) Contains(member string) bool {
	_, ok := s.set[member]
	return ok
}

func (s StringSet) Subtract(other StringSet) StringSet {
	difference := New([]string{})
	for member, _ := range s.set {
		if !other.Contains(member) {
			difference.set[member] = struct{}{}
		}
	}
	return difference
}
