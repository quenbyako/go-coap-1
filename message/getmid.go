package message

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	mathRand "math/rand"
	"sync/atomic"
	"time"
)

func init() {
	mathRand.Seed(time.Now().UnixNano())
}

var msgID = uint32(RandMID())

// GetMID generates a message id for UDP. (0 <= mid <= 65535)
func GetMID() int32 {
	return int32(uint16(atomic.AddUint32(&msgID, 1)))
}

func RandMID() int32 {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		// fallback to cryptographically insecure pseudo-random generator
		return int32(uint16(mathRand.Uint32() >> 16))
	}
	return int32(uint16(binary.BigEndian.Uint32(b)))
}

// ValidateMID validates a message id for UDP. (0 <= mid <= 65535)
func ValidateMID(mid int32) bool {
	return mid >= 0 && mid <= math.MaxUint16
}
