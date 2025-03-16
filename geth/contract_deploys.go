package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"github.com/code/go-ethereum/stores"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func ContractDeployS() {
	client := comm.Client()

	//privateKey, err := crypto.GenerateKey()
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//privateKeyHex := hex.EncodeToString(privateKeyBytes)
	//fmt.Println("Private Key:", privateKeyHex)
	//a81980631cdd38cac1b2d20d75d51e2a5b2ace45475a69a69d865ec9309d7d5c
	//2debf1367e5e2105bdacac7dfacf3cb86e54afc2922bf4f47d6a425c16bf11b9
	privateKey := comm.PrivateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	println("ContractDeployS nonce:", nonce)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background()) //11155111
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	println("ContractDeployS gasPrice:", gasPrice)
	//auth.GasPrice = gasPrice
	auth.GasPrice = big.NewInt(int64(9504439364)) //7504439364

	input := "1.0"
	address, tx, instance, err := stores.DeployStores(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ContractDeployS:", address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

	//	ContractDeployS gasPrice: 0xc000180b80
	//ContractDeployS: 0x130ed9Fd8aeb5F9087Cc9f2c450B59fB6f38Dc4A
	//	0x2691bea18b5bd2181ca736377d7a9ccf6a08a5cfdbe21fd6ac4d204d1573aa04

	//自定义12
	//ContractDeployS: 0x61458f8ffa1F3bF9D477Aa70522CBF10d1eF61E5
	//	0x0135846bfd29cae270846040e3a3ed630f21bc0d834f0d1bea88e7c94d90a156
	//stores 自定义 15
	//ContractDeployS: 0x79426b6712bDf7143B2dbBe5B0C605A618e34E28
	//0x528cb77ddd387ab2b39533d8e61c7c034b7bea0f4ce2fb79cf039b81174c6f1d

}
