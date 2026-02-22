package dictionary

import "errors"

type Dictionary map[string]string

type Actions interface {
	Search (string) (string, error)
	Add (string, string) error
}

var (
	ErrorNotFound = errors.New("could not find the word you were looking for")
	ErrorWordExists = errors.New("already exists")
)

func (d Dictionary) Search (keyword string) (string, error) {
	definition, ok := d[keyword]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add (word string, definition string) error {
	_, err := d.Search(word)
	switch err {
		case ErrorNotFound:
			d[word] = definition
		case nil:
			d[word] = definition
		default:
			return err
	}

	return nil
}