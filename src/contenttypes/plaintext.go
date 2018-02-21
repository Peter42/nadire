package contenttypes

import (
	"math/rand"
	"time"
)

type plaintext struct {
	rng *rand.Rand
}

func newPlaintext() *plaintext {
	return &plaintext{rand.New(rand.NewSource(time.Now().UnixNano()))}
}

const randomStringContent = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.:" // length of 64
func (p *plaintext) FillBuffer(buffer []byte) {
	reqSize := len(buffer)

	i := 0
	for ; i+10 < reqSize; i += 10 {
		rnum := p.rng.Uint64()
		for j := 0; j < 10; j++ {
			b := randomStringContent[rnum&63]
			buffer[i+j] = b
			rnum >>= 6
		}
	}

	for ; i < reqSize; i++ {
		buffer[i] = randomStringContent[p.rng.Uint32()&63]
	}
}
