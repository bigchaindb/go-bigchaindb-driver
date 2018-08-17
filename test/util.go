package test

import (
	"github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
)

func getKeyPair() *transaction.KeyPair {
	return &transaction.KeyPair{
		PublicKey:  []byte{238, 233, 56, 119, 156, 141, 42, 29, 253, 127, 90, 143, 85, 176, 232, 94, 61, 228, 84, 134, 62, 198, 205, 11, 230, 140, 98, 158, 27, 200, 224, 88},
		PrivateKey: []byte{128, 45, 107, 177, 154, 49, 30, 71, 155, 12, 49, 165, 136, 6, 78, 168, 107, 237, 101, 235, 199, 189, 205, 128, 116, 57, 162, 38, 194, 216, 81, 145, 238, 233, 56, 119, 156, 141, 42, 29, 253, 127, 90, 143, 85, 176, 232, 94, 61, 228, 84, 134, 62, 198, 205, 11, 230, 140, 98, 158, 27, 200, 224, 88},
	}
}

func AlicePubKey() string {
	return "G7J7bXF8cqSrjrxUKwcF8tCriEKC5CgyPHmtGwUi4BK3"
}

func AlicePrivKey() string {
	return "CT6nWhSyE7dF2znpx3vwXuceSrmeMy9ChBfi9U92HMSP"
}

func BobPubKey() string {
	return "2dBVUoATxEzEqRdsi64AFsJnn2ywLCwnbNwW7K9BuVuS"
}

func BobPrivKey() string {
	return "4S1dzx3PSdMAfs59aBkQefPASizTs728HnhLNpYZWCad"
}

func AliceKeyPair() transaction.KeyPair {
	return transaction.KeyPair{
		PrivateKey: []byte(AlicePrivKey()),
		PublicKey:  []byte(AlicePubKey()),
	}
}

func BobKeyPair() transaction.KeyPair {
	return transaction.KeyPair{
		PrivateKey: []byte(BobPrivKey()),
		PublicKey:  []byte(BobPubKey()),
	}
}
