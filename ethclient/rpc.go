package ethclient

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/moremorefun/mcommon"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
)

var client *Client
var networkID int64

func InitClient(uri string) {
	var err error
	client, err = Dial(uri)
	if err != nil {
		mcommon.Log.Fatalf("eth client dial error: [%T] %s", err, err.Error())
	}
}

func RpcBlockNumber(ctx context.Context) (int64, error) {
	blockNum, err := client.BlockNumber(ctx)
	if nil != err {
		return 0, err
	}
	return int64(blockNum), nil
}

func RpcBlockByNum(ctx context.Context, blockNum int64) (*types.Block, error) {
	resp, err := client.BlockByNumber(ctx, big.NewInt(blockNum))
	if nil != err {
		return nil, err
	}
	return resp, nil
}

func RpcNonceAt(ctx context.Context, address string) (int64, error) {
	count, err := client.NonceAt(
		ctx,
		common.HexToAddress(address),
		nil,
	)
	if nil != err {
		return 0, err
	}
	return int64(count), nil
}

func RpcNetworkID(ctx context.Context) (int64, error) {
	if networkID != 0 {
		return networkID, nil
	}
	resp, err := client.NetworkID(ctx)
	if nil != err {
		return 0, err
	}
	networkID = resp.Int64()
	return resp.Int64(), nil
}

func RpcSendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := client.SendTransaction(
		ctx,
		tx,
	)
	if nil != err {
		return err
	}
	return nil
}

func RpcTransactionByHash(ctx context.Context, txHashStr string) (*types.Transaction, error) {
	txHash := common.HexToHash(txHashStr)
	tx, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}
	if isPending {
		return nil, nil
	}
	return tx, nil
}

func RpcTransactionReceipt(ctx context.Context, txHashStr string) (*types.Receipt, error) {
	txHash := common.HexToHash(txHashStr)
	tx, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RpcBalanceAt(ctx context.Context, address string) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if nil != err {
		return nil, err
	}
	return balance, nil
}

func RpcFilterLogs(ctx context.Context, startBlock int64, endBlock int64, contractAddresses []string, event abi.Event) ([]types.Log, error) {
	var warpAddresses []common.Address
	for _, contractAddress := range contractAddresses {
		warpAddresses = append(warpAddresses, common.HexToAddress(contractAddress))
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(startBlock),
		ToBlock:   big.NewInt(endBlock),
		Addresses: warpAddresses,
		Topics: [][]common.Hash{
			{event.ID},
		},
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func RpcTokenBalance(ctx context.Context, tokenAddress string, address string) (*big.Int, error) {
	tokenAddressHash := common.HexToAddress(tokenAddress)

	contractAbi, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return nil, err
	}
	input, err := contractAbi.Pack(
		"balanceOf",
		common.HexToAddress(address),
	)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		From:  common.HexToAddress(address),
		To:    &tokenAddressHash,
		Value: nil,
		Data:  input,
	}
	out, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}
	res, err := contractAbi.Unpack("balanceOf", out)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("error call res")
	}
	out0, ok := res[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("error call res")
	}
	return out0, nil
}
