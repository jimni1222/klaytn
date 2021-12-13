// Copyright 2021 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/governance"
	"github.com/klaytn/klaytn/networks/rpc"
	"github.com/klaytn/klaytn/node/cn/filters"
)

// EthereumAPI provides an API to access the Klaytn through the `eth` namespace.
// TODO-Klaytn: Removed unused variable
type EthereumAPI struct {
	publicFilterAPI   *filters.PublicFilterAPI
	governanceKlayAPI *governance.GovernanceKlayAPI

	publicKlayAPI            *PublicKlayAPI
	publicBlockChainAPI      *PublicBlockChainAPI
	publicTransactionPoolAPI *PublicTransactionPoolAPI
	publicAccountAPI         *PublicAccountAPI
}

// NewEthereumAPI creates a new ethereum API.
// EthereumAPI operates using Klaytn's API internally without overriding.
// Therefore, it is necessary to use APIs defined in two different packages(cn and api),
// so those apis will be defined through a setter.
func NewEthereumAPI() *EthereumAPI {
	return &EthereumAPI{nil, nil, nil, nil, nil, nil}
}

// SetPublicFilterAPI sets publicFilterAPI
func (api *EthereumAPI) SetPublicFilterAPI(publicFilterAPI *filters.PublicFilterAPI) {
	api.publicFilterAPI = publicFilterAPI
}

// SetGovernanceKlayAPI sets governanceKlayAPI
func (api *EthereumAPI) SetGovernanceKlayAPI(governanceKlayAPI *governance.GovernanceKlayAPI) {
	api.governanceKlayAPI = governanceKlayAPI
}

// SetPublicKlayAPI sets publicKlayAPI
func (api *EthereumAPI) SetPublicKlayAPI(publicKlayAPI *PublicKlayAPI) {
	api.publicKlayAPI = publicKlayAPI
}

// SetPublicBlockChainAPI sets publicBlockChainAPI
func (api *EthereumAPI) SetPublicBlockChainAPI(publicBlockChainAPI *PublicBlockChainAPI) {
	api.publicBlockChainAPI = publicBlockChainAPI
}

// SetPublicTransactionPoolAPI sets publicTransactionPoolAPI
func (api *EthereumAPI) SetPublicTransactionPoolAPI(publicTransactionPoolAPI *PublicTransactionPoolAPI) {
	api.publicTransactionPoolAPI = publicTransactionPoolAPI
}

// SetPublicAccountAPI sets publicAccountAPI
func (api *EthereumAPI) SetPublicAccountAPI(publicAccountAPI *PublicAccountAPI) {
	api.publicAccountAPI = publicAccountAPI
}

// Etherbase is the address that mining rewards will be send to
func (api *EthereumAPI) Etherbase() (common.Address, error) {
	// TODO-Klaytn: Not implemented yet.
	return common.StringToAddress("0x0"), nil
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase)
func (api *EthereumAPI) Coinbase() (common.Address, error) {
	// TODO-Klaytn: Not implemented yet.
	return common.StringToAddress("0x0"), nil
}

// Hashrate returns the POW hashrate
func (api *EthereumAPI) Hashrate() hexutil.Uint64 {
	// TODO-Klaytn: Not implemented yet.
	return 0
}

// Mining returns an indication if this node is currently mining.
func (api *EthereumAPI) Mining() bool {
	// TODO-Klaytn: Not implemented yet.
	return false
}

// GetWork returns a work package for external miner.
//
// The work package consists of 3 strings:
//   result[0] - 32 bytes hex encoded current block header pow-hash
//   result[1] - 32 bytes hex encoded seed hash used for DAG
//   result[2] - 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
//   result[3] - hex encoded block number
func (api *EthereumAPI) GetWork() ([4]string, error) {
	// TODO-Klaytn: Not implemented yet.
	return [4]string{}, nil
}

// A BlockNonce is a 64-bit hash which proves (combined with the
// mix-hash) that a sufficient amount of computation has been carried
// out on a block.
type BlockNonce [8]byte

// SubmitWork can be used by external miner to submit their POW solution.
// It returns an indication if the work was accepted.
// Note either an invalid solution, a stale work a non-existent work will return false.
func (api *EthereumAPI) SubmitWork(nonce BlockNonce, hash, digest common.Hash) bool {
	// TODO-Klaytn: Not implemented yet.
	return false
}

// SubmitHashrate can be used for remote miners to submit their hash rate.
// This enables the node to report the combined hash rate of all miners
// which submit work through this node.
//
// It accepts the miner hash rate and an identifier which must be unique
// between nodes.
func (api *EthereumAPI) SubmitHashrate(rate hexutil.Uint64, id common.Hash) bool {
	// TODO-Klaytn: Not implemented yet.
	return false
}

// GetHashrate returns the current hashrate for local CPU miner and remote miner.
func (api *EthereumAPI) GetHashrate() uint64 {
	// TODO-Klaytn: Not implemented yet.
	return 0
}

// NewPendingTransactionFilter creates a filter that fetches pending transaction hashes
// as transactions enter the pending state.
//
// It is part of the filter package because this filter can be used through the
// `eth_getFilterChanges` polling method that is also used for log filters.
//
// https://eth.wiki/json-rpc/API#eth_newpendingtransactionfilter
func (api *EthereumAPI) NewPendingTransactionFilter() rpc.ID {
	// TODO-Klaytn: Not implemented yet.
	return ""
}

// NewPendingTransactions creates a subscription that is triggered each time a transaction
// enters the transaction pool and was signed from one of the transactions this nodes manages.
func (api *EthereumAPI) NewPendingTransactions(ctx context.Context) (*rpc.Subscription, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
//
// https://eth.wiki/json-rpc/API#eth_newblockfilter
func (api *EthereumAPI) NewBlockFilter() rpc.ID {
	// TODO-Klaytn: Not implemented yet.
	return ""
}

// NewHeads send a notification each time a new (header) block is appended to the chain.
func (api *EthereumAPI) NewHeads(ctx context.Context) (*rpc.Subscription, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// Logs creates a subscription that fires for all new log that match the given filter criteria.
func (api *EthereumAPI) Logs(ctx context.Context, crit FilterCriteria) (*rpc.Subscription, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// FilterCriteria represents a request to create a new filter.
type FilterCriteria filters.FilterCriteria

// NewFilter creates a new filter and returns the filter id. It can be
// used to retrieve logs when the state changes. This method cannot be
// used to fetch logs that are already stored in the state.
//
// Default criteria for the from and to block are "latest".
// Using "latest" as block number will return logs for mined blocks.
// Using "pending" as block number returns logs for not yet mined (pending) blocks.
// In case logs are removed (chain reorg) previously returned logs are returned
// again but with the removed property set to true.
//
// In case "fromBlock" > "toBlock" an error is returned.
//
// https://eth.wiki/json-rpc/API#eth_newfilter
func (api *EthereumAPI) NewFilter(crit FilterCriteria) (rpc.ID, error) {
	// TODO-Klaytn: Not implemented yet.
	return "", nil
}

// GetLogs returns logs matching the given argument that are stored within the state.
//
// https://eth.wiki/json-rpc/API#eth_getlogs
func (api *EthereumAPI) GetLogs(ctx context.Context, crit FilterCriteria) ([]*types.Log, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// UninstallFilter removes the filter with the given filter id.
//
// https://eth.wiki/json-rpc/API#eth_uninstallfilter
func (api *EthereumAPI) UninstallFilter(id rpc.ID) bool {
	// TODO-Klaytn: Not implemented yet.
	return false
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
//
// https://eth.wiki/json-rpc/API#eth_getfilterlogs
func (api *EthereumAPI) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*types.Log, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetFilterChanges returns the logs for the filter with the given id since
// last time it was called. This can be used for polling.
//
// For pending transaction and block filters the result is []common.Hash.
// (pending)Log filters return []Log.
//
// https://eth.wiki/json-rpc/API#eth_getfilterchanges
func (api *EthereumAPI) GetFilterChanges(id rpc.ID) (interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GasPrice returns a suggestion for a gas price for legacy transactions.
func (api *EthereumAPI) GasPrice(ctx context.Context) (*hexutil.Big, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
func (api *EthereumAPI) MaxPriorityFeePerGas(ctx context.Context) (*hexutil.Big, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

type feeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// DecimalOrHex unmarshals a non-negative decimal or hex parameter into a uint64.
type DecimalOrHex uint64

func (api *EthereumAPI) FeeHistory(ctx context.Context, blockCount DecimalOrHex, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*feeHistoryResult, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
// yet received the latest block headers from its pears. In case it is synchronizing:
// - startingBlock: block number this node started to synchronise from
// - currentBlock:  block number this node is currently importing
// - highestBlock:  block number of the highest block header this node has received from peers
// - pulledStates:  number of state entries processed until now
// - knownStates:   number of known state entries that still need to be pulled
func (api *EthereumAPI) Syncing() (interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// ChainId is the EIP-155 replay-protection chain id for the current ethereum chain config.
func (api *EthereumAPI) ChainId() (*hexutil.Big, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// BlockNumber returns the block number of the chain head.
func (api *EthereumAPI) BlockNumber() hexutil.Uint64 {
	// TODO-Klaytn: Not implemented yet.
	return 0
}

// GetBalance returns the amount of wei for the given address in the state of the
// given block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta
// block numbers are also allowed.
func (api *EthereumAPI) GetBalance(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// EthAccountResult structs for GetProof
// AccountResult in go-ethereum has been renamed to EthAccountResult.
// AccountResult is defined in go-ethereum's internal package, so AccountResult is redefined here as EthAccountResult.
type EthAccountResult struct {
	Address      common.Address     `json:"address"`
	AccountProof []string           `json:"accountProof"`
	Balance      *hexutil.Big       `json:"balance"`
	CodeHash     common.Hash        `json:"codeHash"`
	Nonce        hexutil.Uint64     `json:"nonce"`
	StorageHash  common.Hash        `json:"storageHash"`
	StorageProof []EthStorageResult `json:"storageProof"`
}

// StorageResult in go-ethereum has been renamed to EthStorageResult.
// StorageResult is defined in go-ethereum's internal package, so StorageResult is redefined here as EthStorageResult.
type EthStorageResult struct {
	Key   string       `json:"key"`
	Value *hexutil.Big `json:"value"`
	Proof []string     `json:"proof"`
}

// GetProof returns the Merkle-proof for a given account and optionally some storage keys.
func (api *EthereumAPI) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*EthAccountResult, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetHeaderByNumber returns the requested canonical block header.
// * When blockNr is -1 the chain head is returned.
// * When blockNr is -2 the pending chain head is returned.
func (api *EthereumAPI) GetHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (map[string]interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetHeaderByHash returns the requested header by hash.
func (api *EthereumAPI) GetHeaderByHash(ctx context.Context, hash common.Hash) map[string]interface{} {
	// TODO-Klaytn: Not implemented yet.
	return nil
}

// GetBlockByNumber returns the requested canonical block.
// * When blockNr is -1 the chain head is returned.
// * When blockNr is -2 the pending chain head is returned.
// * When fullTx is true all transactions in the block are returned, otherwise
//   only the transaction hash is returned.
func (api *EthereumAPI) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetBlockByHash returns the requested block. When fullTx is true all transactions in the block are returned in full
// detail, otherwise only the transaction hash is returned.
func (api *EthereumAPI) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetUncleByBlockNumberAndIndex returns the uncle block for the given block hash and index. When fullTx is true
// all transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (api *EthereumAPI) GetUncleByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) (map[string]interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetUncleByBlockHashAndIndex returns the uncle block for the given block hash and index. When fullTx is true
// all transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (api *EthereumAPI) GetUncleByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) (map[string]interface{}, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetUncleCountByBlockNumber returns number of uncles in the block for the given block number
func (api *EthereumAPI) GetUncleCountByBlockNumber(ctx context.Context, blockNr rpc.BlockNumber) *hexutil.Uint {
	// TODO-Klaytn: Not implemented yet.
	return nil
}

// GetUncleCountByBlockHash returns number of uncles in the block for the given block hash
func (api *EthereumAPI) GetUncleCountByBlockHash(ctx context.Context, blockHash common.Hash) *hexutil.Uint {
	// TODO-Klaytn: Not implemented yet.
	return nil
}

// GetCode returns the code stored at the given address in the state for the given block number.
func (api *EthereumAPI) GetCode(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// GetStorageAt returns the storage from the state at the given address, key and
// block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta block
// numbers are also allowed.
func (api *EthereumAPI) GetStorageAt(ctx context.Context, address common.Address, key string, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// EthOverrideAccount indicates the overriding fields of account during the execution
// of a message call.
// Note, state and stateDiff can't be specified at the same time. If state is
// set, message execution will only use the data in the given state. Otherwise
// if statDiff is set, all diff will be applied first and then execute the call
// message.
// OverrideAccount in go-ethereum has been renamed to EthOverrideAccount.
// OverrideAccount is defined in go-ethereum's internal package, so OverrideAccount is redefined here as EthOverrideAccount.
type EthOverrideAccount struct {
	Nonce     *hexutil.Uint64              `json:"nonce"`
	Code      *hexutil.Bytes               `json:"code"`
	Balance   **hexutil.Big                `json:"balance"`
	State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff *map[common.Hash]common.Hash `json:"stateDiff"`
}

// EthStateOverride is the collection of overridden accounts.
// StateOverride in go-ethereum has been renamed to EthStateOverride.
// StateOverride is defined in go-ethereum's internal package, so StateOverride is redefined here as EthStateOverride.
type EthStateOverride map[common.Address]EthOverrideAccount

// Call executes the given transaction on the state for the given block number.
//
// Additionally, the caller can specify a batch of contract for fields overriding.
//
// Note, this function doesn't make and changes in the state/blockchain and is
// useful to execute and retrieve values.
func (api *EthereumAPI) Call(ctx context.Context, args EthTransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *EthStateOverride) (hexutil.Bytes, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// EstimateGas returns an estimate of the amount of gas needed to execute the
// given transaction against the current pending block.
func (api *EthereumAPI) EstimateGas(ctx context.Context, args EthTransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	// TODO-Klaytn: Not implemented yet.
	return 0, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block with the given block number.
func (api *EthereumAPI) GetBlockTransactionCountByNumber(ctx context.Context, blockNr rpc.BlockNumber) *hexutil.Uint {
	// TODO-Klaytn: Not implemented yet.
	return nil
}

// GetBlockTransactionCountByHash returns the number of transactions in the block with the given hash.
func (api *EthereumAPI) GetBlockTransactionCountByHash(ctx context.Context, blockHash common.Hash) *hexutil.Uint {
	// TODO-Klaytn: Not implemented yet.
	return nil
}

// accessListResult returns an optional accesslist
// Its the result of the `debug_createAccessList` RPC call.
// It contains an error if the transaction itself failed.
type accessListResult struct {
	Accesslist *AccessList    `json:"accessList"`
	Error      string         `json:"error,omitempty"`
	GasUsed    hexutil.Uint64 `json:"gasUsed"`
}

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"        gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys"    gencodec:"required"`
}

// CreateAccessList creates a EIP-2930 type AccessList for the given transaction.
// Reexec and BlockNrOrHash can be specified to create the accessList on top of a certain state.
func (api *EthereumAPI) CreateAccessList(ctx context.Context, args EthTransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (*accessListResult, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// EthRPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
// RPCTransaction in go-ethereum has been renamed to EthRPCTransaction.
// RPCTransaction is defined in go-ethereum's internal package, so RPCTransaction is redefined here as EthRPCTransaction.
type EthRPCTransaction struct {
	BlockHash        *common.Hash    `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	GasFeeCap        *hexutil.Big    `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex *hexutil.Uint64 `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	Type             hexutil.Uint64  `json:"type"`
	Accesses         *AccessList     `json:"accessList,omitempty"`
	ChainID          *hexutil.Big    `json:"chainId,omitempty"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

// newEthRPCTransactionFromBlockAndIndex creates an EthRPCTransaction from block and index parameters.
func newEthRPCTransactionFromBlockAndIndex(b *types.Block, index uint64) (*EthRPCTransaction, error) {
	txs := b.Transactions()
	if index >= uint64(len(txs)) {
		return nil, errors.New("invalid transaction index")
	}
	return newEthRPCTransaction(txs[index], b.Hash(), b.NumberU64(), index)
}

// newEthRPCTransaction creates an EthRPCTransaction from Klaytn transaction.
func newEthRPCTransaction(tx *types.Transaction, blockHash common.Hash, blockNumber, index uint64) (*EthRPCTransaction, error) {
	// When an unknown transaction is requested through rpc call,
	// nil is returned by Klaytn API, and it is handled.
	if tx == nil {
		return nil, nil
	}

	from := getFrom(tx)

	// If to is nil, it is fills with from.
	to := tx.To()
	if to == nil {
		to = &from
	}

	// If tx is not TxTypeLegacyTransaction, the type is converted to TxTypeLegacyTransaction.
	// TODO-Klaytn: In the case of Ethereum transaction type,
	//  it must be returned as it is without converting the type.
	typeInt := hexutil.Uint64(tx.Type())
	if types.TxType(typeInt) != types.TxTypeLegacyTransaction {
		typeInt = hexutil.Uint64(types.TxTypeLegacyTransaction)
	}

	signature := tx.GetTxInternalData().RawSignatureValues()[0]

	result := &EthRPCTransaction{
		Type:     typeInt,
		From:     from,
		Gas:      hexutil.Uint64(tx.Gas()),
		GasPrice: (*hexutil.Big)(tx.GasPrice()),
		Hash:     tx.Hash(),
		Input:    tx.Data(),
		Nonce:    hexutil.Uint64(tx.Nonce()),
		To:       to,
		Value:    (*hexutil.Big)(tx.Value()),
		V:        (*hexutil.Big)(signature.V),
		R:        (*hexutil.Big)(signature.R),
		S:        (*hexutil.Big)(signature.S),
	}

	if blockHash != (common.Hash{}) {
		result.BlockHash = &blockHash
		result.BlockNumber = (*hexutil.Big)(new(big.Int).SetUint64(blockNumber))
		result.TransactionIndex = (*hexutil.Uint64)(&index)
	}

	// TODO-Klaytn: Have to add additional fields for ethereum transaction types.

	return result, nil
}

// newEthRPCPendingTransaction creates an EthRPCTransaction for pending tx.
func newEthRPCPendingTransaction(tx *types.Transaction) (*EthRPCTransaction, error) {
	return newEthRPCTransaction(tx, common.Hash{}, 0, 0)
}

// GetTransactionByBlockNumberAndIndex returns the transaction for the given block number and index.
func (api *EthereumAPI) GetTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) *EthRPCTransaction {
	block, err := api.publicTransactionPoolAPI.b.BlockByNumber(ctx, blockNr)
	if block != nil && err == nil {
		ethTx, err := newEthRPCTransactionFromBlockAndIndex(block, uint64(index))
		if ethTx == nil || err != nil {
			return nil
		}
		return ethTx
	}
	return nil
}

// GetTransactionByBlockHashAndIndex returns the transaction for the given block hash and index.
func (api *EthereumAPI) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) *EthRPCTransaction {
	block, err := api.publicTransactionPoolAPI.b.BlockByHash(ctx, blockHash)
	if block != nil && err == nil {
		ethTx, err := newEthRPCTransactionFromBlockAndIndex(block, uint64(index))
		if ethTx == nil || err != nil {
			return nil
		}
		return ethTx
	}
	return nil
}

// GetRawTransactionByBlockNumberAndIndex returns the bytes of the transaction for the given block number and index.
func (api *EthereumAPI) GetRawTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) hexutil.Bytes {
	rawTx, err := api.publicTransactionPoolAPI.GetRawTransactionByBlockNumberAndIndex(ctx, blockNr, index)
	if rawTx == nil || err != nil {
		return nil
	}

	return rawTx
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (api *EthereumAPI) GetRawTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) hexutil.Bytes {
	rawTx, err := api.publicTransactionPoolAPI.GetRawTransactionByBlockHashAndIndex(ctx, blockHash, index)
	if rawTx == nil || err != nil {
		return nil
	}

	return rawTx
}

// GetTransactionCount returns the number of transactions the given address has sent for the given block number
func (api *EthereumAPI) GetTransactionCount(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	return api.publicTransactionPoolAPI.GetTransactionCount(ctx, address, blockNrOrHash)
}

// GetTransactionByHash returns the transaction for the given hash
func (api *EthereumAPI) GetTransactionByHash(ctx context.Context, hash common.Hash) (*EthRPCTransaction, error) {
	// Try to return an already finalized transaction
	if tx, blockHash, blockNumber, index := api.publicTransactionPoolAPI.b.ChainDB().ReadTxAndLookupInfo(hash); tx != nil {
		return newEthRPCTransaction(tx, blockHash, blockNumber, index)
	}
	// No finalized transaction, try to retrieve it from the pool
	if tx := api.publicTransactionPoolAPI.b.GetPoolTransaction(hash); tx != nil {
		return newEthRPCPendingTransaction(tx)
	}
	// Transaction unknown, return as such
	return nil, nil
}

// GetRawTransactionByHash returns the bytes of the transaction for the given hash.
func (api *EthereumAPI) GetRawTransactionByHash(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	rawTx, err := api.publicTransactionPoolAPI.GetRawTransactionByHash(ctx, hash)
	if rawTx == nil || err != nil {
		return nil, err
	}

	return rawTx, nil
}

// GetTransactionReceipt returns the transaction receipt for the given transaction hash.
func (api *EthereumAPI) GetTransactionReceipt(ctx context.Context, hash common.Hash) (map[string]interface{}, error) {
	// Formats return Klaytn Transaction Receipt to the Ethereum Transaction Receipt.
	ethTx, err := newEthTransactionReceipt(api.publicTransactionPoolAPI.b.GetTxLookupInfoAndReceipt(ctx, hash))
	if ethTx == nil || err != nil {
		return nil, err
	}
	return ethTx, nil
}

// newEthTransactionReceipt creates a transaction receipt in Ethereum format.
func newEthTransactionReceipt(tx *types.Transaction, blockHash common.Hash, blockNumber, index uint64, receipt *types.Receipt) (map[string]interface{}, error) {
	// When an unknown transaction receipt is requested through rpc call,
	// nil is returned by Klaytn API, and it is handled.
	if tx == nil || receipt == nil {
		return nil, nil
	}

	from := getFrom(tx)

	// If to is nil, it is fills with from.
	to := tx.To()
	if to == nil {
		to = &from
	}

	// If tx is not TxTypeLegacyTransaction, the type is converted to TxTypeLegacyTransaction.
	// TODO-Klaytn: In the case of Ethereum transaction type,
	//  it must be returned as it is without converting the type.
	typeInt := tx.Type()
	if typeInt != types.TxTypeLegacyTransaction {
		typeInt = types.TxTypeLegacyTransaction
	}

	fields := map[string]interface{}{
		"blockHash":         blockHash,
		"blockNumber":       hexutil.Uint64(blockNumber),
		"transactionHash":   tx.Hash(),
		"transactionIndex":  hexutil.Uint64(index),
		"from":              getFrom(tx),
		"to":                to,
		"gasUsed":           hexutil.Uint64(receipt.GasUsed),
		"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
		"contractAddress":   nil,
		"logs":              receipt.Logs,
		"logsBloom":         receipt.Bloom,
		"type":              hexutil.Uint(typeInt),
	}

	fields["effectiveGasPrice"] = tx.GasPrice()

	// Always use the "status" field and Ignore the "root" field.
	fields["status"] = hexutil.Uint(receipt.Status)

	if receipt.Logs == nil {
		fields["logs"] = [][]*types.Log{}
	}
	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if receipt.ContractAddress != (common.Address{}) {
		fields["contractAddress"] = receipt.ContractAddress
	}

	return fields, nil
}

// EthTransactionArgs represents the arguments to construct a new transaction
// or a message call.
// TransactionArgs in go-ethereum has been renamed to EthTransactionArgs.
// TransactionArgs is defined in go-ethereum's internal package, so TransactionArgs is redefined here as EthTransactionArgs.
type EthTransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *AccessList  `json:"accessList,omitempty"`
	ChainID    *hexutil.Big `json:"chainId,omitempty"`
}

// SendTransaction creates a transaction for the given argument, sign it and submit it to the
// transaction pool.
func (api *EthereumAPI) SendTransaction(ctx context.Context, args EthTransactionArgs) (common.Hash, error) {
	// TODO-Klaytn: Not implemented yet.
	return common.HexToHash("0x"), nil
}

// EthSignTransactionResult represents a RLP encoded signed transaction.
// SignTransactionResult in go-ethereum has been renamed to EthSignTransactionResult.
// SignTransactionResult is defined in go-ethereum's internal package, so SignTransactionResult is redefined here as EthSignTransactionResult.
type EthSignTransactionResult struct {
	Raw hexutil.Bytes `json:"raw"`
	Tx  *Transaction  `json:"tx"`
}

// Transaction is an Ethereum transaction.
type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

// TxData is the underlying data of a transaction.
//
// This is implemented by DynamicFeeTx, LegacyTx and AccessListTx.
type TxData interface {
	txType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	chainID() *big.Int
	accessList() AccessList
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	gasTipCap() *big.Int
	gasFeeCap() *big.Int
	value() *big.Int
	nonce() uint64
	to() *common.Address

	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(chainID, v, r, s *big.Int)
}

// FillTransaction fills the defaults (nonce, gas, gasPrice or 1559 fields)
// on a given unsigned transaction, and returns it to the caller for further
// processing (signing + broadcast).
func (api *EthereumAPI) FillTransaction(ctx context.Context, args EthTransactionArgs) (*EthSignTransactionResult, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// SendRawTransaction will add the signed transaction to the transaction pool.
// The sender is responsible for signing the transaction and using the correct nonce.
func (api *EthereumAPI) SendRawTransaction(ctx context.Context, input hexutil.Bytes) (common.Hash, error) {
	return api.publicTransactionPoolAPI.SendRawTransaction(ctx, input)
}

// Sign calculates an ECDSA signature for:
// keccack256("\x19Ethereum Signed Message:\n" + len(message) + message).
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The account associated with addr must be unlocked.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sign
func (api *EthereumAPI) Sign(addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	return api.publicTransactionPoolAPI.Sign(addr, data)
}

// SignTransaction will sign the given transaction with the from account.
// The node needs to have the private key of the account corresponding with
// the given from address and it needs to be unlocked.
func (api *EthereumAPI) SignTransaction(ctx context.Context, args EthTransactionArgs) (*EthSignTransactionResult, error) {
	// TODO-Klaytn: Not implemented yet.
	return nil, nil
}

// PendingTransactions returns the transactions that are in the transaction pool
// and have a from address that is one of the accounts this node manages.
func (api *EthereumAPI) PendingTransactions() ([]*EthRPCTransaction, error) {
	pending, err := api.publicTransactionPoolAPI.b.GetPoolTransactions()
	if err != nil {
		return nil, err
	}
	accounts := getAccountsFromWallets(api.publicTransactionPoolAPI.b.AccountManager().Wallets())
	transactions := make([]*EthRPCTransaction, 0, len(pending))
	for _, tx := range pending {
		from := getFrom(tx)
		if _, exists := accounts[from]; exists {
			ethTx, err := newEthRPCPendingTransaction(tx)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, ethTx)
		}
	}
	return transactions, nil
}

// Resend accepts an existing transaction and a new gas price and limit. It will remove
// the given transaction from the pool and reinsert it with the new gas price and limit.
func (api *EthereumAPI) Resend(ctx context.Context, sendArgs EthTransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error) {
	// TODO-Klaytn: Not implemented yet.
	return common.HexToHash("0x"), nil
}

// Accounts returns the collection of accounts this node manages
func (api *EthereumAPI) Accounts() []common.Address {
	// TODO-Klaytn: Not implemented yet.
	return nil
}
