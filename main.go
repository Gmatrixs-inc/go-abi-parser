package main

import (
	"context"
	"fmt"
	"gabiparser/ethclient"
	"gabiparser/hethd"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"strings"
)

func main() {
	rpcUri := "https://mainnet.infura.io/v3/ee07e33cb6414781a72deaf3b303ca3b"
	ethclient.InitClient(rpcUri)

	configTokenRowAddresses := []string{"0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"}
	contractAbi, err := abi.JSON(strings.NewReader(ethclient.ERC721ABI))
	if err != nil {
		fmt.Println(err)
	}

	startI := int64(1)
	endI := int64(1000)
	var tokeinds []string
	for i := startI; i < endI; i++ {
		if len(configTokenRowAddresses) > 0 {
			logs, err := ethclient.RpcFilterLogs(
				context.Background(),
				i,
				i,
				configTokenRowAddresses,
				contractAbi.Events["Transfer"],
			)
			if err != nil {
				return
			}

			for _, log := range logs {
				txid := log.TxHash.Hex()
				toaddr := hethd.CommonHashToAddrssStringLower(log.Topics[2])
				fromaddr := hethd.CommonHashToAddrssStringLower(log.Topics[1])
				tokenId, ok := math.ParseBig256(log.Topics[3].Hex())
				if !ok {
					fmt.Printf("invalid hex or decimal integer %d", tokenId)
				}

				if fromaddr == "0x0000000000000000000000000000000000000000" {
					tokeinds = append(tokeinds, tokenId.String())
				}
				fmt.Println(txid, fromaddr, toaddr, tokenId)
			}
			fmt.Println(len(tokeinds))
		}
	}
}
