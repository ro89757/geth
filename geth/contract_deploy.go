package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"github.com/code/go-ethereum/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func ContractDeploy() {
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
	println("gasPrice:", gasPrice)
	auth.GasPrice = big.NewInt(int64(7504439364)) //7504439364

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

	//0x426DEB4A2EAE7A6164D5260a2aB1638C791DeCD6
	//0xf03b46341943b4c38d76d1fc74cb703c0f4cd0800e65d60a0f0665a4fc0c3aaa

}
