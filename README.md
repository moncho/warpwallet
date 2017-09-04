[![GitHub license](https://img.shields.io/badge/license-MIT-green.svg)](https://raw.githubusercontent.com/moncho/warpwallet/master/LICENSE)
[![Build Status](https://travis-ci.org/moncho/warpwallet.svg?branch=master)](https://travis-ci.org/moncho/warpwallet)
[![codecov](https://codecov.io/gh/moncho/warpwallet/branch/master/graph/badge.svg)](https://codecov.io/gh/moncho/warpwallet)
[![Go Report Card](https://goreportcard.com/badge/github.com/moncho/warpwallet)](https://goreportcard.com/report/github.com/moncho/warpwallet)

# Warpwallet

**warpwallet** is a deterministic bitcoin address generator for the command line.

The address-generation algorithm has been implemented as described in [Keybase/warp](https://keybase.io/warp). All credit for the algorithm goes to the original authors.

## Disclaimer

Handle with care, if you intend to use this tool for generating a Bitcoin wallet, make sure you understand the [risks](https://eprint.iacr.org/2016/103.pdf) of using [brain wallets](https://en.bitcoin.it/wiki/Brainwallet). 

Said that, I am fairly confident that the implementation faithfully matches the original implementation (all tests created for the original tool have been added and are passing). 
So, if you feel confident that the original tool is a safe way to generate a Bitcoin wallet, this tool should be safe for usage as well.

## Usage

Run `warpwallet` on a terminal, it will ask for a passphrase and a salt, then it will generate a private key on [WIF](https://en.bitcoin.it/wiki/Wallet_import_format) format and the corresponding [version 1 Bitcoin public address](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses). 

## Why

I thought it would be cool to run **warpwallet** on a terminal. 

