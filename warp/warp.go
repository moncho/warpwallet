package warp

//NewSeed creates a seed from the given passphrase and salt
//using https://keybase.io/warp algorithm for generating the seed.
func NewSeed(passphrase, salt string) ([]byte, error) {
	return parallelSeedGeneration(passphrase, salt)
}

//serialSeedGeneration generates the seed by running algorithm steps one at a time
func serialSeedGeneration(passphrase, salt string) ([]byte, error) {
	k1, err := scryptSeed(passphrase, salt)()
	if err != nil {
		return nil, err
	}
	k2, err := pbkdf2Seed(passphrase, salt)()
	if err != nil {
		return nil, err
	}

	blockXOR(k1, k2)
	return k1, nil
}

//parallelSeedGeneration runs in parallel intermediate steps of the seed generation algorithm
func parallelSeedGeneration(passphrase, salt string) ([]byte, error) {
	return xorSeeds(
		scryptSeed(passphrase, salt),
		pbkdf2Seed(passphrase, salt))()
}
