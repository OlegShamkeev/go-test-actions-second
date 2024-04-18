package main

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != a {
		t.Errorf("expected %d, got %d", a, res)
	}
}

/*func TestMain(m *testing.M) {
	main()
}*/
