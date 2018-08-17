package transaction

import (
	"testing"

	"encoding/json"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

//func TestNewTransactionFailsWithUnknownOperation(t *testing.T) {
//
//	txn, err := New("UNKNOWN_OPERATION", Asset{}, nil, nil, nil)
//
//	assert.Error(t, err, "Error while making new Create Transaction")
//	assert.Equal(t, txn.ID, "", "Returned txn should be empty")
//}
//
//func TestNewCreateTransaction(t *testing.T) {
//	// Given
//	asset := asset()
//	metadata := metadata()
//	outputs := outputs()
//	publicKeys := publicKeys()
//	expTxn := Transaction{
//		Inputs: []Input{{
//			Fulfills:     OutputLocation{},
//			Fulfillment:  "",
//			OwnersBefore: publicKeys,
//		}},
//	}
//
//	// When
//	txn, err := NewCreateTransaction(asset, metadata, outputs, publicKeys)
//
//	// Then
//	assert.NoError(t, err, "NewCreateTransaction Should not give error")
//	assert.NotNil(t, txn, "Expecting not nil for Txn")
//	assert.Equal(t, expTxn.Inputs[0].OwnersBefore[0], txn.Inputs[0].OwnersBefore[0], "Input should have public keys set correctly")
//}
//
//func TestNewOutputDefaultsToAmountOne(t *testing.T) {
//
//	// Given
//	amount := ""
//	condition := condition()
//
//	// When
//	output, _ := NewOutput(condition, amount)
//
//	// Then
//	assert.NotNil(t, output, "Output should not be nil")
//	assert.Equal(t, "1", output.Amount)
//}

//func TestNewOutputCreatesUniquePubKeysSlice(t *testing.T) {
//
//	// Given
//	condition := condition()
//	// Add another identical subcondition
//	condition.Details = append(condition.Details, subcondition())
//
//	// When
//	output, _ := NewOutput(condition, "1")
//
//	// Then
//	assert.Equal(t, 1, len(output.PublicKeys))
//}

func TestCalculateID(t *testing.T) {

	// Given
	ctxnBytes, err := ioutil.ReadFile("../../fixtures/mock_create_txn.json")
	assert.NoError(t, err)

	var createTxn *Transaction
	_ = json.Unmarshal(ctxnBytes, &createTxn)

	// When
	id, err := createTxn.createID()
	assert.NoError(t, err, "Could not create ID")

	// Then
	assert.Equal(t, "4957744b3ac54434b8270f2c854cc1040228c82ea4e72d66d2887a4d3e30b317", id)

}

//func TestCreateCondition(t *testing.T) {
//	// Given
//	pubKey := []byte("4zvwRjXUKGfvwnParsHAS3HuSVzV5cA4McphgmoCtajS")
//
//	// When
//	cond, err := NewEd25519Condition(pubKey)
//	assert.NoError(t, err)
//
//	// Then
//	assert.Equal(t, "ni:///sha-256;uLdVX7FEjLWVDkAkfMAkEqPPwFqZj7qfiGE152t_x5c?fpt=ed25519-sha-256&cost=131072", cond.Uri, "Expected matching condition URIs")
//}

//func asset() Asset {
//	data := make(map[string]interface{})
//	data["message"] = rand.Int()
//
//	asset := Asset{
//		Data: data,
//	}
//	return asset
//}
//
//func metadata() Metadata {
//	md := make(map[string]interface{})
//	md["planet"] = "earth"
//	return md
//}
//
//func outputs() []Output {
//	cond := condition()
//	op := Output{
//		Amount:     "1",
//		Condition:  cond,
//		PublicKeys: publicKeys(),
//	}
//	res := make([]Output, 1)
//	res[0] = op
//	return res
//}
//
//var s = []byte("GtMiC8DQ274d3wsW7R4MMNvMU6DLcd4U9vg43tu9fFTp")
//
//func publicKeys() []ed25519.PublicKey {
//	return []ed25519.PublicKey{s}
//}
//
//func condition() Condition {
//	det := cryptoconditions.NewSimpleCondition(cryptoconditions.CTEd25519Sha256, s, 1)
//	return Condition{
//		Details: det,
//		Uri:     "ni:///sha-256;raKXoRoO8MhGqW0XqOWNUXoheQzS-0NLHvnWPpOcwIg?fpt=ed25519-sha-256&cost=131072",
//	}
//}

//func subcondition() SubCondition {
//	return SubCondition{
//		PublicKey: "GtMiC8DQ274d3wsW7R4MMNvMU6DLcd4U9vg43tu9fFTp",
//		Type:      "ed25519-sha-256",
//	}
//}
