package tests

import (
	"1hw/storage"

	"strconv"
	"testing"
)

// go test -bench=BenchmarkSetFunc
func BenchmarkSetFunc(b *testing.B) {
	addTests := []Test{}
	for i := 0; i < 10; i++ {
		addTests = append(addTests, Test{strconv.Itoa(i), strconv.Itoa(i)})
	}
	storage := storage.NewStorage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range addTests {
			storage.Set(test.key, test.value)
		}
	}
}

// go test -bench=BenchmarkGetFunc
func BenchmarkGetFunc(b *testing.B) {
	storage := storage.NewStorage()
	for i := 0; i < 10; i++ {
		storage.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			storage.Get(strconv.Itoa(i))
		}
	}

}

// go test -bench=BenchmarkSetGetFunc
func BenchmarkSetGetFunc(b *testing.B) {
	addTests := []Test{}
	for i := 0; i < 10; i++ {
		addTests = append(addTests, Test{strconv.Itoa(i), strconv.Itoa(i)})
	}
	storage := storage.NewStorage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range addTests {
			storage.Set(test.key, test.value)
		}
		for _, test := range addTests {
			storage.Get(test.key)
		}
	}
}
