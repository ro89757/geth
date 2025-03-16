package geth

import (
	"context"
	"fmt"
	"github.com/code/go-ethereum/comm"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func BlockSubscribe() {
	client := comm.Client()
	headers := make(chan *types.Header) //创建一个新的通道，用于接收最新的区块头
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
