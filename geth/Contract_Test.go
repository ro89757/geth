package geth

import (
	"context"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func ContractTest() {
	// 连接到 Ethereum 节点（可以是本地节点或者 Infura 提供的节点）
	client := comm.Client()

	// 需要检查的合约地址
	contractAddress := common.HexToAddress("0x61458f8ffa1F3bF9D477Aa70522CBF10d1eF61E5")

	// 获取指定地址的代码
	code, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatalf("Failed to get contract code: %v", err)
	}

	// 如果代码为空，表示该地址没有合约
	if len(code) == 0 {
		fmt.Println("No contract code found at the specified address!")
	} else {
		fmt.Println("Contract code found at the address.")
	}
}
