package warp

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

type result struct {
	seed []byte
	err  error
}

type seedGenerator func() ([]byte, error)

func scryptSeed(passphrase, salt string) seedGenerator {
	return func() ([]byte, error) {
		return scrypt.Key([]byte(passphrase+"\u0001"), []byte(salt+"\u0001"), 262144, 8, 1, 32)
	}

}

func pbkdf2Seed(passphrase, salt string) seedGenerator {
	return func() ([]byte, error) {
		dk := pbkdf2.Key([]byte(passphrase+"\u0002"), []byte(salt+"\u0002"), 65536, 32, sha256.New)
		return dk, nil
	}
}

//xorSeeds xor's the result of the given seedGenerators
func xorSeeds(generators ...seedGenerator) seedGenerator {
	return func() ([]byte, error) {

		genCount := len(generators)
		results := make(chan result, genCount)
		done := make(chan struct{})
		defer close(results)
		defer close(done)

		for _, g := range generators {

			go func(s seedGenerator) {
				b, err := s()
				select {
				case results <- result{b, err}:
				case <-done:
				}

			}(g)
		}

		var finalKey []byte

		for i := 0; i < genCount; i++ {
			r := <-results
			if r.err != nil {
				return nil, r.err
			}
			if finalKey == nil {
				finalKey = r.seed
			} else {
				blockXOR(finalKey, r.seed)
			}
		}

		return finalKey, nil
	}
}

func blockXOR(dst, src []byte) {
	for i, v := range src {
		dst[i] ^= v
	}
}
