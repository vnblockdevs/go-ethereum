package parser

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type ParserLog struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}
type ParserReceipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8        `json:"type,omitempty"`
	PostState         []byte       `json:"root"`
	Status            uint64       `json:"status"`
	CumulativeGasUsed uint64       `json:"cumulativeGasUsed" gencodec:"required"`
	Logs              []*ParserLog `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          common.Hash    `json:"transactionHash" gencodec:"required"`
	ContractAddress common.Address `json:"contractAddress"`
	GasUsed         uint64         `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	BlockTime        uint64
	TransactionIndex uint `json:"transactionIndex"`
}

func CopyReceipt(r *ethTypes.Receipt) *ParserReceipt {
	pr := new(ParserReceipt)

	pr.BlockHash = r.BlockHash
	pr.BlockNumber = r.BlockNumber
	pr.TransactionIndex = r.TransactionIndex

	pr.TxHash = r.TxHash
	pr.ContractAddress = r.ContractAddress
	pr.GasUsed = r.GasUsed

	pr.Logs = make([]*ParserLog, 0)

	for _, l := range r.Logs {
		pl := &ParserLog{l.Address, l.Topics, l.Data}
		pr.Logs = append(pr.Logs, pl)
	}

	pr.Status = r.Status
	pr.PostState = r.PostState
	pr.Type = r.Type

	return pr
}
