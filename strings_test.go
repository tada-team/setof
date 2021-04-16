package setof

import "testing"

func TestSet(t *testing.T) {
	s := NewStrings().Update("a").Update("a")
	if len(s.SortedList()) != 1 || s.SortedList()[0] != "a" {
		t.Error("StringsSet fail")
	}
}
