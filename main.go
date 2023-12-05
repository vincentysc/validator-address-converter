package main

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("canto", "cantopub")
	// config.SetBech32PrefixForAccount("cro", "cropub")
	// config.SetBech32PrefixForValidator("crocncl", "crocncl")
	// config.SetBech32PrefixForConsensusNode("crocnclcons", "crocnclconspub")
	config.Seal()

	// Account address => Operator address
	accAddr, _ := sdk.AccAddressFromBech32("canto184h3psqp3ud6hp8n8z06sfm8khcv9k60m6kdv7")
	valAddr, _ := sdk.ValAddressFromHex(hex.EncodeToString(accAddr.Bytes()))
	fmt.Println("Operator address:", valAddr.String())
	fmt.Println("Evm address:", hex.EncodeToString(accAddr.Bytes()))

	// // Operator address => account address
	// valAddr, _ := sdk.ValAddressFromBech32("crocncl1amcxa06n0djycz0qd2cvvdyvd35uunnfej85g6")
	// accAddr, _ := sdk.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))
	// fmt.Println("Account address:", accAddr.String())
}
