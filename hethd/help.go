package hethd

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func CommonHashToAddrss(bytes common.Hash) common.Address {
	var b common.Address
	b.SetBytes(bytes[:])
	return b
}

func CommonHashToAddrssString(bytes common.Hash) string {
	return CommonHashToAddrss(bytes).Hex()
}

func CommonHashToAddrssStringLower(bytes common.Hash) string {
	return strings.ToLower(CommonHashToAddrss(bytes).Hex())
}
