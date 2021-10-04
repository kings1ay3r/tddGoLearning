package main

import (
	"bytes"
	"fmt"
	"testing"
)

type Spysleeper struct {
	Calls int
}

func (s *Spysleeper) Sleep() {
	s.Calls++
	println("sleep")
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &Spysleeper{}

	CountDown(buffer, spy)

	got := buffer.String()
	want := "321Go!"
	fmt.Sprintf(got)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
