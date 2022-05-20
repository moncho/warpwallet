package bitcoin

import (
	"bytes"
	"testing"
)

func BenchmarkBase58Encode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b58encode(data)
	}
}
