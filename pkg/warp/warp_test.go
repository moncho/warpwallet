package warp

import (
	"encoding/hex"
	"testing"
)

type args struct {
	passphrase string
	salt       string
}

var tests = []struct {
	name    string
	args    args
	want    string
	wantErr bool
}{
	{
		"test 1",
		args{
			"ER8FT+HFjk0",
			"7DpniYifN6c",
		},
		"6f2552e159f2a1e1e26c2262da459818fd56c81c363fcc70b94c423def42e59f",
		false,
	},
	{
		"test 2",
		args{
			"YqIDBApDYME",
			"G34HqIgjrIc",
		},
		"da009602a5781a8795d55c6e68a4b4d52969a75955ea70255869dd17c3398592",
		false,
	},
	{
		"test 3",
		args{
			"FPdAxCygMJg",
			"X+qaSwhUYXw",
		},
		"2f6af9ad997b831963f4de48278c044e687ff3cecc25739d1564985b929cb3dd",
		false,
	},
	{
		"test 4",
		args{
			"uyVkW5vKXX3RpvnUcj7U3Q",
			"zXrlmk3p5Lxr0vjJKdcJWQ",
		},
		"00c1eecb8f56387ba4ecef1574abc1d0078cf8366a301ef15e94fb313a853453",
		false,
	},
}

func Test_parallelSeedGeneration(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parallelSeedGeneration(tt.args.passphrase, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("parallelSeedGeneration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if hex.EncodeToString(got) != tt.want {
				t.Errorf("%q. parallelSeedGeneration() = %v, want %v", tt.name, hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_serialSeedGeneration(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serialSeedGeneration(tt.args.passphrase, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("serialSeedGeneration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if hex.EncodeToString(got) != tt.want {
				t.Errorf("%q. serialSeedGeneration() = %v, want %v", tt.name, hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func TestNewSeed(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSeed(tt.args.passphrase, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if hex.EncodeToString(got) != tt.want {
				t.Errorf("%q. NewSeed() = %v, want %v", tt.name, hex.EncodeToString(got), tt.want)
			}
		})
	}
}
