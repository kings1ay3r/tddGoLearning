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

	t.Run("Read Known Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		got, _ := dictionary.search("test")
		want := "defenition of test"

		testAssertion(t, got, want)
	})
	t.Run("Read UnKnown Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		_, got := dictionary.search("retest")
		want := ErrNOtFound
		testAssertion(t, got.Error(), want.Error())
	})
	t.Run("Create New Word", func(t *testing.T) {
		dictionary := Haystack{}
		dictionary.Add("test", "defenition of test")
		got, _ := dictionary.search("test")
		want := "defenition of test"

		testAssertion(t, got, want)
	})
	t.Run("Create New Word, Exist Error", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		got := dictionary.Add("test", "defeniton of test")
		want := ErrWordExist

		testAssertion(t, got.Error(), want.Error())
	})
	t.Run("Updating Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		dictionary.Update("test", "test Redefined")
		got, _ := dictionary.search("test")
		want := "test Redefined"
		testAssertion(t, got, want)
	})
	t.Run("Updating Word : Exist Error", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		got := dictionary.Update("retest", "meaning of retest")
		want := ErrNOtFound
		testAssertion(t, got.Error(), want.Error())
	})
	t.Run("Delete Word", func(t *testing.T) {
		dictionary := Haystack{"test": "defenition of test"}
		dictionary.Delete("test")
		_, got := dictionary.search("test")
		want := ErrNOtFound
		testAssertion(t, got.Error(), want.Error())
	})
	t.Run("Delete Word : Unknowm Error", func(t *testing.T) {
		dictionary := Haystack{}
		got := dictionary.Delete("test")
		want := ErrNOtFound
		testAssertion(t, got.Error(), want.Error())
	})
}
