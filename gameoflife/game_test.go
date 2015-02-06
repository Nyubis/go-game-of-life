package gameoflife

import "testing"

type point struct {
	x, y int
	val bool
}

func Test(t *testing.T) {
	//Simple tests
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
			t.Errorf("GetAt(%d, %d) == %t, want %t", w.x, w.y, w.val)
		}
	}
	//Beacon test
	b = MakeBoard(4, 4)
	b.SetAt(0,0,true)
	b.SetAt(0,1,true)
	b.SetAt(1,0,true)
	b.SetAt(2,3,true)
	b.SetAt(3,2,true)
	b.SetAt(3,3,true)
	for i := 0; i<5; i++ {
		first := b.GetAt(1,1)
		second := b.GetAt(2,2)
		if (first != (i%2==1) || second != (i%2==1)) {
			t.Errorf("Beacon iteration %d: cells are %t and %t", i, first, second)
		}
		b.Step()
	}
}
