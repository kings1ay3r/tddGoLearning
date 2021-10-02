package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	testAssertion := func(t testing.TB, got, want string) {
		t.Helper() //! TODO tounderstand

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testing Known Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		got, _ := dictionary.search("test")
		want := "defenition of test"

		testAssertion(t, got, want)
	})

	t.Run("Testing Known Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		_, got := dictionary.search("retest")
		want := ErrNOtFound
		testAssertion(t, got.Error(), want.Error())
	})
}
