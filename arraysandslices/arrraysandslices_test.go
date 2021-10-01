package arraysandslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	got := Sum(numbers)
	want := 21

	if got != want {
		var gotString string = fmt.Sprint(got)
		var wantString string = fmt.Sprint(want)
		errormsg := "Expecting " + wantString + " and recieved " + gotString
		t.Errorf(errormsg)
	}
}
