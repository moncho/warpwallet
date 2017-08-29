package warp

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

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

func xorSeeds(seeds ...seedGenerator) seedGenerator {
	return func() ([]byte, error) {
		var keys [][]byte
		var err error
		var finalKey []byte
		for _, s := range seeds {
			if b, sErr := s(); sErr == nil {
				keys = append(keys, b)
			} else {
				err = sErr
			}
		}

		if len(keys) > 0 {
			finalKey = make([]byte, len(keys[0]))
			copy(finalKey, keys[0])
			for _, key := range keys[1:] {
				blockXOR(finalKey, key)
			}
		}
		return finalKey, err
	}
}

func blockXOR(dst, src []byte) {
	for i, v := range src {
		dst[i] ^= v
	}
}
