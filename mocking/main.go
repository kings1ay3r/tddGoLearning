package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type defaultsleeper struct {
	Calls int
	Sleep func()
}

type configurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	fmt.Fprint(out, "3")
	sleeper.Sleep()
	fmt.Fprint(out, "2")
	sleeper.Sleep()
	fmt.Fprint(out, "1")
	sleeper.Sleep()
	fmt.Fprint(out, "Go!")
}

// TODO check coverage of this test and understnad how we can cover defaultsleeper
// TODO Understand Mocking Better
func main() {
	CountDown(os.Stdout, &defaultsleeper{})
}
