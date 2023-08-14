package addition_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veegrace/go-benchmarking/addition"
)

func TestAddition(t *testing.T) {
	a := 3.0
	b := 4.0
	wantSum := 7.0
	gotSum := addition.Addition(a, b)
	assert.Equal(t, wantSum, gotSum, "Sum not the same")
}

// func BenchmarkAddition(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		a := 3.0
// 		b := 4.0
// 		addition.Addition(a, b)
// 	}
// }

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
