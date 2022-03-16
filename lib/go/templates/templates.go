package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"[^"\s].*/NonFungibleToken.cdc"`)
	placeholderModernMusicianNFT       = regexp.MustCompile(`"[^"\s].*/ModernMusicianNFT.cdc"`)
	placeholderMetadataViews    = regexp.MustCompile(`"[^"\s].*/MetadataViews.cdc"`)
)

func replaceAddresses(code string, nftAddress, ModernMusicianNFTAddress, metadataAddress flow.Address) []byte {
	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderModernMusicianNFT.ReplaceAllString(code, "0x"+ModernMusicianNFTAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())
	return []byte(code)
}
