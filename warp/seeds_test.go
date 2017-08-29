package warp

import (
	"encoding/hex"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		passphrase string
		salt       string
	}
	type seeds struct {
		scryptSeed string
		pbkdf2Seed string
		xorSeed    string
	}
	tests := []struct {
		name  string
		args  args
		seeds seeds
	}{
		{
			"test 1",
			args{
				"ER8FT+HFjk0",
				"7DpniYifN6c",
			},
			seeds{
				"b58e47817de4d3901694b68bc8566ed5af9bec21e7a3bd56be114e2004ac148b",
				"daab156024167271f4f894e91213f6cd52cd243dd19c7126075d0c1debeef114",
				"6f2552e159f2a1e1e26c2262da459818fd56c81c363fcc70b94c423def42e59f",
			},
		},
		{
			"test 2",
			args{
				"YqIDBApDYME",
				"G34HqIgjrIc",
			},
			seeds{
				"145c2af767a1116477fec267e758cf57614e5d7fa175f7d9acde5b4984d44181",
				"ce5cbcf5c2d90be3e22b9e098ffc7b824827fa26f49f87fcf4b7865e47edc413",
				"da009602a5781a8795d55c6e68a4b4d52969a75955ea70255869dd17c3398592",
			},
		},
		{
			"test 3",
			args{
				"FPdAxCygMJg",
				"X+qaSwhUYXw",
			},
			seeds{
				"653083a95e681134205d7f4bf9c3f58a48bb8e7ef7715cba4e800bcc802c368a",
				"4a5a7a04c713922d43a9a103de4ff1c420c47db03b542f275be4939712b08557",
				"2f6af9ad997b831963f4de48278c044e687ff3cecc25739d1564985b929cb3dd",
			},
		},
		{
			"test 4",
			args{
				"uyVkW5vKXX3RpvnUcj7U3Q",
				"zXrlmk3p5Lxr0vjJKdcJWQ",
			},
			seeds{
				"235d2b1480eee07bdc04114fd83294facd8f252ac142685595046bb2e680e9dc",
				"239cc5df0fb8d80078e8fe5aac99552aca03dd1cab7276a4cb909083dc05dd8f",
				"00c1eecb8f56387ba4ecef1574abc1d0078cf8366a301ef15e94fb313a853453",
			},
		},
	}
	for _, tt := range tests {
		scSeed, _ := scryptSeed(tt.args.passphrase, tt.args.salt)()
		pbSeed, _ := pbkdf2Seed(tt.args.passphrase, tt.args.salt)()
		xorSeed, _ := xorSeeds(
			scryptSeed(tt.args.passphrase, tt.args.salt),
			pbkdf2Seed(tt.args.passphrase, tt.args.salt))()

		scSeedHex := hex.EncodeToString(scSeed)
		pbSeedHex := hex.EncodeToString(pbSeed)
		xorSeedHex := hex.EncodeToString(xorSeed)
		if scSeedHex != tt.seeds.scryptSeed {
			t.Errorf("%s. scryptSeed() = %s, want %s", tt.name, scSeedHex, tt.seeds.scryptSeed)
		}
		if pbSeedHex != tt.seeds.pbkdf2Seed {
			t.Errorf("%s. pbkdf2Seed() = %s, want %s", tt.name, pbSeedHex, tt.seeds.pbkdf2Seed)
		}
		if xorSeedHex != tt.seeds.xorSeed {
			t.Errorf("%s. xorSeeds() = %s, want %s", tt.name, xorSeedHex, tt.seeds.xorSeed)
		}
	}
}
