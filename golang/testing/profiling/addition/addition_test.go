package addition_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veegrace/playground/golang/testing/profiling/addition"
)

func TestAddition(t *testing.T) {
	a := 3.0
	b := 4.0
	wantSum := 7.0
	gotSum := addition.Addition(a, b)
	assert.Equal(t, wantSum, gotSum, "Sum not the same")
}

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 3.0
		b := 4.0
		addition.Addition(a, b)
	}
}

func BenchmarkSample_Multiple(b *testing.B) {
	benchmarks := []struct {
		scenario string
		a        float64
		b        float64
		wantSum  float64
	}{
		{"both numbers positive", 3.0, 4.0, 7.0},
		{"both numbers negative", -3.0, -4.0, -7.0},
		{"larger number is negative", 3.0, -4.0, -1.0},
		{"smaller number is negative", -3.0, 4.0, 1.0},
		{"sum is zero", -3.0, 3.0, 0.0},
	}
	for _, bm := range benchmarks {
		b.Run(bm.scenario, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gotSum := addition.Addition(bm.a, bm.b)
				assert.Equal(b,
					bm.wantSum, gotSum, "Sum not the same",
				)
			}
		})
	}
}
