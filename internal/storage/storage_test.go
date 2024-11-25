package storage

import (
	"errors"
	"testing"
)

func TestStorage_Get(t *testing.T) {
	const key = "key"
	const valueWant = "value"

	Storage[key] = valueWant
	defer delete(Storage, key)

	valueGot, err := Get(key)

	if err != nil {
		t.Errorf("Expected to work without errors, but got: '%v'", err)
	}

	if valueGot != valueWant {
		t.Fatalf("Expected value: '%s'; got: '%s'", valueWant, valueGot)
	}
}

func TestStorage_GetNoItem(t *testing.T) {
	const key = "key"

	_, errGot := Get(key)

	if !errors.Is(errGot, ErrNoSuchKey) {
		t.Fatalf("Expected to get no such key error")
	}
}

func TestStorage_Delete(t *testing.T) {
	const key = "key"
	const value = "value"

	Storage[key] = value
	defer delete(Storage, key)

	errDeleteOnce := Delete(key)
	errDeleteTwice := Delete(key)

	if errDeleteOnce != nil || errDeleteTwice != nil {
		t.Error("deletion should return no error")
	}

	if len(Storage) > 0 {
		t.Error("deletion should clean up the underlying map")
	}
}

func TestStorage_Put(t *testing.T) {
	const key = "key"
	const valueWant = "value"

	errPut := Put(key, valueWant)
	errPutTwice := Put(key, valueWant)

	if errPut != nil || errPutTwice != nil {
		t.Error("put operation should return no error")
	}

	if len(Storage) != 1 {
		t.Error("put operation should put the item into map in idempotent manner")
	}

	if val, ok := Storage[key]; !ok || val != valueWant {
		t.Error("put operation should put a correct item to the map")
	}
}
