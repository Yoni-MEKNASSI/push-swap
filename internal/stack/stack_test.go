package stack

import (
	"reflect"
	"testing"
)

func TestExampleFromSubject(t *testing.T) {
	s := &Stacks{A: []int{2, 1, 3, 6, 8, 5}}
	for _, op := range []string{"sa", "pb", "pb", "pb", "rb", "rrr"} {
		s.Apply(op)
	}
	if !reflect.DeepEqual(s.A, []int{5, 6, 8}) || !reflect.DeepEqual(s.B, []int{3, 2, 1}) {
		t.Errorf("a=%v b=%v", s.A, s.B)
	}
	s.Apply("pa")
	s.Apply("pa")
	s.Apply("pa")
	if !s.IsSorted() {
		t.Errorf("a devrait être triée : a=%v b=%v", s.A, s.B)
	}
}

func TestEmptyStacksAndUnknown(t *testing.T) {
	s := &Stacks{}
	for _, op := range []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"} {
		if !s.Apply(op) {
			t.Errorf("%s devrait être valide", op)
		}
	}
	if s.Apply("xx") {
		t.Error("xx ne devrait pas être valide")
	}
}

func TestParse(t *testing.T) {
	if nums, ok := Parse([]string{"2 1 3", "-4"}); !ok || !reflect.DeepEqual(nums, []int{2, 1, 3, -4}) {
		t.Errorf("nums=%v ok=%v", nums, ok)
	}
	for _, bad := range []string{"0 one 2", "1 2 2", "1.5 2", "99999999999999999999"} {
		if _, ok := Parse([]string{bad}); ok {
			t.Errorf("%q devrait échouer", bad)
		}
	}
}
