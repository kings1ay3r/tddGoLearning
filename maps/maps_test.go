package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "defenition of test"}
	got := DicitonarySearch(dictionary, "test")
	want := "defenition of test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
