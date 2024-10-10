package tests

import (
	"1hw/storage"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	addTests := []Test{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
		{"key4", "value4"},
		{"key5", "value5"},
		{"key6", "value6"},
		{"key7", "value7"},
		{"key8", "value8"},
		{"key9", "value9"},
	}
	storage := storage.NewStorage()
	for _, test := range addTests {
		storage.Set(test.key, test.value)
	}
	for _, test := range addTests {
		got := storage.Get(test.key)
		if test.value != *got {
			t.Errorf("got %s ; wanted %s", *got, test.value)
		}
	}

}

func TestGetKind(t *testing.T) {
	addTests := []Test{
		{"key1", "1"},
		{"key2", "value2"},
		{"key3", "3"},
		{"key4", "value4"},
		{"key5", "5"},
		{"key6", "value6"},
		{"key7", "7"},
		{"key8", "value8"},
		{"key9", "9"},
	}
	storage := storage.NewStorage()
	for i, test := range addTests {
		storage.Set(test.key, test.value)
		got := storage.GetKind(test.key)

		if i%2 == 0 && got != "D" {
			t.Errorf("got %s ; wanted %s", got, "D")
		} else if i%2 != 0 && got != "S" {
			t.Errorf("got %s ; wanted %s", got, "S")
		}

	}
}
