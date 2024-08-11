package sf

import (
	"testing"
)

type SimpleAttributes struct {
	A string
	B NestedAttributes
}

type NestedAttributes struct {
	C int
	D NestedNestedAttributes
}

type NestedNestedAttributes struct {
	E int
}

func TestAttributes(t *testing.T) {
	s := SimpleAttributes{
		A: "z",
		B: NestedAttributes{
			C: 1,
			D: NestedNestedAttributes{
				E: 2,
			},
		},
	}

	AttributeAssign(&s, "A", "zzz")
	if s.A != "zzz" {
		t.Errorf("attribute value was not assigned correctly to 'zzz'")
	}

	AttributeAssign(&s, "B.C", 100)
	if s.B.C != 100 {
		t.Errorf("attribute value was not assigned correctly to '100'")
	}

	AttributeAssign(&s, "B.D.E", 200)
	if s.B.D.E != 200 {
		t.Errorf("attribute value was not assigned correctly to '200'")
	}

	AttributeAssign(&s.B, "D.E", 300)
	if s.B.D.E != 300 {
		t.Errorf("attribute value was not assigned correctly to '300'")
	}
}
