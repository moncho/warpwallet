package bitcoin

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
)

const (
	//Base58BitcoinSymbolChart see https://en.bitcoin.it/wiki/Base58Check_encoding
	Base58BitcoinSymbolChart     = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	base58BitcoinSymbolChartIdx0 = "1"

	//PrivateKeyVersionPrefix is the prefix of WIF encoded strings
	//see https://en.bitcoin.it/wiki/List_of_address_prefixes
	PrivateKeyVersionPrefix = 0x80

	// BitcoinPrivKeyBytesLen defines the length in bytes of a private key for Bitcoin
	BitcoinPrivKeyBytesLen = 32
)

var zero = big.NewInt(0)
var base = big.NewInt(58)

//PrivateToWIF encodes the given private key to WIF, following instructions
//on https://en.bitcoin.it/wiki/Wallet_import_format
func PrivateToWIF(key ecdsa.PrivateKey) (string, error) {
	b := key.D.Bytes()
	if len(b) < BitcoinPrivKeyBytesLen {
		b = paddedPrepend(BitcoinPrivKeyBytesLen, b)
	}
	return Base58CheckEncode(PrivateKeyVersionPrefix, b)
}

func b58encode(b []byte) string {

	//Convert bytes to big integer
	x := new(big.Int).SetBytes(b)
	r := new(big.Int)
	var result string

	for x.Cmp(zero) > 0 {
		// x, r = (x / 58, x % 58)
		x, r = x.QuoRem(x, base, r)
		result = string(Base58BitcoinSymbolChart[r.Int64()]) + result
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		result = base58BitcoinSymbolChartIdx0 + result
	}

	return result
}

//Base58CheckEncode encodes the given byte array in Bitcoin into human-typable strings.
//See https://en.bitcoin.it/wiki/Base58Check_encoding Creating a Base58Check string
func Base58CheckEncode(ver uint8, b []byte) (string, error) {

	//1. Take the version byte and payload bytes, and concatenate them together (bytewise).
	bcpy := append([]byte{ver}, b...)

	//2. Take the first four bytes of SHA256(SHA256(results of step 1))
	hash1 := sha256.Sum256(bcpy)
	hash2 := sha256.Sum256(hash1[:])

	//3. Concatenate the results of step 1 and the results of step 2 together (bytewise).
	bcpy = append(bcpy, hash2[0:4]...)

	//4. Convert to base-58
	result := b58encode(bcpy)

	return result, nil
}

// paddedPrepend prepends zero to the given byte slice
// until the slice is of the expected given size.
func paddedPrepend(size uint, src []byte) []byte {
	var zeroes []byte
	for i := 0; i < int(size)-len(src); i++ {
		zeroes = append(zeroes, 0)
	}
	return append(zeroes, src...)
}
