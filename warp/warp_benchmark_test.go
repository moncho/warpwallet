package warp

import "testing"

//BenchmarkSeedGeneration benchmarks warpwallet seed generation
func BenchmarkParallelSeedGeneration(b *testing.B) {
	pass := "uyVkW5vKXX3RpvnUcj7U3Q"
	salt := "zXrlmk3p5Lxr0vjJKdcJWQ"
	for i := 0; i < b.N; i++ {
		parallelSeedGeneration(pass, salt)
	}
}

func BenchmarkSerialSeedGeneration(b *testing.B) {
	pass := "uyVkW5vKXX3RpvnUcj7U3Q"
	salt := "zXrlmk3p5Lxr0vjJKdcJWQ"
	for i := 0; i < b.N; i++ {
		serialSeedGeneration(pass, salt)
	}
}
