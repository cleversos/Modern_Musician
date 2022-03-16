package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"regexp"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"[^"\s].*/NonFungibleToken.cdc"`)
	placeholderMetadataViews    = regexp.MustCompile(`"[^"\s].*/MetadataViews.cdc"`)
)

const (
	filenameNonFungibleToken = "NonFungibleToken.cdc"
	filenameModernMusicianNFT       = "ModernMusicianNFT.cdc"
	filenameMetadataViews    = "MetadataViews.cdc"
)

// NonFungibleToken returns the NonFungibleToken contract interface.
func NonFungibleToken() []byte {
	return assets.MustAsset(filenameNonFungibleToken)
}

// ModernMusicianNFT returns the ModernMusicianNFT contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func ModernMusicianNFT(nftAddress, metadataAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameModernMusicianNFT)

	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())

	return []byte(code)
}

func MetadataViews() []byte {
	return assets.MustAsset(filenameMetadataViews)
}
