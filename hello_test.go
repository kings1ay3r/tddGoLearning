package main

import "testing"

func TestHello(t *testing.T) {

	testAssertion := func(t testing.TB, got, want string) {
		t.Helper() //! TODO tounderstand

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello to somone in Spanish", func(t *testing.T) {
		got := Hello("Eragon", "es")
		want := "Hola, Eragon"
		testAssertion(t, got, want)

	})

	t.Run("Saying Hello to somone in French", func(t *testing.T) {
		got := Hello("Arya", "fr")
		want := "Bonjour, Arya"
		testAssertion(t, got, want)

	})
	t.Run("Saying Hello World When empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		testAssertion(t, got, want)

	})
	t.Run("Saying Hello World When empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		testAssertion(t, got, want)

	})
}
