package gameoflife

import "testing"

type point struct {
	x, y int
	val bool
}

func Test(t *testing.T) {
	b := MakeBoard(10, 10)
	var tests = []struct {
		give, want point
	}{
		{point{2,3,true}, point{2,3,true}},      //Basic point check
		{point{1,1,true}, point{1,2,false}},     //Initialise check
		{point{1,2,true}, point{1,1,true}},      //Older point check
		{point{10,10,true}, point{10,10,false}}, //Out of bounds check
	}
	for _, test := range tests {
		g := test.give
		w := test.want
		b.SetAt(g.x, g.y, g.val)
		got := b.GetAt(w.x, w.y)
		if got != w.val {
			t.Errorf("GetAt(%d, %d) == %b, want %b", w.x, w.y, w.val)
		}
	}
}
