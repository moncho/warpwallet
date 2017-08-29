package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/moncho/warpwallet/bitcoin"
	"github.com/moncho/warpwallet/warp"
)

func main() {

	stat, _ := os.Stdin.Stat()
	piped := false
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		piped = true
	}
	pass, salt := askForSeeds(piped)

	key, err := warp.NewSeed(pass, salt)

	//log.Printf("Generating wallet seed from %s, %s", pass, salt)

	if err != nil {
		fmt.Printf("Could not generate wallet seed from %s, %s: %s", pass, salt, err.Error())
		os.Exit(-1)
	}

	wif, pubAddress := generate(key)
	if wif == "" || pubAddress == "" {
		os.Exit(-1)
	}
	fmt.Printf("Public bitcoin address: %v\n", pubAddress)
	fmt.Printf("Private key (don't share): %v\n", wif)

}

func askForSeeds(piped bool) (string, string) {

	reader := bufio.NewReader(os.Stdin)
	if !piped {
		fmt.Print("Passphrase: ")
	}
	pass, _ := reader.ReadString('\n')
	pass = strings.Trim(pass, "\n")
	if !piped {
		fmt.Print("Your email [as a salt]: ")
	}
	salt, _ := reader.ReadString('\n')
	salt = strings.Trim(salt, "\n")
	return pass, salt
}

func generate(key []byte) (string, string) {
	priv, err := bitcoin.NewBitcoinPrivKey(key)
	if err != nil {
		fmt.Printf("Could not generate wallet private key: %s", err.Error())
		return "", ""
	}
	wif, err := bitcoin.PrivateToWIF(*priv)
	if err != nil {
		fmt.Printf("Could not generate WIF from private key: %s", err.Error())
		return "", ""
	}
	address, err := bitcoin.ToBTCAddress(priv.PublicKey)
	if err != nil {
		fmt.Printf("Could not generate compressed Bitcoin address from public key: %s", err.Error())
		return "", ""
	}

	return wif, address
}
