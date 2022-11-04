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

type Ethfilter struct {
	contract    []string
	contractAbi abi.ABI
}

func NewEthfilter() *Ethfilter {
	rpcUri := viper.GetString("Eth.Rpcurl")
	cabi, err := abi.JSON(strings.NewReader(ethclient.ERC721ABI))
	if err != nil {
		fmt.Println(err)
	}
	ethclient.InitClient(rpcUri)
	return &Ethfilter{
		contract:    []string{viper.GetString("Eth.ContractAddress")},
		contractAbi: cabi,
	}
}

func (e *Ethfilter) Scan(startI, endI int64) {
	fullZeroAddress := "0x0000000000000000000000000000000000000000"
	var tokeinds []string
	for i := startI; i < endI; i++ {
		if len(e.contract) > 0 {
			logs, err := ethclient.RpcFilterLogs(
				context.Background(),
				i,
				i,
				e.contract,
				e.contractAbi.Events["Transfer"],
			)
			if err != nil {
				return
			}

			for _, log := range logs {
				txid := log.TxHash.Hex()
				toaddr := hethd.CommonHashToAddrssStringLower(log.Topics[2])
				fromaddr := hethd.CommonHashToAddrssStringLower(log.Topics[1])
				tokenId, ok := math.ParseBig256(log.Topics[3].Hex())
				logger.Debugf("txid: %v toaddr %v fromaddr: %v", txid, toaddr, fromaddr)

				if !ok {
					fmt.Printf("invalid hex or decimal integer %d", tokenId)
					logger.Fatal("invalid hex or decimal integer %d", tokenId)
				}

				if fromaddr == fullZeroAddress {
					tokeinds = append(tokeinds, tokenId.String())
				}
			}
		}
	}
}

var ethCmd = &cobra.Command{
	Use:   "eth",
	Short: "Eth log parser",
	Long:  `Eth log parser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eth log parser")
		ethfilter := NewEthfilter()
		ethfilter.Scan(startI, endI)
	},
}
