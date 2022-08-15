package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/ttvs-blockchain/local-chain/internal/ledger"
)

const (
	numTXs     = 90
	bindingLen = 32
)

func main() {
	lc := ledger.NewController()
	defer lc.Close()
	// lc.GetAllTXs()
	// return
	for i := 0; i < numTXs; i++ {
		b := createDummyTX()
		id, err := lc.SubmitTX(b)
		handleErr(err)
		log.Println("No.", i, "id: ", id)
	}
}

func createDummyTX() string {
	byteBinding := make([]byte, bindingLen)
	_, err := rand.Read(byteBinding)
	if err != nil {
		panic(err)
	}
	strBinding := hex.EncodeToString(byteBinding)
	return strBinding
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
