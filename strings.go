package setof

import "sort"

type Strings map[string]struct{}

func NewStrings(v ...string) Strings {
	set := make(Strings)
	if len(v) > 0 {
		set.Update(v...)
	}
	return set
}

func (set Strings) Update(v ...string) Strings {
	for _, s := range v {
		set[s] = struct{}{}
	}
	return set
}

func (set Strings) Remove(v ...string) Strings {
	for _, s := range v {
		delete(set, s)
	}
	return set
}

func (set Strings) Contains(v string) bool {
	_, ok := set[v]
	return ok
}

func (set Strings) SortedList() []string {
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
