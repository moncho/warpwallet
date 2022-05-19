package bitcoin

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

//NewBitcoinPrivKey creates a valid Bitcoin private key from the given byte slice
//See https://en.bitcoin.it/wiki/Private_key
func NewBitcoinPrivKey(key []byte) (*ecdsa.PrivateKey, error) {
	curve := btcec.S256()

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(key)
	return priv, nil

}
