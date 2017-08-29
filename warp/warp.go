package warp

//NewSeed creates a seed from the given passphrase and salt
//using https://keybase.io/warp algorithm for generating the seed.
func NewSeed(passphrase, salt string) ([]byte, error) {
	return xorSeeds(
		scryptSeed(passphrase, salt),
		pbkdf2Seed(passphrase, salt))()
}
