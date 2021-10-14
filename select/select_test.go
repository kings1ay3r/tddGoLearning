package select_

import "testing"

func TestRacer(t *testing.T) {
	slowurl := "http://facebook.com"
	fasturl := "https://www.google.com"

	want := fasturl
	got := Racer(slowurl, fasturl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
