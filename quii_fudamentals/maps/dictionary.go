package maps

type Dictionary map[string]string
type DictionaryError string

var (
    ErrNotFound = DictionaryError("could not find the word that you were looking for")
    ErrWordExists = DictionaryError("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryError("cannot perform operation on word because it does not exist")
)

func (e DictionaryError) Error() string {
    return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }
    return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)
    switch err {
        case ErrNotFound:
		    d[word] = definition
        case nil:
            return ErrWordExists
        default:
            return err
    }
    return nil
}

func (d Dictionary) Update(word, definition string) error {
    _, err := d.Search(word)

	switch err {
        case ErrNotFound:
            return ErrWordDoesNotExist
        case nil:
            d[word] = definition
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
