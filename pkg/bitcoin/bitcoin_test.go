package bitcoin

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/moncho/warpwallet/pkg/warp"
)

type warpWalletInput struct {
	passphrase string
	salt       string
}

func Test_BitcoinPrivKeyGenerationWarpWalletTestSpec(t *testing.T) {

	tests := []struct {
		name           string
		args           warpWalletInput
		seed           string
		privateKey     string
		bitcoinAddress string
	}{
		{
			"Generate bitcoin keypair from passphrase ER8FT+HFjk0",
			warpWalletInput{
				passphrase: "ER8FT+HFjk0",
				salt:       "7DpniYifN6c",
			},
			"6f2552e159f2a1e1e26c2262da459818fd56c81c363fcc70b94c423def42e59f",
			"5JfEekYcaAexqcigtFAy4h2ZAY95vjKCvS1khAkSG8ATo1veQAD",
			"1J32CmwScqhwnNQ77cKv9q41JGwoZe2JYQ",
		},
		{
			"Generate bitcoin keypair from passphrase YqIDBApDYME",
			warpWalletInput{
				passphrase: "YqIDBApDYME",
				salt:       "G34HqIgjrIc",
			},

			"da009602a5781a8795d55c6e68a4b4d52969a75955ea70255869dd17c3398592",

			"5KUJA5iZ2zS7AXkU2S8BiBVY3xj6F8GspLfWWqL9V7CajXumBQV",
			"19aKBeXe2mi4NbQRpYUrCLZtRDHDUs9J7J",
		},
		{
			"Generate bitcoin keypair from passphrase FPdAxCygMJg",
			warpWalletInput{
				passphrase: "FPdAxCygMJg",
				salt:       "X+qaSwhUYXw",
			},

			"2f6af9ad997b831963f4de48278c044e687ff3cecc25739d1564985b929cb3dd",

			"5JBAonQ4iGKFJxENExZghDtAS6YB8BsCw5mwpHSvZvP3Q2UxmT1",
			"14Pqeo9XNRxjtKFFYd6TvRrJuZxVpciS81",
		},
		{
			"Generate bitcoin keypair from passphrase gdoyAj5Y+jA",

			warpWalletInput{
				passphrase: "gdoyAj5Y+jA",
				salt:       "E+6ZzCnRqVM",
			},
			"5ab0b9ef00b03d19a6fd571a612300492233f252febb0e8aa6ab90e286fa1178",
			"5JWE9LBvFM5xRE9YCnq3FD35fDVgTPNBmksGwW2jj5fi6cSgHsC",
			"1KiiYhv9xkTZfcLYwqPhYHrSbvwJFFUgKv",
		},
		{
			"Generate bitcoin keypair from passphrase bS7kqw6LDMJbvHwNFJiXlw",
			warpWalletInput{
				passphrase: "bS7kqw6LDMJbvHwNFJiXlw",
				salt:       "tzsvA87xk+Rrw/5qj580kg",
			},

			"cc104757e5504c199a1e71dbda23a04751ab3bf44b1031952df20305ddad020a",
			"5KNA7T1peej9DwF5ZALUeqBq2xN4LHetaKvH9oRr8RHdTgwohd7",
			"17ZcmAbJ35QJzAbwqAj4evo4vL5PwA8e7C",
		},
		{
			"Generate bitcoin keypair from passphrase uyVkW5vKXX3RpvnUcj7U3Q",
			warpWalletInput{
				passphrase: "uyVkW5vKXX3RpvnUcj7U3Q",
				salt:       "zXrlmk3p5Lxr0vjJKdcJWQ",
			},
			"00c1eecb8f56387ba4ecef1574abc1d0078cf8366a301ef15e94fb313a853453",
			"5Hpcw1rqoojG7LTHo4MrEHBwmBQBXQQmH6dEa89ayw5qMXvZmEZ",
			"1ACJ7MhCRRTPaEvr2utat6FQjsQgC6qpE6",
		},
		{
			"Generate bitcoin keypair from passphrase 5HoGwwEMOgclQyzH72B9pQ",
			warpWalletInput{
				passphrase: "5HoGwwEMOgclQyzH72B9pQ",
				salt:       "UGKv/5nY3ig8bZvMycvIxQ",
			},
			"265486f06e0a90da5b0202ec41a1d28e5614daef13a3ed3cc5668135454ac08e",
			"5J7Ag5fBArgKN9ocVJs4rcQw1chZjHrqAb4YRuny6YiieJc5iG3",
			"1Mtb2o7AsTRAR3vjtSYjq1rgB8Q6A76avD",
		},
		{
			"Generate bitcoin keypair from passphrase TUMBDBWh8ArOK0+jO5glcA",
			warpWalletInput{
				passphrase: "TUMBDBWh8ArOK0+jO5glcA",
				salt:       "dAMOvN2WaEUTC/V5yg0eQA",
			},

			"f52bb6f6d755b9cf972a7962f9f288d9d91d94b35e6ac6fc4f5ce02dbdc4886d",
			"5KgG93ePJJ8HC2tnTerThNUnXbjyeBpUCBDRn5ZxMRB9GxiwJEK",
			"1B2VuTAHERd2GmBK522LFLUTwYWcW1vXH6",
		},
		{
			"Generate bitcoin keypair from passphrase rDrc5eIhSt2qP8pnpnSMu1u2/mP6KTqS",
			warpWalletInput{
				passphrase: "rDrc5eIhSt2qP8pnpnSMu1u2/mP6KTqS",
				salt:       "HGM1/qHoT3XX61NXw8H1nQ"},

			"13a7fdb19b0a5501c21020b5d06870f19d1eec7c705ee744aed2f31bd8c45b0f",
			"5HxwfzgQ2yem9uY5UxdiaKYPgUR251FCVHw1VuHC5bSW5NVLaok",
			"12XD7BtiU1gydRzQm3cAoui2RQjhVJfNPg",
		},

		{
			"Generate bitcoin keypair from passphrase Brd8TB3EDhegSx2wy2ffW0oGNC29vkCo",
			warpWalletInput{
				passphrase: "Brd8TB3EDhegSx2wy2ffW0oGNC29vkCo",
				salt:       "dUBIrYPiUZ6BD/l+zBhthA",
			},

			"bbf6322cb1006b475f224644a53a9afa0ae497507ff3b5c060e4ad31f4ca21c6",
			"5KF4ozGWXGZAqNydQg65JQ4XnJaUpBkU9g59C287GrbLfWVmYHL",
			"1CD93Tgj74uKh87dENR2GMWB1kpCidLZiS",
		},
		{
			"Generate bitcoin keypair from passphrase eYuYtFxU4KrePYrbHSi/8ncAKEb+KbNH",
			warpWalletInput{
				passphrase: "eYuYtFxU4KrePYrbHSi/8ncAKEb+KbNH",
				salt:       "le5MMmWaj4AlGcRevRPEdw",
			},

			"b5b5fe55e8e332ebe469c54d5672b69ace7acb40c5c032a4fb5b75c1bd0b409e",

			"5KCK9EtgvjsQcPcZcfMoqcHwZKzA1MLfPUvDCYE1agiNf56CfAk",
			"18mugeQN8uecTBE9psW2uQrhRBXZJkhyB7",
		},
		{
			"Generate bitcoin keypair from passphrase TRGmdIHpnsSXjEnLc+U+MrRV3ryo8trG",

			warpWalletInput{

				passphrase: "TRGmdIHpnsSXjEnLc+U+MrRV3ryo8trG",
				salt:       "DhZNEt9hx08i6uMXo5DOyg",
			},
			"739112892d83ddb8513c6f3879e8f630f241a3e66fcc08e505e12e7fec6627a8",

			"5JhBaSsxgNBjvZWVfdVQsnMzYf4msHMQ7HRaHLvvMy1CEgsTstg",
			"19QCgqHnKw8vrJph7wWP3nKg9tFixqYwiB",
		},
	}
	for _, tt := range tests {
		seed, err := warp.NewSeed(tt.args.passphrase, tt.args.salt)

		if err != nil {
			t.Errorf("%q. seed generation error = %v", tt.name, err)
			continue
		}

		if hex.EncodeToString(seed) != tt.seed {
			t.Errorf("%q. seed() = %v, want %v", tt.name, hex.EncodeToString(seed), tt.seed)
			continue

		}

		got, err := NewBitcoinPrivKey(seed)

		if err != nil {
			t.Errorf("%q. NewBitcoinPrivKey() error = %v", tt.name, err)
			continue
		}

		wif, err := PrivateToWIF(*got)

		if err != nil {
			t.Errorf("%q. PrivateToWIF() error = %v", tt.name, err)
			continue
		}
		if wif != tt.privateKey {
			t.Errorf(
				"%s. PrivateToWIF() = %v, want %v",
				tt.name,
				wif,
				tt.privateKey)
			continue

		}
		publicAddress, err := ToBTCAddress(got.PublicKey)

		if err != nil {
			t.Errorf("%q. ToBTCAddress() error = %v", tt.name, err)
			continue
		}
		if publicAddress != tt.bitcoinAddress {
			t.Errorf(
				"%s. ToBTCAddress() = %v, want %v",
				tt.name,
				publicAddress,
				tt.bitcoinAddress)
		}
	}
}

func Test_BitcoinPrivKeyGeneration(t *testing.T) {

	tests := []struct {
		name           string
		args           warpWalletInput
		privateKey     string
		bitcoinAddress string
	}{
		{
			"Generate bitcoin keypair from passphrase 'correct horse battery staple'",
			warpWalletInput{
				passphrase: "correct horse battery staple",
				salt:       "correct@horse.battery",
			},
			"5JYTXQWoMaG2GLysDzVWfdMWazqHUyrcjpQwNbnReg7vLQzhFLQ",
			"13w7aHdp6bYSjxCFj47h2hMKaE5xyQNG84",
		},
	}

	for _, tt := range tests {
		seed, err := warp.NewSeed(tt.args.passphrase, tt.args.salt)

		if err != nil {
			t.Errorf("%q. seed generation error = %v", tt.name, err)
			continue
		}

		got, err := NewBitcoinPrivKey(seed)

		if err != nil {
			t.Errorf("%q. NewBitcoinPrivKey() error = %v", tt.name, err)
			continue
		}
		wif, err := PrivateToWIF(*got)

		if err != nil {
			t.Errorf("%q. PrivateToWIF() error = %v", tt.name, err)
			continue
		}
		if wif != tt.privateKey {
			t.Errorf(
				"%s. PrivateToWIF() = %v, want %v",
				tt.name,
				wif,
				tt.privateKey)
		}

		publicAddress, err := ToBTCAddress(got.PublicKey)

		if err != nil {
			t.Errorf("%q. ToBTCAddress() error = %v", tt.name, err)
			continue
		}

		if publicAddress != tt.bitcoinAddress {
			t.Errorf(
				"%s. ToBTCAddress() = %v, want %v",
				tt.name,
				publicAddress,
				tt.bitcoinAddress)
		}
	}

}

func Test_paddedPrepend(t *testing.T) {
	type args struct {
		size uint
		src  []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"paddedPrepend does not do anything when slice has the wanted length",
			args{
				16,
				[]byte("YELLOW SUBMARINE"),
			},
			[]byte("YELLOW SUBMARINE"),
		},
		{
			"paddedPrepend prepends 0 to reach the expected length",
			args{
				18,
				[]byte("YELLOW SUBMARINE"),
			},
			[]byte("\x00\x00YELLOW SUBMARINE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paddedPrepend(tt.args.size, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paddedPrepend() = %v, want %v", got, tt.want)
			}
		})
	}
}
