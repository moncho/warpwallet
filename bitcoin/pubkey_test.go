package bitcoin

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"reflect"
	"testing"
)

func Test_isOdd(t *testing.T) {
	type args struct {
		a *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test odd",
			args{
				big.NewInt(3),
			},
			true,
		},
		{
			"test even",
			args{
				big.NewInt(2),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.a); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serializeUncompressed(t *testing.T) {
	type args struct {
		pub ecdsa.PublicKey
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"ECDSA public key uncompressed serialization",
			args{
				ecdsa.PublicKey{
					X: big.NewInt(0),
					Y: big.NewInt(1),
				},
			},
			append(append([]byte{'\x04'}, bytes.Repeat([]byte{'\x00'}, 63)...), '\x01'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serializeUncompressed(tt.args.pub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serializeUncompressed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serializeCompressed(t *testing.T) {
	type args struct {
		pub ecdsa.PublicKey
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"ECDSA public key compressed serialization",
			args{
				ecdsa.PublicKey{
					X: big.NewInt(1),
					Y: big.NewInt(2),
				},
			},
			append(append([]byte{PubKeyCompressedEvenPrefix}, bytes.Repeat([]byte{'\x00'}, 31)...), '\x01'),
		},
		{
			"ECDSA public key compressed serialization",
			args{
				ecdsa.PublicKey{
					X: big.NewInt(1),
					Y: big.NewInt(1),
				},
			},
			append(append([]byte{PubKeyCompressedOddPrefix}, bytes.Repeat([]byte{'\x00'}, 31)...), '\x01'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serializeCompressed(tt.args.pub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serializeCompressed() = %v, want %v", got, tt.want)
			}
		})
	}
}
