package main

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("cro", "cropub")
	config.SetBech32PrefixForValidator("crocncl", "crocncl")
	config.SetBech32PrefixForConsensusNode("crocnclcons", "crocnclconspub")
	config.Seal()

	// // Account address => Operator address
	// accAddr, _ := sdk.AccAddressFromBech32("cro1amcxa06n0djycz0qd2cvvdyvd35uunnf6lya2x")
	// valAddr, _ := sdk.ValAddressFromHex(hex.EncodeToString(accAddr.Bytes()))
	// fmt.Println("Operator address:", valAddr.String())

	// Operator address => account address
	valAddr, _ := sdk.ValAddressFromBech32("crocncl1amcxa06n0djycz0qd2cvvdyvd35uunnfej85g6")
	accAddr, _ := sdk.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))
	fmt.Println("Account address:", accAddr.String())
}
