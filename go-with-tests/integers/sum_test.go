package integers

import "testing"

func TestSum(t *testing.T) {
 sum := Sum(1, 2)
 want := 3

 if sum != want {
	t.Errorf("expected '%d' but got '%d'", want, sum)
 }
}