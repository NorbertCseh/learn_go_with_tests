package notmaps

const (
	ErrNotFound         = DictionaryErr("Could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("Cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("Cannot find word")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, text string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = text
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, text string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = text
	default:
		return err
	}
	return nil

}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
