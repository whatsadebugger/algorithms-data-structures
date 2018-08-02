package selection

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const size int = 100000

var list = make([]int, size)
var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

func TestSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected []int
		input    []int
	}{
		{"positive values", []int{1, 2, 3, 4, 5, 6, 7}, []int{3, 4, 1, 5, 7, 2, 6}},
		{"negative values", []int{-7, -6, -5, -4, -3, -2, -1}, []int{-1, -3, -6, -2, -4, -5, -7}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Sort(test.input)
			assert.True(t, Equal(test.input, test.expected), "sorted array was not equal to expected value")
		})
	}
}

// Equal tells whether a and b contain the same elements.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		Sort(list)
	}
}
