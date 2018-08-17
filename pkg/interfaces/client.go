package interfaces

import "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"

type Client interface {
	GetBlock(blockHeight string) (transaction.Block, error)
	GetTransaction(transactionId string) (transaction.Transaction, error)
	ListBlocks(transactionID string) ([]transaction.Block, error)
	ListOutputs(pubKey string, spent bool) ([]transaction.OutputLocation, error)
	ListTransactions(assetID, operation string) ([]transaction.Transaction, error)
	ListVotes(blockID string) (interface{}, error)
	PostTransaction(txn transaction.Transaction, mode string) (transaction.Transaction, error)
	SearchAsset(search string, limit int) ([]transaction.Asset, error)
	SearchMetadata(search string, limit int) ([]transaction.Metadata, error)
}

//'blocks': 'blocks',
//'blocksDetail': 'blocks/%(blockHeight)s',
//'outputs': 'outputs',
//'transactions': 'transactions',
//'transactionsSync': 'transactions?mode=sync',
//'transactionsCommit': 'transactions?mode=commit',
//'transactionsDetail': 'transactions/%(transactionId)s',
//'assets': 'assets',
//'metadata': 'metadata',
//'votes': 'votes'
