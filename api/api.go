package api

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	C "github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CkbApi interface {
	C.Client
	mercury.Client
}

type DefaultCkbApi struct {
	ckb     C.Client
	mercury mercury.Client
}

func (cli *DefaultCkbApi) GetTransactionProof(ctx context.Context, txHashes []string, blockHash *types.Hash) (*types.TransactionProof, error) {
	return cli.ckb.GetTransactionProof(ctx, txHashes, blockHash)
}

func (cli *DefaultCkbApi) VerifyTransactionProof(ctx context.Context, proof *types.TransactionProof) ([]*types.Hash, error) {
	return cli.ckb.VerifyTransactionProof(ctx, proof)
}

func (cli *DefaultCkbApi) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	return cli.mercury.GetBalance(payload)
}

func (cli *DefaultCkbApi) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildTransferTransaction(payload)
}

func (cli *DefaultCkbApi) BuildAssetAccountCreationTransaction(payload *model.CreateAssetAccountPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildAssetAccountCreationTransaction(payload)
}

func (cli *DefaultCkbApi) BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildAssetCollectionTransaction(payload)
}

func (cli *DefaultCkbApi) RegisterAddresses(normalAddresses []string) ([]string, error) {
	return cli.mercury.RegisterAddresses(normalAddresses)
}

func (cli *DefaultCkbApi) GetGenericTransaction(txHash string) (*resp.GetGenericTransactionResponse, error) {
	return cli.mercury.GetGenericTransaction(txHash)
}

func (cli *DefaultCkbApi) GetGenericBlock(payload *model.GetGenericBlockPayload) (*resp.GenericBlockResponse, error) {
	return cli.mercury.GetGenericBlock(payload)
}

func (cli *DefaultCkbApi) QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error) {
	return cli.mercury.QueryGenericTransactions(payload)
}

func (cli *DefaultCkbApi) GetTipBlockNumber(ctx context.Context) (uint64, error) {
	return cli.ckb.GetTipBlockNumber(ctx)
}

func (cli *DefaultCkbApi) GetTipHeader(ctx context.Context) (*types.Header, error) {
	return cli.ckb.GetTipHeader(ctx)
}

func (cli *DefaultCkbApi) GetCurrentEpoch(ctx context.Context) (*types.Epoch, error) {
	return cli.ckb.GetCurrentEpoch(ctx)
}

func (cli *DefaultCkbApi) GetEpochByNumber(ctx context.Context, number uint64) (*types.Epoch, error) {
	return cli.ckb.GetEpochByNumber(ctx, number)
}

func (cli *DefaultCkbApi) GetBlockHash(ctx context.Context, number uint64) (*types.Hash, error) {
	return cli.ckb.GetBlockHash(ctx, number)
}

func (cli *DefaultCkbApi) GetBlock(ctx context.Context, hash types.Hash) (*types.Block, error) {
	return cli.ckb.GetBlock(ctx, hash)
}

func (cli *DefaultCkbApi) GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error) {
	return cli.ckb.GetHeader(ctx, hash)
}

func (cli *DefaultCkbApi) GetHeaderByNumber(ctx context.Context, number uint64) (*types.Header, error) {
	return cli.ckb.GetHeaderByNumber(ctx, number)
}

func (cli *DefaultCkbApi) GetLiveCell(ctx context.Context, outPoint *types.OutPoint, withData bool) (*types.CellWithStatus, error) {
	return cli.ckb.GetLiveCell(ctx, outPoint, withData)
}

func (cli *DefaultCkbApi) GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error) {
	return cli.ckb.GetTransaction(ctx, hash)
}

func (cli *DefaultCkbApi) GetBlockEconomicState(ctx context.Context, hash types.Hash) (*types.BlockEconomicState, error) {
	return cli.ckb.GetBlockEconomicState(ctx, hash)
}

func (cli *DefaultCkbApi) GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	return cli.ckb.GetBlockByNumber(ctx, number)
}

func (cli *DefaultCkbApi) GetForkBlock(ctx context.Context, blockHash types.Hash) (*types.Block, error) {
	return cli.ckb.GetForkBlock(ctx, blockHash)
}

func (cli *DefaultCkbApi) GetConsensus(ctx context.Context) (*types.Consensus, error) {
	return cli.ckb.GetConsensus(ctx)
}

func (cli *DefaultCkbApi) GetBlockMedianTime(ctx context.Context, blockHash types.Hash) (uint64, error) {
	return cli.ckb.GetBlockMedianTime(ctx, blockHash)
}

func (cli *DefaultCkbApi) DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error) {
	return cli.ckb.DryRunTransaction(ctx, transaction)
}

func (cli *DefaultCkbApi) CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error) {
	return cli.ckb.CalculateDaoMaximumWithdraw(ctx, point, hash)
}

func (cli *DefaultCkbApi) EstimateFeeRate(ctx context.Context, blocks uint64) (*types.EstimateFeeRateResult, error) {
	return cli.ckb.EstimateFeeRate(ctx, blocks)
}

func (cli *DefaultCkbApi) LocalNodeInfo(ctx context.Context) (*types.Node, error) {
	return cli.ckb.LocalNodeInfo(ctx)
}

func (cli *DefaultCkbApi) GetPeers(ctx context.Context) ([]*types.Node, error) {
	return cli.ckb.GetPeers(ctx)
}

func (cli *DefaultCkbApi) GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error) {
	return cli.ckb.GetBannedAddresses(ctx)
}

func (cli *DefaultCkbApi) ClearBannedAddresses(ctx context.Context) error {
	return cli.ckb.ClearBannedAddresses(ctx)
}

func (cli *DefaultCkbApi) SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error {
	return cli.ckb.SetBan(ctx, address, command, banTime, absolute, reason)
}

func (cli *DefaultCkbApi) SyncState(ctx context.Context) (*types.SyncState, error) {
	return cli.ckb.SyncState(ctx)
}

func (cli *DefaultCkbApi) SetNetworkActive(ctx context.Context, state bool) error {
	return cli.ckb.SetNetworkActive(ctx, state)
}

func (cli *DefaultCkbApi) AddNode(ctx context.Context, peerId, address string) error {
	return cli.ckb.AddNode(ctx, peerId, address)
}

func (cli *DefaultCkbApi) RemoveNode(ctx context.Context, peerId string) error {
	return cli.ckb.RemoveNode(ctx, peerId)
}

func (cli *DefaultCkbApi) PingPeers(ctx context.Context) error {
	return cli.ckb.PingPeers(ctx)
}

func (cli *DefaultCkbApi) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	return cli.ckb.SendTransaction(ctx, tx)
}

func (cli *DefaultCkbApi) SendTransactionNoneValidation(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	return cli.ckb.SendTransactionNoneValidation(ctx, tx)
}

func (cli *DefaultCkbApi) TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error) {
	return cli.ckb.TxPoolInfo(ctx)
}

func (cli *DefaultCkbApi) GetRawTxPool(ctx context.Context) (*types.RawTxPool, error) {
	return cli.ckb.GetRawTxPool(ctx)
}

func (cli *DefaultCkbApi) ClearTxPool(ctx context.Context) error {
	return cli.ckb.ClearTxPool(ctx)
}

func (cli *DefaultCkbApi) GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error) {
	return cli.ckb.GetBlockchainInfo(ctx)
}

func (cli *DefaultCkbApi) BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error {
	return cli.ckb.BatchTransactions(ctx, batch)
}

func (cli *DefaultCkbApi) BatchLiveCells(ctx context.Context, batch []types.BatchLiveCellItem) error {
	return cli.ckb.BatchLiveCells(ctx, batch)
}

func (cli *DefaultCkbApi) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return cli.ckb.CallContext(ctx, result, method, args)
}

func (cli *DefaultCkbApi) GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	return cli.ckb.GetCells(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultCkbApi) GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.Transactions, error) {
	return cli.ckb.GetTransactions(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultCkbApi) GetTip(ctx context.Context) (*indexer.TipHeader, error) {
	return cli.ckb.GetTip(ctx)
}

func (cli *DefaultCkbApi) GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error) {
	return cli.ckb.GetCellsCapacity(ctx, searchKey)
}

func (cli *DefaultCkbApi) Close() {
	cli.ckb.Close()
}
func NewCkbApi(address string) (CkbApi, error) {
	dial, err := rpc.Dial(address)
	if err != nil {
		return nil, err
	}

	indexerClient := indexer.NewClient(dial)
	mercuryClient := mercury.NewClient(dial)
	ckbClient := C.NewClientWithIndexer(dial, indexerClient)

	return &DefaultCkbApi{
		ckb:     ckbClient,
		mercury: mercuryClient}, err
}