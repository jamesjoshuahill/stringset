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
