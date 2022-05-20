package bitcoin

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

const (
	//PubKeyCompressedEvenPrefix is the prefix to be used for compressed address format of even keys encoding
	PubKeyCompressedEvenPrefix = 0x02
	//PubKeyCompressedOddPrefix is the prefix to be used for compressed address format of odd keys encoding
	PubKeyCompressedOddPrefix = 0x03

	//PubKeyUncompressedVPrefix is the prefix to be used for uncompressed address format encoding
	PubKeyUncompressedVPrefix = 0x04
	//MainNetworkVPrefix is the prefix that identifes Bitcoin main network
	MainNetworkVPrefix = 0x00
)

// ToBTCAddress converts a Bitcoin public key to a compressed Bitcoin address string.
//See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses
func ToBTCAddress(pubKey ecdsa.PublicKey) (string, error) {

	//1 -
	key := serializeUncompressed(pubKey)

	//2 - Perform SHA-256
	hash1 := sha256.Sum256(key)

	//3 - Perform RIPEMD-160
	ripemd160Hash := ripemd160.New()
	_, err := ripemd160Hash.Write(hash1[:])
	if err != nil {
		return "", err
	}
	hash2 := ripemd160Hash.Sum(nil)

	return Base58CheckEncode(MainNetworkVPrefix, hash2)
}

// serializeUncompressed returns a Bitcoin public key to a 65-byte uncompressed format
// See https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses step 1
func serializeUncompressed(pub ecdsa.PublicKey) []byte {
	x := pub.X.Bytes()
	y := pub.Y.Bytes()
	if len(x) < BitcoinPrivKeyBytesLen {
		x = paddedPrepend(BitcoinPrivKeyBytesLen, x)
	}
	if len(y) < BitcoinPrivKeyBytesLen {
		y = paddedPrepend(BitcoinPrivKeyBytesLen, y)
	}

	result := append([]byte{PubKeyUncompressedVPrefix}, x...)
	result = append(result, y...)
	return result
}

// serializeCompressed serializes a public key to a 33-byte compressed format.
// See https://bitcoin.org/en/developer-guide#mini-private-key-format
func serializeCompressed(pub ecdsa.PublicKey) []byte {
	var prefix byte
	if isOdd(pub.Y) {
		prefix = PubKeyCompressedOddPrefix
	} else {
		prefix = PubKeyCompressedEvenPrefix
	}
	x := pub.X.Bytes()

	if len(x) < BitcoinPrivKeyBytesLen {
		x = paddedPrepend(BitcoinPrivKeyBytesLen, x)
	}

	return append([]byte{prefix}, x...)
}

func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}
