package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/moncho/warpwallet/bitcoin"
	"github.com/moncho/warpwallet/qr"
	"github.com/moncho/warpwallet/terminal"
	"github.com/moncho/warpwallet/warp"
)

func main() {

	piped := seedsFromPipe()
	pass, salt := promptSeeds(piped)

	if pass == "" || salt == "" {
		fmt.Println("A pass and some salt are required")
		os.Exit(-1)
	}

	key, err := warp.NewSeed(pass, salt)

	if err != nil {
		fmt.Printf("Could not generate wallet seed from %s, %s: %s", pass, salt, err.Error())
		os.Exit(-1)
	}

	wif, pubAddress := generate(key)
	if wif == "" || pubAddress == "" {
		os.Exit(-1)
	}

	var wifQR bytes.Buffer
	if err := qr.Copy(&wifQR, strings.NewReader(wif)); err != nil {
		fmt.Printf("Could not generate QR code for WIF: %s", err.Error())
		os.Exit(-1)
	}
	var pubAddressQR bytes.Buffer

	if err := qr.Copy(&pubAddressQR, strings.NewReader(pubAddress)); err != nil {
		fmt.Printf("Could not generate QR code for pubAddress: %s", err.Error())
		os.Exit(-1)
	}
	print(wif, wifQR.String(), pubAddress, pubAddressQR.String())
}

func print(wif, wifQR, pubAddress, pubAddressQR string) {

	fmt.Printf("\n%s %s\t\t %s %s\n\n",
		terminal.White("Public bitcoin address:"),
		terminal.Yellow(pubAddress),
		terminal.White("Private key (don't share):"),
		terminal.Red(wif))

	pub := strings.Split(pubAddressQR, "\n")
	wif2 := strings.Split(wifQR, "\n")

	for i, line := range pub {
		fmt.Printf("%s\t\t%s\n", string(line), string(wif2[i]))

	}
}

func promptSeeds(piped bool) (string, string) {

	reader := bufio.NewReader(os.Stdin)
	if !piped {
		fmt.Print(
			terminal.White("Passphrase: "))
	}
	pass, _ := reader.ReadString('\n')
	pass = strings.Trim(pass, "\n")
	if !piped {
		fmt.Print(
			terminal.White("Your email [as a salt]: "))
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

func seedsFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	piped := false
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		piped = true
	}
	return piped
}
