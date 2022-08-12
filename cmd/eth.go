package cmd

import (
	"context"
	"fmt"
	"gabiparser/ethclient"
	"gabiparser/hethd"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func logFilter(startI, endI int64) {
	rpcUri := viper.GetString("Eth.Rpcurl")
	ethclient.InitClient(rpcUri)

	configTokenRowAddresses := []string{viper.GetString("Eth.ContractAddress")}
	contractAbi, err := abi.JSON(strings.NewReader(ethclient.ERC721ABI))
	if err != nil {
		fmt.Println(err)
	}

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

var ethCmd = &cobra.Command{
	Use:   "eth",
	Short: "Eth log parser",
	Long:  `Eth log parser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eth log parser")

	},
}
