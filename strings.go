package setof

import "sort"

type strings map[string]struct{}

func Strings(v ...string) strings {
	set := make(strings)
	if len(v) > 0 {
		set.Update(v...)
	}
	return set
}

func (set strings) Update(v ...string) strings {
	for _, s := range v {
		set[s] = struct{}{}
	}
	return set
}

func (set strings) Remove(v ...string) strings {
	for _, s := range v {
		delete(set, s)
	}
	return set
}

func (set strings) Contains(v string) bool {
	_, ok := set[v]
	return ok
}

func (set strings) SortedList() []string {
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
