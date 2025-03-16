package geth

import (
	"github.com/code/go-ethereum/comm"
	"log"

	"github.com/code/go-ethereum/store"
	"github.com/ethereum/go-ethereum/common"
)

const (
	contractAddr = "0x426DEB4A2EAE7A6164D5260a2aB1638C791DeCD6"
)

func ContractLoad() {
	client := comm.Client()
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract
}
