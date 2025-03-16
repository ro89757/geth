package comm

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func Client() *ethclient.Client {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/xx")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("we have a connection")
	return client
}

func PrivateKey() *ecdsa.PrivateKey { // 私钥
	privateKey, err := crypto.HexToECDSA("Private Key")

	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func ToAddress() common.Address { // 目标地址
	toAddress := common.HexToAddress("xx")
	return toAddress
}
func ContractAddress() common.Address { //合约地址
	contractAddress := common.HexToAddress("xx")
	return contractAddress
}
