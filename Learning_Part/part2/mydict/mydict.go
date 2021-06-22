package mydict

import "errors"

type Dict map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Already Exists")
	errCantUpdate = errors.New("Can't Update non-existing word")
)

func (d Dict) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dict) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err != nil {
		return errWordExists
	}
	return nil
}

func (d Dict) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dict) Delete(word string) {
	delete(d, word)
}
