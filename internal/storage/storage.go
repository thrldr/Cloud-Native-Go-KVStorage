package storage

var Storage = make(map[string]string)

func Put(key, value string) error {
	Storage[key] = value

	return nil
}

func Get(key string) (string, error) {
	value, ok := Storage[key]

	if !ok {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(Storage, key)

	return nil
}
