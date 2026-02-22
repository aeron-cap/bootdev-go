package shapes

import "testing"

func TestShapeArea(t *testing.T) {
 for _, test := range areaTests {
	 t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %g want %g", got, test.want)
			}
	 })
 }
}

func BenchmarkShapeArea(b *testing.B) {
	for b.Loop() {
		for _, test := range areaTests {
			test.shape.Area()
		}
	}
}