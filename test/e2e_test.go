package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"fmt"

	cl "github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

const (
	URL = "http://localhost:9984/api/v1/"
)

func TestAssetCreationAndTransferE2E(t *testing.T) {

	// -- PREPARE CREATE TRANSACTION -- //
	var keyPairs []*txn.KeyPair

	keyPair, err := txn.NewKeyPair()
	assert.NoError(t, err, "Could not create keypairs!")

	keyPairs = append(keyPairs, keyPair)

	var outputs []txn.Output
	var issuers []ed25519.PublicKey
	for _, keyPair := range keyPairs {
		// Create conditions
		condition := txn.NewEd25519Condition(keyPair.PublicKey)

		// Create output
		output, err := txn.NewOutput(*condition, "1")
		assert.NoError(t, err, "Cannot creat output")

		outputs = append(outputs, output)

		// Create issuers
		issuers = append(issuers, keyPair.PublicKey)
	}

	data := make(map[string]interface{})
	data["assetID"] = "testID1"
	asset := txn.Asset{
		Data: data,
	}
	metadata := make(map[string]interface{})
	metadata["planet"] = "earth"

	// New create transaction
	txn, err := txn.NewCreateTransaction(asset, metadata, outputs, issuers)
	assert.NoError(t, err, "Not able to create txn")

	assert.NotNil(t, txn)

	// Sign transaction
	err = txn.Sign(keyPairs)
	assert.NoError(t, err, "Not able to sign txn")

	fmt.Println("*** SIGNED TXN:")
	fmt.Println(txn.String())

	// -- SEND TRANSACTION -- //

	// Create client
	cfg := cl.ClientConfig{
		Url: URL,
	}
	client, err := cl.New(cfg)
	assert.NoError(t, err, "Could not create client")

	// -- GET BLOCK -- //
	block1, err := client.GetBlock("0")
	assert.NoError(t, err, "Could not get block 0")
	fmt.Printf("Block height: %d\n", block1.Height)

	// Post transaction
	err = client.PostTransaction(txn)
	assert.NoError(t, err, "Could not post transaction")

	assert.NotNil(t, block1)

	// -- LIST TRANSACTION -- //
	txns, err := client.ListTransactions("1", "CREATE")
	assert.NoError(t, err, "Could not get list of transactions for asset with ID %s", asset.ID)

	assert.NotNil(t, txns)
	//assert.Equal(t, 1, len(txns))

	// -- PREPARE TRANSFER TRANSACTION -- //

}

func TestNewClientCreatesConnectionToTestNetwork(t *testing.T) {

	// Make sure to add dbd_key.json file for your own test network
	headerBytes, err := ioutil.ReadFile("../fixtures/dbd_key.json")
	assert.NoError(t, err)

	var h map[string]string
	err = json.Unmarshal(headerBytes, &h)
	assert.NoError(t, err)

	cfg := cl.ClientConfig{
		Url:     "https://test.bigchaindb.com/api/v1/",
		Headers: h,
	}

	client, err := cl.New(cfg)

	assert.NoError(t, err, "Could not create client")

	block, err := client.GetBlock("0")

	assert.NoError(t, err, "Could not retrieve genesis block")

	assert.Equal(t, 0, block.Height)
}

func TestNewClientCreatesConnectionLocalhost(t *testing.T) {

	cfg := cl.ClientConfig{
		Url: URL,
	}

	client, err := cl.New(cfg)

	assert.NoError(t, err, "Could not create client")

	block, err := client.GetBlock("0")

	assert.NoError(t, err, "Could not retrieve genesis block")

	assert.Equal(t, 0, block.Height)
}

func TestPostCreateTransaction(t *testing.T) {

	ctxnBytes, err := ioutil.ReadFile("../fixtures/mock_create_txn.json")
	assert.NoError(t, err)

	var createTxn *txn.Transaction
	_ = json.Unmarshal(ctxnBytes, &createTxn)

	cfg := cl.ClientConfig{
		Url: URL,
	}

	client, _ := cl.New(cfg)

	err = client.PostTransaction(createTxn)

	assert.NoError(t, err, "Could not post transaction")
}
