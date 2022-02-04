package main

import (
	"flag"
	"fmt"

	metadata "github.com/mcdayoub/go-solana-nft-metadata/pkg"
)

func main() {
	command := flag.String("command", "account", "a get command, either 'account' or 'nft'")
	address := flag.String("address", "", "a solana address")

	flag.Parse()

	if *command == "account" {
		fmt.Println(metadata.AllNFTsForAddress(*address))
	} else if *command == "nft" {
		fmt.Println(metadata.NFTMetadata(*address))
	}
}
