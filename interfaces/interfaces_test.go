package interfaces

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {

	assertTest := func(t testing.TB, got, want float64) {
		if got != want {
			var gotString string = fmt.Sprint(got)
			var wantString string = fmt.Sprint(want)
			errormsg := "Expecting " + wantString + " and recieved " + gotString
			t.Errorf(errormsg)
		}
	}
	/* 	t.Run("Testing Perimeter", func(t *testing.T) {
		obj := Rectangle{10.0, 10.0}
		got := Perimeter(obj)
		want := 40.0
		assertTest(t, got, want)
	}) */
	t.Run("Testing Interface anonymous struct", func(t *testing.T) {
		perimeterTest := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 36},
			{Circle{6}, 36},
			{Triangle{[3]float64{3, 3, 3}}, 9},
		}
		for _, testObj := range perimeterTest {

			obj := testObj.shape
			got := obj.Perimeter()
			want := testObj.want
			assertTest(t, got, want)
		}
	})
	t.Run("Testing Perimeter", func(t *testing.T) {
		obj := Rectangle{13.30, 10.0}
		got := obj.Perimeter()
		want := 46.6
		assertTest(t, got, want)
	})
	t.Run("Testing Area Rectangle", func(t *testing.T) {
		obj := Rectangle{13.30, 10.0}
		got := obj.Area()
		want := 133.0
		assertTest(t, got, want)
	})
	t.Run("Testing Area Circle", func(t *testing.T) {
		obj := Circle{10.0}
		got := obj.Area()
		want := 300.0
		assertTest(t, got, want)
	})
	t.Run("Testing Area Triangle", func(t *testing.T) {
		obj := Triangle{[3]float64{3, 3, 3}}
		got := obj.Area()
		want := 27.0
		assertTest(t, got, want)
	})

}
