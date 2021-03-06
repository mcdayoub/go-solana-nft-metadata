package examples

import (
	"fmt"

	metadata "github.com/mcdayoub/go-solana-nft-metadata/pkg"
)

func accountOwnerExample() {
	accountOwner := "9j3Mcte8bwh97SsUBqZgApG5xieGCaXHYKCjFSwxZ14t"
	allNFTs, err := metadata.AllNFTsForAddress(accountOwner)
	if err != nil {
		fmt.Println(err)
	}
	for _, metadata := range allNFTs {
		fmt.Println(string(metadata))
	}
}

func mintAddressExample() {
	mintAddress := "9qtN3RDr8sykdzmJoDjYbAAQmjoQcAfsqj8ipT6rh413"
	metadata := metadata.NFTMetadata(mintAddress)
	fmt.Println(string(metadata))
}
