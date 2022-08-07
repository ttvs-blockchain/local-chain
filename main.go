package main

import (
	"log"
	"time"

	"github.com/ttvs-blockchain/local-chain/internal/ledger"
)

func main() {
	lc := ledger.NewController()
	defer lc.Close()
	id, err := lc.SubmitTX("test", time.Now().UnixNano())
	handleErr(err)
	err = lc.GetAllTXs()
	handleErr(err)
	log.Println("id: ", id)
	log.Println("FindTX")
	err = lc.FindTX(id)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
