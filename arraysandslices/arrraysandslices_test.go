package arraysandslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	assertTest := func(t testing.TB, got, want int) {
		if got != want {
			var gotString string = fmt.Sprint(got)
			var wantString string = fmt.Sprint(want)
			errormsg := "Expecting " + wantString + " and recieved " + gotString
			t.Errorf(errormsg)
		}
	}
	t.Run("Collection of 6", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
		assertTest(t, got, want)
	})
	t.Run("Collection of 5", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertTest(t, got, want)
	})
}
