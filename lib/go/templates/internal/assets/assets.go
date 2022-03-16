// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/destroy_nft.cdc (509B)
// ../../../transactions/mint_nft.cdc (1.324kB)
// ../../../transactions/scripts/borrow_nft.cdc (581B)
// ../../../transactions/scripts/get_collection_length.cdc (465B)
// ../../../transactions/scripts/get_nft_metadata.cdc (1.371kB)
// ../../../transactions/scripts/get_total_supply.cdc (118B)
// ../../../transactions/setup_account.cdc (928B)
// ../../../transactions/transfer_nft.cdc (1.179kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _destroy_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x6f\xe2\x40\x0c\x85\xef\xf9\x15\x4f\x39\xec\x86\xc3\x26\x97\xd5\x1e\x22\xb6\x08\xd1\x22\x71\x41\x55\x4b\x7f\xc0\x30\x71\xc8\xa8\xc1\x8e\x1c\x47\x69\x55\xf1\xdf\xab\x40\x49\xa1\xc2\xa7\xd1\xf8\x7d\xef\xd9\x0e\xfb\x46\xd4\xb0\x16\x5e\x76\xbc\x0b\xdb\x9a\x36\xf2\x4a\x8c\x52\x65\x8f\x38\x4d\x33\x2f\x6c\xea\xbc\xb5\xd9\x4f\x4d\xea\x0b\x1f\x47\x5f\x06\x0f\x6f\x6e\xdf\xd4\xb4\x5e\x6e\x6e\xa1\xdf\xdd\x13\x14\x99\x3a\x6e\x9d\xb7\x20\x9c\x84\x22\xc7\xcb\x8a\xed\xdf\xdf\x09\x3e\x22\x00\x68\x94\x1a\xa7\x94\xb4\x61\xc7\xa4\x39\xe6\x9d\x55\x73\xef\xa5\x63\x3b\x4b\x86\xaa\xc9\xe0\xa5\xae\xe9\xe8\xf3\x44\x25\xfe\xe3\x84\xa4\x5b\x51\x95\x7e\xfa\xeb\x22\x78\x31\x2a\xef\x92\x61\xc4\x1c\x37\x9b\xcf\x26\xea\x76\xf4\xe8\xac\x9a\x8c\x49\x43\xcd\x66\x68\x1c\x07\x9f\xc4\x0b\xe9\xea\x02\x2c\x86\x53\x0c\x1c\x94\x4a\x52\x62\x4f\x30\x81\x55\x04\xe9\x99\xf4\x77\x7b\x31\x5f\x3c\x89\x46\xbf\x2c\x43\x1f\xac\x2a\xd4\xf5\x47\xf5\x78\xb6\xdb\xe8\xd5\xc6\x5c\x1a\xa6\x7f\xae\x17\x4f\xcf\x6e\xc9\xf9\xb1\xba\xcf\x11\x8a\x8b\xc8\x82\x5a\x53\x79\x1f\xf0\xe3\xdf\x21\x3a\x44\x9f\x01\x00\x00\xff\xff\xbd\x0b\x03\xcd\xfd\x01\x00\x00"

func destroy_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_destroy_nftCdc,
		"destroy_nft.cdc",
	)
}

func destroy_nftCdc() (*asset, error) {
	bytes, err := destroy_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "destroy_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0x4f, 0x57, 0xb2, 0xc6, 0x27, 0xe0, 0x9e, 0xca, 0x4e, 0x50, 0xc7, 0xe3, 0xd1, 0xd3, 0xb1, 0x36, 0xe8, 0x38, 0x59, 0xb7, 0x69, 0x24, 0x74, 0x82, 0x70, 0x76, 0x1a, 0x11, 0x9c, 0xeb, 0x86}}
	return a, nil
}

var _mint_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xcd\x8e\x9b\x40\x0c\xbe\xf3\x14\x56\x0e\x2d\x91\x56\x70\x8f\xda\x5d\xa5\x51\x23\xf5\xd0\x68\xa5\xcd\x0b\x0c\x83\x03\x56\x07\x0f\x9a\x31\x9b\xad\x56\x79\xf7\x6a\x18\xc2\x4f\x36\x6d\x39\x24\x60\xfb\xfb\x6c\x7f\xb6\xa9\x69\xad\x13\x38\x58\xde\x77\x5c\x51\x61\xf0\x68\x7f\x21\xc3\xc9\xd9\x06\x56\x59\x96\x6b\xcb\xe2\x94\x16\x9f\xdf\xc6\x64\xba\xd4\xab\x64\x20\xf8\xfe\xa6\x9a\xd6\xe0\x61\x7f\xbc\x07\x9d\xbc\x11\x94\xe4\x39\x1c\x6b\xf2\xe0\xb5\xa3\x56\xa0\xf3\xe8\x41\x6a\x84\xc3\xfe\xf8\x93\x58\xd0\x81\x43\x6f\x3b\xa7\x11\xc4\x42\x43\x2c\xa0\x80\xf1\x1c\x02\x02\xf8\x87\x40\xd3\x79\x81\x02\xc1\x75\x0c\x67\x92\xba\xc7\x2b\xad\x6d\xc7\x02\x52\x2b\x81\x5a\x45\xd2\x66\xc9\x18\xf0\x5e\xac\xc3\x12\x88\x21\x0f\xaf\xaa\xc2\x7c\x4c\x9d\x24\xe2\x14\x7b\xa5\x85\x2c\xa7\x09\x00\x80\x43\x4d\x2d\x21\xcb\x06\xb6\x65\xe9\xd0\xfb\x87\xde\xce\xaa\xc1\x0d\xbc\x88\x23\xae\xa2\xa5\xc4\xd8\x12\x59\x5e\x3a\xa4\xee\x9a\x82\x15\x99\xc9\xbc\x86\xf7\xa4\xf7\xe5\x39\x18\xab\x95\x81\x57\xe5\x48\x15\x06\xe1\x64\x5d\x5f\x23\x71\xb5\x6c\xe1\x84\x0e\x59\x63\x0f\x33\x28\x83\x63\x03\x9f\x66\x12\xcf\x3a\x09\x61\xad\xc3\x56\x39\x4c\x3d\x55\x1c\x42\xb7\x9d\xd4\xdb\xa8\x53\xa8\x00\x86\x27\xcf\xa1\xb0\xce\xd9\x33\xa8\x29\x4f\x50\xff\x2f\x73\x21\x86\x41\xba\x91\xc2\xa3\x39\x65\x43\xa9\x5f\x21\xe6\xcb\x22\xe9\x97\xbb\x05\x3e\xa6\x61\x59\x36\xb3\xed\xc9\xa2\xe3\x25\x32\x3f\x2b\xa9\xd7\x23\x7d\x78\x9e\x9e\xa0\x55\x4c\x3a\x5d\xed\x6c\x67\x4a\x60\x2b\xff\x2e\x7b\x50\x68\x15\x69\x2e\x51\x12\x7c\x43\xdd\x09\x2e\xbb\xff\x16\x69\x02\x6e\x9c\xf7\x67\x0f\x6d\x57\x18\xd2\x3d\x95\xb6\xc6\x60\xbf\x16\x37\x93\xb8\x4e\xc3\xa1\x46\x7a\xed\xbb\xaf\x50\x06\x91\xd3\x91\x6d\xd9\x4a\x56\xa1\xec\x54\xab\x0a\x32\x24\xbf\xd3\x99\x06\xbb\x31\xcf\x73\x9f\xfc\xa3\x0c\xa3\xaa\xef\x1f\xae\xf2\x16\x7c\x79\x4c\xff\x2f\x61\x35\x2f\xfe\xae\x8c\x13\xed\x6a\x9d\xcc\x65\x0b\x03\x1b\xa3\x14\x97\x50\x62\x6b\x3d\x09\x90\x5c\xf1\x73\x39\x27\x0d\xef\xad\x4d\xff\x77\xd8\x1f\xd3\x45\xc5\xb3\xf3\xbb\x16\xf9\xb0\x08\x88\x77\x18\x7e\x97\xf6\xc5\x35\xce\x3e\x96\x51\xb3\xd3\x1c\x5f\xa7\x88\xeb\xe2\x5c\x92\x3f\x01\x00\x00\xff\xff\x83\xdf\xa1\xe1\x2c\x05\x00\x00"

func mint_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_mint_nftCdc,
		"mint_nft.cdc",
	)
}

func mint_nftCdc() (*asset, error) {
	bytes, err := mint_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mint_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb3, 0x3a, 0x5b, 0x9, 0xa5, 0x6e, 0xc6, 0x70, 0xa4, 0xc, 0x88, 0xda, 0x5d, 0xd4, 0xbe, 0xb9, 0x5b, 0xd8, 0xb9, 0xee, 0x91, 0x7b, 0x4a, 0xb2, 0x6, 0xf5, 0x21, 0xc4, 0x89, 0xb0, 0xcf, 0xdf}}
	return a, nil
}

var _scriptsBorrow_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x6b\xe3\x30\x10\x85\xef\xfe\x15\x0f\x1f\x16\x1b\x16\xfb\xb2\xec\x21\x34\x0d\x69\x68\xa0\x97\x10\x8a\x7b\x2e\xf2\x78\x9c\x0c\x95\x25\x21\xcb\xb4\x25\xe4\xbf\x97\xc4\xb1\x9d\xa6\x54\x27\x09\xbd\x6f\xa4\xf7\x49\xe3\xac\x0f\xd8\x58\xb3\xee\xcc\x4e\x4a\xcd\x85\x7d\x63\x83\xda\xdb\x06\x71\x96\xe5\x59\x96\x93\x35\xc1\x2b\x0a\x6d\x7e\x1b\xcb\xa8\xa2\x38\xba\xcc\x78\xfc\x50\x8d\xd3\xbc\x59\x17\xbf\xd0\x53\xa0\xe7\xa2\x3c\x47\xb1\x97\x16\x2d\x79\x71\x01\xa5\xf5\xde\xbe\xb7\x50\x06\xe3\x10\x05\xb2\x5a\x33\x05\xb1\x26\x72\x5d\x89\xba\x33\x68\x94\x98\x44\x55\x95\xe7\xb6\x9d\x61\xd9\x6f\xfe\x42\xaa\x19\x5e\x9e\x4c\xf8\xff\x2f\xc5\x21\x02\x00\xcd\x01\x8a\xc8\x76\x26\x60\x8e\x1d\x87\x65\x7f\x18\xe0\x34\x1a\x63\xd3\x33\xcf\x5c\x63\x3e\x60\xe7\xfb\xd3\xca\x76\x1c\x56\xca\xa9\x52\xb4\x84\xcf\xe4\xaa\xcb\x6a\x24\xb7\x5d\xa9\x85\xb6\x2a\xec\xd3\x89\xeb\x5b\xdd\xfd\x39\xfc\x90\x77\x0b\x1e\xef\x93\x89\x5b\x2c\xe0\x94\x11\x4a\xe2\x95\xed\x74\x05\x63\x07\x41\xa0\xf1\x1b\xbd\x23\x77\xa6\xaf\x1a\xc4\x97\x5e\x79\x8e\x87\x1e\x51\xf0\x5c\xb3\x67\x43\x8c\x60\xa1\xd0\x3a\x26\xa9\x85\xce\xa6\xc5\x20\xec\xf9\xda\xf4\x60\xe5\x15\xf3\xef\x66\x2e\x75\x36\xeb\x22\x39\xe9\x96\x2a\x8d\x8e\xd1\x57\x00\x00\x00\xff\xff\x42\x5e\x85\x9f\x45\x02\x00\x00"

func scriptsBorrow_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBorrow_nftCdc,
		"scripts/borrow_nft.cdc",
	)
}

func scriptsBorrow_nftCdc() (*asset, error) {
	bytes, err := scriptsBorrow_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/borrow_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6e, 0x35, 0x85, 0x35, 0xa5, 0x35, 0x7c, 0x4e, 0xec, 0x2f, 0xd1, 0x85, 0x20, 0x1, 0xe5, 0x41, 0x36, 0x46, 0x72, 0xa9, 0x90, 0x9b, 0x8e, 0xae, 0x8b, 0x7, 0x11, 0x8c, 0xbe, 0x9f, 0xe5, 0xe6}}
	return a, nil
}

var _scriptsGet_collection_lengthCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\xc5\x87\x62\x5f\xe4\x7b\x68\x1a\x82\xdb\x40\x2e\x21\x94\xfc\x80\x2c\xcb\x8e\xa8\xbc\x2b\xe4\x15\x6d\x09\xf9\xf7\x12\xab\xb1\xd3\x94\xe8\x20\x24\x76\xde\xb0\x33\xb6\xf7\x14\x18\x76\x84\x9b\x88\x9d\xad\x9d\x39\xd0\x87\x41\x68\x03\xf5\x90\x49\x59\x4a\x59\x6a\x42\x0e\x4a\xf3\x50\xde\xcb\xa4\x6e\x74\x26\x7e\x3d\xde\xbe\x54\xef\x9d\xd9\x6d\x0e\x0f\xe8\x59\x90\x38\xe1\x63\x0d\x6d\x44\xe8\x95\xc5\x5c\x35\x4d\x30\xc3\xb0\x80\x75\x7a\x14\x0b\xd8\x22\xc3\x49\x00\x00\x38\xc3\xa0\xb4\xa6\x88\x0c\x4b\xe8\x0c\xaf\xd3\xe7\x4a\x15\x62\x92\x69\x72\xce\x68\xb6\x84\xef\xa6\x85\xe5\x15\x1b\xe7\x97\x23\x3b\xc3\x95\xf2\xaa\xb6\xce\xf2\x77\x7e\xb3\x54\x35\x91\xfb\x58\x3b\xab\xf7\x8a\x8f\xc5\xcc\xd5\x14\x02\x7d\x3e\x3f\x9d\xfe\xb5\x70\x0f\x9e\x5f\xf2\x99\x5b\xad\xc0\x2b\xb4\x3a\xcf\x2a\x8a\xae\x01\x24\x86\x64\x05\x7a\x5a\x23\x35\xe6\x47\xfa\x26\x41\x96\x6c\xc6\x2b\x18\x8e\x01\xff\xc6\xbb\x64\xd9\xbe\x0e\x79\x21\x9d\xc1\x8e\x8f\xe2\x2c\x7e\x02\x00\x00\xff\xff\xaa\x92\x4e\x73\xd1\x01\x00\x00"

func scriptsGet_collection_lengthCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_collection_lengthCdc,
		"scripts/get_collection_length.cdc",
	)
}

func scriptsGet_collection_lengthCdc() (*asset, error) {
	bytes, err := scriptsGet_collection_lengthCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_collection_length.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x89, 0xcd, 0xef, 0x5d, 0x12, 0xb7, 0x51, 0x37, 0xc7, 0x80, 0x78, 0xf2, 0x33, 0x27, 0xa0, 0x48, 0xa5, 0xb7, 0x52, 0x63, 0x43, 0xcd, 0x14, 0xbe, 0x1b, 0x93, 0xd9, 0x22, 0x5f, 0x57, 0x8f}}
	return a, nil
}

var _scriptsGet_nft_metadataCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\x4d\x6f\xdb\x3c\x0c\xc7\xef\xfe\x14\x4c\x0e\x0f\x6c\xe0\x81\x7d\x19\x76\x08\xea\x16\x45\xf7\x82\x1d\x36\x14\x58\xb7\xbb\x2c\xd3\x09\x01\x5b\x32\x24\xba\x59\x50\xe4\xbb\x0f\x92\x6c\xcb\x4e\xbd\x5c\x22\x93\xfc\x8b\x3f\x92\x22\x75\xbd\x36\x0c\x9f\xff\x88\xae\x6f\xf1\xc7\x97\x17\x68\x8c\xee\x60\x9f\xe7\x45\x9e\x17\x52\x2b\x36\x42\xb2\x2d\x62\x40\x2e\x6b\xb9\x4f\x46\xdd\x77\x64\x51\x0b\x16\xbf\x09\xcf\x76\x92\x16\x2b\x6b\x88\x4f\xfa\xa1\x02\xcb\x66\x90\x0c\x2e\xcb\x5b\x02\x00\xe0\x8c\x2d\x32\x28\xd1\xe1\x01\x7e\xb2\x21\x75\x5c\x39\x6a\xb4\xd2\x50\xcf\xa4\xd5\xa6\x9f\x4f\x43\x57\x29\x41\xed\xa6\x57\x9f\x15\x9a\x03\x3c\xd6\xb5\x41\x6b\xd7\xc2\x4b\x1f\x33\x7a\x0f\x29\xe2\xd4\x9f\xdc\x6f\x49\xf4\xff\x6c\xdd\xc0\x89\xce\x77\x2c\xd1\xb5\x06\x89\x76\xd5\xf0\xcb\x02\x24\x38\xb2\xb1\x39\xee\x67\xb1\x6d\x72\xc7\x02\xa5\x47\x5a\x3b\x16\x38\x50\x2e\xe1\xd6\x61\x33\x18\x94\x11\x72\x1d\xe2\x01\xa1\x0c\xa0\x37\xea\x4b\xef\xb3\x07\x54\xef\xbb\x26\xd7\x30\xd0\x66\x50\xd0\x09\x52\xa9\x08\x95\xc5\x12\x81\xea\x03\xfc\xfa\xa6\xf8\xe3\x87\xec\xb0\x98\xb8\xeb\xbd\x90\x52\x0f\x8a\xa1\x84\x23\xf2\x63\xf8\x98\x6e\xc8\x92\x39\x4c\xea\xb6\x45\x39\x16\x37\x6a\x66\xb4\xfc\x88\xfc\x24\x7a\x51\x51\x4b\x7c\x49\x17\xcf\xf3\x69\x96\x3d\x0f\x55\x4b\xf2\x59\xf0\x29\x8b\xba\x4a\x1b\xa3\xcf\x77\xff\xbd\x2d\x24\xf1\x78\x2b\xbe\xde\xa7\x51\xfb\xf0\x00\xbd\x50\x24\xd3\xfd\x93\x1e\xda\x1a\x94\x66\x08\xd7\x81\x00\x83\x0d\x1a\x54\x12\x81\x35\xf0\x09\x17\xf8\xfb\x45\x51\xaa\x71\x75\x47\xdf\xc8\x13\x09\x52\xd7\x38\xaa\xb3\x5d\xd0\x14\x05\x7c\xf5\xef\x1c\xa1\x12\x96\x24\xd4\x64\xfb\x56\x5c\x80\x54\xa3\x4d\x27\x7c\x7b\x1a\x6d\x80\x4f\x64\x5d\x9f\xe7\x4c\xaf\x84\xe7\x30\xb7\xdc\xa0\xd5\xed\x2b\xba\x75\x4c\xdd\x10\xef\xd6\x0b\xfa\x29\x5c\x79\x9f\x66\x53\x56\xbf\x7a\x63\xa2\x32\xdc\x24\xec\x0e\x36\x65\x5e\x30\xab\xd6\x2f\x7d\xcc\xef\x8d\xbb\x7c\x1c\xf1\x6e\xd9\x8c\x97\xf9\x75\xb9\x89\xba\xaf\x74\xec\x96\x41\x1e\x8c\x72\x25\xdd\x6e\xe5\x48\xe6\xf7\xe2\x1f\xbb\x39\x85\x2c\x8c\x9b\x8b\x3a\xc5\xcd\xa6\x7c\x30\x94\x66\xef\x36\xd7\xff\x6d\xec\xed\x78\xc8\xa9\x46\xc5\xd4\xd0\x14\x94\x25\xd7\xe4\x6f\x00\x00\x00\xff\xff\xd1\x3a\xad\x04\x5b\x05\x00\x00"

func scriptsGet_nft_metadataCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_nft_metadataCdc,
		"scripts/get_nft_metadata.cdc",
	)
}

func scriptsGet_nft_metadataCdc() (*asset, error) {
	bytes, err := scriptsGet_nft_metadataCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_nft_metadata.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0xcf, 0x9e, 0x96, 0xd5, 0x9d, 0x15, 0x25, 0x93, 0x69, 0x9d, 0xf5, 0x74, 0xb4, 0xaf, 0x88, 0xa, 0x5d, 0x5b, 0xef, 0xbf, 0x74, 0xfd, 0xc6, 0x62, 0x9c, 0x25, 0xc4, 0x6a, 0x8, 0xd9, 0x11}}
	return a, nil
}

var _scriptsGet_total_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x31\x0a\x42\x31\x0c\x06\xe0\x3d\xa7\xf8\x79\x93\x2e\xe9\x22\x0e\xee\x0a\x2e\x2e\xea\x01\x6a\x7d\x0f\x0a\x6d\x1a\x62\x02\x8a\x78\x77\x47\xdd\x3f\xbe\xda\x75\x98\x63\xff\xcc\x5d\xdb\x7c\x3a\x5c\xb0\xd8\xe8\x98\x98\x13\x73\x2a\x43\xdc\x72\xf1\x47\xfa\x01\x2e\xf7\x32\x11\x69\xdc\xb0\x84\xa0\xe7\x2a\xab\xf5\x0e\xd7\xa3\xf8\x76\x83\x37\x01\x80\xcd\x1e\x26\x7f\x2b\xfb\xf0\xdc\xce\xa1\xda\x5e\xf4\xa1\x6f\x00\x00\x00\xff\xff\xab\xdd\xb2\x0f\x76\x00\x00\x00"

func scriptsGet_total_supplyCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_total_supplyCdc,
		"scripts/get_total_supply.cdc",
	)
}

func scriptsGet_total_supplyCdc() (*asset, error) {
	bytes, err := scriptsGet_total_supplyCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_total_supply.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x9, 0x60, 0xa2, 0xa5, 0x58, 0x7b, 0xb8, 0xa2, 0x87, 0x3a, 0x50, 0x8b, 0x97, 0x82, 0xd3, 0xf7, 0x78, 0xfa, 0x17, 0x8a, 0xda, 0xc8, 0x54, 0x76, 0x3b, 0xe3, 0x9c, 0x92, 0x0, 0x29, 0x87}}
	return a, nil
}

var _setup_accountCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xcd\x6e\xda\x40\x10\xbe\xfb\x29\xbe\xe6\x10\x19\x89\xe0\x7b\x44\x23\x45\x51\x38\xa2\xa8\xe5\x05\x86\x65\x8c\x57\x59\x76\xad\xd9\x71\x28\x42\xbc\x7b\xb5\x36\xb5\x0d\xb5\xd4\xee\xc9\x9a\xfd\x7e\x67\x6d\x0f\x75\x10\xc5\x3a\xf8\x55\xe3\xf7\x76\xeb\x78\x13\x3e\xd9\xa3\x94\x70\xc0\xc3\x62\x51\x98\xe0\x55\xc8\x68\x2c\xee\x31\x0b\xb3\x33\x0f\xd9\x55\xe0\xfd\x17\x1d\x6a\xc7\xeb\xd5\x66\x8a\x3a\xdc\x76\xa4\xac\x28\xb0\xa9\x6c\x84\x0a\xf9\x48\x46\x6d\xf0\xb0\x11\xc7\x8a\x14\xe4\x41\xc6\x84\xc6\x2b\x8e\xa1\x71\x3b\x48\xe3\x13\x41\x03\x22\x2b\xac\x46\x76\x25\x9a\x3a\x0d\x84\x0d\xdb\x2f\xc6\x7a\xb5\x89\x59\x36\x56\x3b\x67\x19\x00\xd4\xc2\x35\x09\xe7\xd1\xee\x3d\xcb\x33\x5e\x1b\xad\x5e\x3b\xf5\x19\xce\x2d\x24\x9d\xa2\xc0\x0f\xd6\x46\x3c\x98\xc4\x9d\x60\x4b\x68\xc5\x7d\x0e\x72\xc2\xb4\x3b\xa1\xa2\x08\x82\x09\xce\x71\xeb\xd2\xf3\x6d\x89\xce\x61\xb1\x0d\x22\xe1\xb8\x7c\x1c\x55\x7e\xeb\xf1\x2f\x79\x5a\xce\x33\x26\x2f\x7f\x6a\x10\xda\xf3\x07\x69\x35\xc3\xb7\xef\xf0\xd6\x8d\x12\xa6\x23\x6d\xc4\x7e\x74\xc9\xc6\xf9\xdf\x84\x49\x19\x04\xcf\x47\xf0\xa1\xd6\xd3\x54\x50\xc7\x3a\x1a\x63\xf9\x34\xce\x62\x5a\x89\xf7\xc4\x1d\x62\xe5\xb3\x1b\x9b\x48\x5f\x0c\xab\x69\xf9\xa3\x0d\xf5\x88\xeb\x16\x12\x2a\x5f\x3e\x0d\x4e\x73\x68\xf8\x8f\xde\x37\x56\xe6\x4f\xa3\xba\xd9\x3a\x6b\x60\xa8\xa6\xad\x75\x56\x4f\x28\x83\xb4\xf6\x13\x0d\xaf\x09\x9c\xf5\x9f\xcb\xc7\xf3\x5f\x3f\xed\xe0\xfb\xd1\xaa\xce\xc7\xa1\x86\xcf\x7b\xd8\xe5\x25\xbf\x79\x8a\xc9\x26\x1d\x34\x15\x99\xdf\x80\x95\x64\xcf\xfa\xef\xfa\x3d\x69\x96\x75\xef\x7b\xc9\x7e\x07\x00\x00\xff\xff\x73\xab\x6a\xaa\xa0\x03\x00\x00"

func setup_accountCdcBytes() ([]byte, error) {
	return bindataRead(
		_setup_accountCdc,
		"setup_account.cdc",
	)
}

func setup_accountCdc() (*asset, error) {
	bytes, err := setup_accountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setup_account.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf3, 0xd4, 0x97, 0x91, 0xe4, 0x73, 0xc5, 0x59, 0x21, 0xf4, 0x4e, 0x1d, 0xff, 0xc3, 0x43, 0xef, 0x32, 0xd2, 0x85, 0x28, 0xe3, 0x57, 0x89, 0x3e, 0x83, 0x20, 0x7, 0x87, 0x51, 0x45, 0x7c, 0x74}}
	return a, nil
}

var _transfer_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xc1\x6e\x9b\x40\x10\x86\xef\x3c\xc5\xc8\x87\x16\x4b\x0d\x5c\xaa\x1e\x90\x93\xc8\x72\x1a\x29\x97\x28\x6a\xdd\x07\x58\x96\x01\xb6\xc5\x33\x68\x76\xa8\x5b\x45\x7e\xf7\x0a\x83\xc1\xc4\xc4\x97\x72\x42\x30\xff\xff\xcf\x7c\xbb\xe3\x76\x35\x8b\xc2\x33\xd3\x63\x43\x85\x4b\x2b\xdc\xf2\x2f\x24\xc8\x85\x77\xb0\x88\xa2\xd8\x32\xa9\x18\xab\x3e\x7e\x5b\x13\xd9\xcc\x2e\x82\xde\xe0\xeb\x1f\xb3\xab\x2b\x7c\x7e\xdc\xce\x49\xc7\xbf\x9d\x28\x88\x63\xd8\x96\xce\x83\x8a\x21\x6f\xac\x3a\x26\x70\x1e\x72\x96\xee\x53\x8e\x22\x8e\x0a\x30\x94\xc1\xc9\xb3\x15\x31\x21\x18\x6b\xb9\x21\x05\x65\x30\xc4\x5a\xa2\x04\xc1\x99\x4f\x28\x68\x5d\xed\x90\x34\x81\x75\x96\x09\x7a\xff\x09\xf6\x4e\xcb\x4c\xcc\xfe\xe9\x21\x81\x1f\x4f\xa4\x5f\x3e\x2f\xe1\x35\x08\x00\x00\x6a\xc1\xda\x08\x86\xde\x15\x84\x92\xc0\xba\xd1\x72\xdd\x45\xb4\x35\xd0\x3f\x71\x0c\x05\x2a\x68\x89\x30\x04\x78\xa8\x9b\xb4\x72\x76\x68\x89\xd3\x9f\x68\x75\xd0\x54\xa8\x63\x31\xdc\xb6\x06\xbd\xf3\xd8\xe4\x32\x38\x8f\x48\x59\x84\xf7\x60\x40\x30\x47\x41\xb2\xd8\x8e\xd9\x86\x76\xed\x7d\xf4\x47\x1c\x96\xab\x0a\x8f\xd3\x4e\xb2\xc6\xcf\xdf\x30\x87\xdb\x5e\x33\x94\xb4\x4f\xd4\x25\xac\x3e\x9c\x1d\xc9\x66\x90\xdd\x85\x2d\xe8\x04\x66\x7f\x7e\x57\x16\x53\xe0\x8b\xd1\x72\x39\xf1\xbc\xbf\x87\xda\x90\xb3\xe1\x62\xc3\x4d\x95\x01\xb1\x5e\x19\x84\xf7\xdd\x1c\x63\xb3\x8b\x77\x18\xf4\x70\x2f\x1c\x04\x2d\xba\xdf\x28\xfe\x3d\x0e\x19\xd6\xec\x9d\x76\x10\x06\xd2\x53\x0e\x05\xea\xc6\xd4\x26\x75\x95\xd3\xbf\xe1\xec\xc0\x2f\xc7\xfc\xcb\x79\x07\x86\xaf\x17\x2b\xf1\x56\x7c\xb8\x0b\xff\x87\xd5\x69\xd2\x6b\xb8\x4e\x57\xfb\x28\x18\xf6\x6f\x9e\xf4\x84\x12\xe5\x0a\xab\x9b\xe9\xa5\x89\x4e\x6e\xe1\xf9\xc6\x8c\xef\xd3\xe8\x87\x8e\xf3\x90\xec\x68\xba\x1f\xf3\xd9\xe3\xe9\x44\xfd\x6b\xa8\x2d\xbc\x04\x56\x37\x94\x6b\xc7\xeb\x10\x1c\xfe\x05\x00\x00\xff\xff\xe4\xc8\xf9\xb9\x9b\x04\x00\x00"

func transfer_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_nftCdc,
		"transfer_nft.cdc",
	)
}

func transfer_nftCdc() (*asset, error) {
	bytes, err := transfer_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x83, 0xaa, 0xe1, 0x14, 0x3d, 0xe6, 0x1f, 0xed, 0xd7, 0x9b, 0xa1, 0x31, 0xc3, 0xc6, 0x3a, 0xac, 0x67, 0x61, 0xc5, 0xdc, 0x5a, 0xf7, 0x8c, 0x67, 0x8d, 0x26, 0xc0, 0x51, 0x22, 0x8c, 0xb1, 0xf4}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"destroy_nft.cdc":                   destroy_nftCdc,
	"mint_nft.cdc":                      mint_nftCdc,
	"scripts/borrow_nft.cdc":            scriptsBorrow_nftCdc,
	"scripts/get_collection_length.cdc": scriptsGet_collection_lengthCdc,
	"scripts/get_nft_metadata.cdc":      scriptsGet_nft_metadataCdc,
	"scripts/get_total_supply.cdc":      scriptsGet_total_supplyCdc,
	"setup_account.cdc":                 setup_accountCdc,
	"transfer_nft.cdc":                  transfer_nftCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"destroy_nft.cdc": {destroy_nftCdc, map[string]*bintree{}},
	"mint_nft.cdc": {mint_nftCdc, map[string]*bintree{}},
	"scripts": {nil, map[string]*bintree{
		"borrow_nft.cdc": {scriptsBorrow_nftCdc, map[string]*bintree{}},
		"get_collection_length.cdc": {scriptsGet_collection_lengthCdc, map[string]*bintree{}},
		"get_nft_metadata.cdc": {scriptsGet_nft_metadataCdc, map[string]*bintree{}},
		"get_total_supply.cdc": {scriptsGet_total_supplyCdc, map[string]*bintree{}},
	}},
	"setup_account.cdc": {setup_accountCdc, map[string]*bintree{}},
	"transfer_nft.cdc": {transfer_nftCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}