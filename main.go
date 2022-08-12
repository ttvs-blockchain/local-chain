package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/ttvs-blockchain/local-chain/internal/ledger"
)

const (
	numTXs     = 90
	bindingLen = 64
)

func main() {
	lc := ledger.NewController()
	defer lc.Close()
	// lc.GetAllTXs()
	// return
	for i := 0; i < numTXs; i++ {
		b, t := createDummyTX()
		id, err := lc.SubmitTX(b, t)
		handleErr(err)
		log.Println("No.", i, "id: ", id)
	}
}

func createDummyTX() (string, int64) {
	byteBinding := make([]byte, bindingLen)
	_, err := rand.Read(byteBinding)
	if err != nil {
		return "", 0
	}
	strBinding := hex.EncodeToString(byteBinding)
	timestamp := time.Now().UnixNano()
	// strTimestamp := strconv.Itoa(int(timestamp))
	return strBinding, timestamp
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
