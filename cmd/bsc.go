package cmd

import (
	"context"
	"fmt"
	"gabiparser/common"
	"gabiparser/ethclient"
	"gabiparser/hethd"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var logger *log.Entry

type Bscfilter struct {
	contract    []string
	contractAbi abi.ABI
}

func init() {
	logger = common.NewLogger()
}

func NewBscFilter() *Bscfilter {
	rpcUri := viper.GetString("Bsc.Rpcurl")
	cabi, err := abi.JSON(strings.NewReader(ethclient.ERC721ABI))
	if err != nil {
		logger.Errorf("reader err:%+v", err)
	}
	ethclient.InitClient(rpcUri)
	return &Bscfilter{
		contract:    []string{viper.GetString("Bsc.ContractAddress")},
		contractAbi: cabi,
	}
}

func (e *Bscfilter) Scan(startI, endI int64) {
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
				fmt.Println("scan err:", err)
				return
			}

			for _, log := range logs {
				txid := log.TxHash.Hex()
				toaddr := hethd.CommonHashToAddrssStringLower(log.Topics[2])
				fromaddr := hethd.CommonHashToAddrssStringLower(log.Topics[1])
				tokenId, ok := math.ParseBig256(log.Topics[3].Hex())
				logger.Debugf("txid: %v toaddr %v fromaddr: %v", txid, toaddr, fromaddr)

				if !ok {
					logger.Errorf("invalid hex or decimal integer %d", tokenId)
					logger.Fatal("invalid hex or decimal integer %d", tokenId)
				}

				if fromaddr == fullZeroAddress {
					tokeinds = append(tokeinds, tokenId.String())
				}
				fmt.Println(txid, fromaddr, toaddr, tokenId)
				logger.Info("txid: ", txid, "fromaddr: ", fromaddr, "toaddr: ", toaddr, "tokenid: ", tokenId)
			}
			logger.Debugln(len(tokeinds))
		}
	}
}

var bscCmd = &cobra.Command{
	Use:   "bnb",
	Short: "Bsc log parser",
	Long:  `Bsc log parser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bsc log parser")
		bscfilter := NewBscFilter()
		bscfilter.Scan(startI, endI)
	},
}
