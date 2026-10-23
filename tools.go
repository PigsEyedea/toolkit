package toolkit

import (
	"crypto/rand"
	"math/big"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type with have acces
// to the pointer methods

type Tools struct{}

// retruns random string of length n
func (t *Tools) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomStringSource)
	randMax := big.NewInt(int64(len(r)))

	for idx := range s {

		randIdx, err := rand.Int(rand.Reader, randMax)

		if err != nil {
			panic(err)
		}

		s[idx] = r[randIdx.Int64()]

	}

	return string(s)
}
