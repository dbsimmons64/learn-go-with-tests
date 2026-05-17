package dictionary

import (
	"errors"
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound         = errors.New("could not find the word you are looking for")
	ErrWordExists       = errors.New("word exists in dictionary")
	ErrWordDoesNotExist = errors.New("word does not exist in dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = meaning
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
