package maps

import "errors"

type Haystack map[string]string

var ErrNOtFound = errors.New("invalid word")
var ErrWordExist = errors.New("Word Exists")

func (h Haystack) search(needle string) (string, error) {
	defenition, ok := h[needle]
	if !ok {
		return defenition, ErrNOtFound
	}
	return defenition, nil
}

func (h Haystack) Add(key, value string) error {
	_, ok := h[key]
	if ok {
		return ErrWordExist
	}
	h[key] = value
	return nil
}

func (h Haystack) Update(key, value string) error {
	_, ok := h[key]
	if ok {
		h[key] = value
		return nil
	} else {
		return ErrNOtFound

	}
}

func (h Haystack) Delete(key string) error {
	_, ok := h[key]
	if ok {
		delete(h, key)
		return nil
	} else {
		return ErrNOtFound
	}
}
