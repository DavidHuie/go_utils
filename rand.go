package go_utils

import (
	"math/rand"
)

// Returns a random uint64 integer
func RandUint64() uint64 {
	var x uint64 = uint64(rand.Uint32())
	var y uint64 = uint64(rand.Uint32())
	var big_rand uint64
	big_rand = big_rand | x
	big_rand = big_rand | y<<32
	return big_rand
}
