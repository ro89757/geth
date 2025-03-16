package geth

import (
	"context"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"log"
	"math/big"
)

func Blcoks() {
	client := comm.Client()
	// 查询最新区块的头信息 nil
	//header, err := client.HeaderByNumber(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(header.Number.String())
	// 查询指定区块的头信息
	blockNumber := big.NewInt(7788526)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 7788526
	fmt.Println(header.Time)                // 1740559224
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xc4af3a714c04023fcea4808ce99ef265e2a0dee38a155130f1cd00ae0e6901e2

	if err != nil {
		log.Fatal(err)
	}
	// 查询指定区块的全部信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 7788526
	fmt.Println(block.Time())                // 1740559224
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xc4af3a714c04023fcea4808ce99ef265e2a0dee38a155130f1cd00ae0e6901e2
	fmt.Println(len(block.Transactions()))   // 192
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 192
}
