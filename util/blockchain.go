package util

import (
	"strconv"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/sirupsen/logrus"
)

func GetShortIdOfHash(hash string) uint64 {
	// first 8 bytes of hex string with 0x prefixed
	id, err := strconv.ParseUint(hash[2:18], 16, 64)
	if err != nil {
		logrus.WithError(err).WithField("hash", hash).Fatalf("Failed convert hash to short id")
	}

	return id
}

func GetSummaryOfBlock(block *types.Block) *types.BlockSummary {
	summary := types.BlockSummary{
		BlockHeader:  block.BlockHeader,
		Transactions: make([]types.Hash, 0, len(block.Transactions)),
	}

	for _, tx := range block.Transactions {
		summary.Transactions = append(summary.Transactions, tx.Hash)
	}

	return &summary
}

// StripLogExtraFields strips extra unnecessary fields from logs to comply with fullnode rpc
func StripLogExtraFieldsForRPC(logs []types.Log) {
	for i := 0; i < len(logs); i++ {
		log := &logs[i]

		log.BlockHash, log.EpochNumber = nil, nil
		log.TransactionHash, log.TransactionIndex = nil, nil
		log.LogIndex, log.TransactionLogIndex = nil, nil
	}
}
