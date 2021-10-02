package maps

import "errors"

type Haystack map[string]string

var ErrNOtFound = errors.New("invalid word")

func (h Haystack) search(needle string) (string, error) {
	defenition, ok := h[needle]
	if !ok {
		return defenition, ErrNOtFound
	}
	return defenition, nil
}
