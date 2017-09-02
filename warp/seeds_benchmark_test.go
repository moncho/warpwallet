package warp

import "testing"

//BenchmarkSeedGeneration benchmarks warpwallet seed generation
func BenchmarkSeedGeneration(b *testing.B) {
	pass := "uyVkW5vKXX3RpvnUcj7U3Q"
	salt := "zXrlmk3p5Lxr0vjJKdcJWQ"
	for i := 0; i < b.N; i++ {
		xorSeeds(
			scryptSeed(pass, salt),
			pbkdf2Seed(pass, salt))()
	}
}
