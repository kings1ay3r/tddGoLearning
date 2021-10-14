package select_

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)
	if bDuration > aDuration {
		return a
	}
	return b
}
