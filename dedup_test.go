package dedup

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

/*
BenchmarkAccessStructure show compare metrics between data strucuture and number of items.
*/
func BenchmarkDedup(b *testing.B) {
	for size := 2; size < 100000000; size = size * 2 {
		benchmarkDedup(b, size)
	}
}

func benchmarkDedup(b *testing.B, size int) {
	var strings = make([]string, size, size)

	for i := 0; i < size; i++ {
		u, err := uuid.NewV4()
		if err != nil {
			panic("UUUUUID!")
		}
		strings[i] = u.String()
	}

	b.ResetTimer()

	b.Run(
		fmt.Sprintf("Dedup_justMap_%d", size),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := justMap()
				_ = dedup(strings, m)
			}
		},
	)

	b.Run(
		fmt.Sprintf("Dedup_preinitMap_%d", size),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := preinitMap(size)
				_ = dedup(strings, m)
			}
		},
	)
}
