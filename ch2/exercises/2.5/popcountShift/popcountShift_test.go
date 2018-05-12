package popcountShift

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func benchmarkPopCount(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(u)
	}
}

func benchmarkPopCountShift(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(u)
	}
}

func benchmarkPopCountShiftRecurse(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftRecurse(u)
	}
}

func benchmarkPopCountClearRightBit(u uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRightBit(u, 0)
	}
}
func BenchmarkPopCount69(b *testing.B)           { benchmarkPopCount(69, b) }
func BenchmarkPopCount420(b *testing.B)          { benchmarkPopCount(420, b) }
func BenchmarkPopCount6969(b *testing.B)         { benchmarkPopCount(6969, b) }
func BenchmarkPopCount4206969(b *testing.B)      { benchmarkPopCount(4206969, b) }
func BenchmarkPopCount694206942069(b *testing.B) { benchmarkPopCount(694206942069, b) }

func BenchmarkPopCountShift69(b *testing.B)           { benchmarkPopCountShift(69, b) }
func BenchmarkPopCountShift420(b *testing.B)          { benchmarkPopCountShift(420, b) }
func BenchmarkPopCountShift6969(b *testing.B)         { benchmarkPopCountShift(6969, b) }
func BenchmarkPopCountShift4206969(b *testing.B)      { benchmarkPopCountShift(4206969, b) }
func BenchmarkPopCountShift694206942069(b *testing.B) { benchmarkPopCountShift(694206942069, b) }

func BenchmarkPopCountShiftRecurse69(b *testing.B)      { benchmarkPopCountShiftRecurse(69, b) }
func BenchmarkPopCountShiftRecurse420(b *testing.B)     { benchmarkPopCountShiftRecurse(420, b) }
func BenchmarkPopCountShiftRecurse6969(b *testing.B)    { benchmarkPopCountShiftRecurse(6969, b) }
func BenchmarkPopCountShiftRecurse4206969(b *testing.B) { benchmarkPopCountShiftRecurse(4206969, b) }
func BenchmarkPopCountShiftRecurse694206942069(b *testing.B) {
	benchmarkPopCountShiftRecurse(694206942069, b)
}

func BenchmarkPopCountCRB69(b *testing.B)           { benchmarkPopCountClearRightBit(69, b) }
func BenchmarkPopCountCRB420(b *testing.B)          { benchmarkPopCountClearRightBit(420, b) }
func BenchmarkPopCountCRB6969(b *testing.B)         { benchmarkPopCountClearRightBit(6969, b) }
func BenchmarkPopCountCRB4206969(b *testing.B)      { benchmarkPopCountClearRightBit(4206969, b) }
func BenchmarkPopCountCRB694206942069(b *testing.B) { benchmarkPopCountClearRightBit(694206942069, b) }
