package arraysandslices

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {
	assertTest := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			var gotString string = fmt.Sprint(got)
			var wantString string = fmt.Sprint(want)
			errormsg := "Expecting " + wantString + " and recieved " + gotString
			t.Errorf(errormsg)
		}
	}
	got := sumAll([]int{1, 1, 1}, []int{3, 7, 8, 9, 4})
	want := []int{3, 31}
	assertTest(t, got, want)
}
func TestSumAllTails(t *testing.T) {
	assertTest := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			var gotString string = fmt.Sprint(got)
			var wantString string = fmt.Sprint(want)
			errormsg := "Expecting " + wantString + " and recieved " + gotString
			t.Errorf(errormsg)
		}
	}
	t.Run("Normal Tails", func(t *testing.T) {

		got := sumAllTails([]int{1, 1, 1}, []int{3, 7, 8, 9, 4})
		want := []int{2, 28}
		assertTest(t, got, want)
	})
	t.Run("Empty Tails", func(t *testing.T) {

		got := sumAllTails([]int{1, 1, 1}, []int{3, 7, 8, 9, 4}, []int{})
		want := []int{2, 28, 0}
		assertTest(t, got, want)
	})
}
