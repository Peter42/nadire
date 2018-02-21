package contenttypes

import (
	"math/rand"
	"time"
)

type rawdata struct {
	rng *rand.Rand
}

func newRawdata() *rawdata {
	return &rawdata{rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (r *rawdata) FillBuffer(buffer []byte) {
	reqSize := len(buffer)

	i := 0
	for ; i+8 < reqSize; i += 8 {
		rnum := r.rng.Uint64()

		// This loop seems imperformant/unnecessary to me (cast to int64 pointer and write directly)
		// But maybe casting around kills compiler optimizations
		// TODO: Benchmark both options
		for j := 0; j < 8; j++ {
			buffer[i+j] = byte(rnum & 255)
			rnum >>= 8
		}
	}

	for ; i < reqSize; i++ {
		buffer[i] = byte(r.rng.Uint32() & 255)
	}
}
