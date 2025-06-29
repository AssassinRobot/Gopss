package generator

import (
	"os"
	"testing"
)

func BenchmarkGenerator(b *testing.B) {
	file, err := os.CreateTemp("","password_list_*.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(file.Name())

	infos := []string{"John", "Doe", "123456789", "1990", "01", "01"}

	for i := 0; i < b.N; i++ {
		Generate(file, infos)

		file.Seek(0, 0)
	}
}

