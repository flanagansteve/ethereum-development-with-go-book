package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

func main() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyID) // 0ec5cfe4e215239756054992dbc2e10f011db1cdfc88b9ba6301e2f9ea1b58d2

	publicKey, err := client.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(hexutil.Encode(publicKey)) // 0x04f17356fd52b0d13e5ede84f998d26276f1fc9d08d9e73dcac6ded5f3553405db38c2f257c956f32a0c1fca4c3ff6a38a2c277c1751e59a574aecae26d3bf5d1d

	messageHash, err := client.Post(context.Background(), whisperv6.NewMessage{
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Payload:   []byte("Hello"),
		PublicKey: publicKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageHash) // 0xdbfc815d3d122a90d7fb44d1fc6a46f3d76ec752f3f3d04230fe5f1b97d2209a
}
