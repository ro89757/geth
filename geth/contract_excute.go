package geth

import (
	"context"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"log"
	"math/big"

	"github.com/code/go-ethereum/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ContractExcute() {
	client := comm.Client()
	privateKey := comm.PrivateKey()

	storeContract, err := store.NewStore(common.HexToAddress("0x426DEB4A2EAE7A6164D5260a2aB1638C791DeCD6"), client)
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))
	// 初始化交易opt实例
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	// 调用合约方法
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("valueInContract:", valueInContract)
	fmt.Println("value:", value)
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
	//	tx hash: 0x29ecd3da800cb0f496d923e724dba645f6b9a55c2dc4281c43aa2142d9809217
	//valueInContract: [100 101 109 111 95 115 97 118 101 95 118 97 108 117 101 49 49 49 49 49 0 0 0 0 0 0 0 0 0 0 0 0]
	//	value: [100 101 109 111 95 115 97 118 101 95 118 97 108 117 101 49 49 49 49 49 0 0 0 0 0 0 0 0 0 0 0 0]
	//	is value saving in contract equals to origin value: true
}
