package stringset

type StringSet struct {
	set map[string]struct{}
}

func New(slice []string) StringSet {
	set := make(map[string]struct{})
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return StringSet{set: set}
}

func (s StringSet) Contains(item string) bool {
	_, ok := s.set[item]
	return ok
}
