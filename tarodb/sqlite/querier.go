// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlite

import (
	"context"
	"database/sql"
)

type Querier interface {
	AllAssets(ctx context.Context) ([]Asset, error)
	AllInternalKeys(ctx context.Context) ([]InternalKey, error)
	AllMintingBatches(ctx context.Context) ([]AllMintingBatchesRow, error)
	AnchorGenesisPoint(ctx context.Context, arg AnchorGenesisPointParams) error
	AnchorPendingAssets(ctx context.Context, arg AnchorPendingAssetsParams) error
	ApplySpendDelta(ctx context.Context, arg ApplySpendDeltaParams) (int32, error)
	AssetsByGenesisPoint(ctx context.Context, prevOut []byte) ([]AssetsByGenesisPointRow, error)
	AssetsInBatch(ctx context.Context, rawKey []byte) ([]AssetsInBatchRow, error)
	BindMintingBatchWithTx(ctx context.Context, arg BindMintingBatchWithTxParams) error
	ConfirmChainAnchorTx(ctx context.Context, arg ConfirmChainAnchorTxParams) error
	ConfirmChainTx(ctx context.Context, arg ConfirmChainTxParams) error
	DeleteAssetWitnesses(ctx context.Context, assetID int32) error
	DeleteManagedUTXO(ctx context.Context, outpoint []byte) error
	DeleteNode(ctx context.Context, arg DeleteNodeParams) (int64, error)
	DeleteSpendProofs(ctx context.Context, transferID int32) error
	FetchAddrByTaprootOutputKey(ctx context.Context, taprootOutputKey []byte) (FetchAddrByTaprootOutputKeyRow, error)
	FetchAddrEvent(ctx context.Context, id int32) (FetchAddrEventRow, error)
	FetchAddrs(ctx context.Context, arg FetchAddrsParams) ([]FetchAddrsRow, error)
	FetchAssetDeltas(ctx context.Context, transferID int32) ([]FetchAssetDeltasRow, error)
	FetchAssetDeltasWithProofs(ctx context.Context, transferID int32) ([]FetchAssetDeltasWithProofsRow, error)
	FetchAssetProof(ctx context.Context, tweakedScriptKey []byte) (FetchAssetProofRow, error)
	FetchAssetProofs(ctx context.Context) ([]FetchAssetProofsRow, error)
	FetchAssetWitnesses(ctx context.Context, assetID sql.NullInt32) ([]FetchAssetWitnessesRow, error)
	FetchAssetsByAnchorTx(ctx context.Context, anchorUtxoID sql.NullInt32) ([]Asset, error)
	// We use a LEFT JOIN here as not every asset has a family key, so this'll
	// generate rows that have NULL values for the faily key fields if an asset
	// doesn't have a family key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	FetchAssetsForBatch(ctx context.Context, rawKey []byte) ([]FetchAssetsForBatchRow, error)
	FetchChainTx(ctx context.Context, txid []byte) (ChainTxn, error)
	FetchChildren(ctx context.Context, arg FetchChildrenParams) ([]FetchChildrenRow, error)
	FetchChildrenSelfJoin(ctx context.Context, arg FetchChildrenSelfJoinParams) ([]FetchChildrenSelfJoinRow, error)
	FetchGenesisByID(ctx context.Context, genAssetID int32) (FetchGenesisByIDRow, error)
	FetchGenesisPointByAnchorTx(ctx context.Context, anchorTxID sql.NullInt32) (GenesisPoint, error)
	FetchManagedUTXO(ctx context.Context, arg FetchManagedUTXOParams) (FetchManagedUTXORow, error)
	FetchMintingBatch(ctx context.Context, rawKey []byte) (FetchMintingBatchRow, error)
	FetchMintingBatchesByInverseState(ctx context.Context, batchState int16) ([]FetchMintingBatchesByInverseStateRow, error)
	FetchMintingBatchesByState(ctx context.Context, batchState int16) ([]FetchMintingBatchesByStateRow, error)
	FetchRootNode(ctx context.Context, namespace string) (MssmtNode, error)
	FetchSeedlingsForBatch(ctx context.Context, rawKey []byte) ([]AssetSeedling, error)
	FetchSpendProofs(ctx context.Context, transferID int32) (FetchSpendProofsRow, error)
	GenesisAssets(ctx context.Context) ([]GenesisAsset, error)
	GenesisPoints(ctx context.Context) ([]GenesisPoint, error)
	GetRootKey(ctx context.Context, id []byte) (Macaroon, error)
	InsertAddr(ctx context.Context, arg InsertAddrParams) (int32, error)
	InsertAssetDelta(ctx context.Context, arg InsertAssetDeltaParams) error
	InsertAssetSeedling(ctx context.Context, arg InsertAssetSeedlingParams) error
	InsertAssetSeedlingIntoBatch(ctx context.Context, arg InsertAssetSeedlingIntoBatchParams) error
	InsertAssetTransfer(ctx context.Context, arg InsertAssetTransferParams) (int32, error)
	InsertAssetWitness(ctx context.Context, arg InsertAssetWitnessParams) error
	InsertBranch(ctx context.Context, arg InsertBranchParams) error
	InsertCompactedLeaf(ctx context.Context, arg InsertCompactedLeafParams) error
	InsertLeaf(ctx context.Context, arg InsertLeafParams) error
	InsertNewAsset(ctx context.Context, arg InsertNewAssetParams) (int32, error)
	InsertRootKey(ctx context.Context, arg InsertRootKeyParams) error
	InsertSpendProofs(ctx context.Context, arg InsertSpendProofsParams) (int32, error)
	NewMintingBatch(ctx context.Context, arg NewMintingBatchParams) error
	// We use a LEFT JOIN here as not every asset has a family key, so this'll
	// generate rows that have NULL values for the family key fields if an asset
	// doesn't have a family key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	QueryAssetBalancesByAsset(ctx context.Context, assetIDFilter interface{}) ([]QueryAssetBalancesByAssetRow, error)
	QueryAssetBalancesByFamily(ctx context.Context, keyFamFilter interface{}) ([]QueryAssetBalancesByFamilyRow, error)
	QueryAssetTransfers(ctx context.Context, arg QueryAssetTransfersParams) ([]QueryAssetTransfersRow, error)
	// We use a LEFT JOIN here as not every asset has a family key, so this'll
	// generate rows that have NULL values for the family key fields if an asset
	// doesn't have a family key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	// This clause is used to select specific assets for a asset ID, general
	// channel balances, and also coin selection. We use the sqlc.narg feature to
	// make the entire statement evaluate to true, if none of these extra args are
	// specified.
	QueryAssets(ctx context.Context, arg QueryAssetsParams) ([]QueryAssetsRow, error)
	QueryEventIDs(ctx context.Context, arg QueryEventIDsParams) ([]QueryEventIDsRow, error)
	ReanchorAssets(ctx context.Context, arg ReanchorAssetsParams) error
	SetAddrManaged(ctx context.Context, arg SetAddrManagedParams) error
	UpdateBatchGenesisTx(ctx context.Context, arg UpdateBatchGenesisTxParams) error
	UpdateMintingBatchState(ctx context.Context, arg UpdateMintingBatchStateParams) error
	UpsertAddrEvent(ctx context.Context, arg UpsertAddrEventParams) (int32, error)
	UpsertAssetFamilyKey(ctx context.Context, arg UpsertAssetFamilyKeyParams) (int32, error)
	UpsertAssetFamilySig(ctx context.Context, arg UpsertAssetFamilySigParams) (int32, error)
	UpsertAssetProof(ctx context.Context, arg UpsertAssetProofParams) error
	UpsertChainTx(ctx context.Context, arg UpsertChainTxParams) (int32, error)
	UpsertGenesisAsset(ctx context.Context, arg UpsertGenesisAssetParams) (int32, error)
	UpsertGenesisPoint(ctx context.Context, prevOut []byte) (int32, error)
	UpsertInternalKey(ctx context.Context, arg UpsertInternalKeyParams) (int32, error)
	UpsertManagedUTXO(ctx context.Context, arg UpsertManagedUTXOParams) (int32, error)
	UpsertRootNode(ctx context.Context, arg UpsertRootNodeParams) error
	UpsertScriptKey(ctx context.Context, arg UpsertScriptKeyParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
