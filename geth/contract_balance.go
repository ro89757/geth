package geth

import (
	"fmt"
	"github.com/code/go-ethereum/comm"
	token "github.com/code/go-ethereum/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math"
	"math/big"
)

func ContractBalance() {
	client := comm.Client()
	// Golem (GNT) Address
	tokenAddress := comm.ContractAddress()
	instance, err := token.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := comm.ToAddress()
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	//name: RccToken
	//symbol: RCC
	//decimals: 18
	//wei: 19999000000000000000000000
	//balance: 19999000.000000
	//
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
