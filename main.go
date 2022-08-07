package main

import "github.com/ttvs-blockchain/local-chain/internal/ledger"

func main() {
	lc := ledger.NewController()
	defer lc.Close()
	err := lc.SubmitTx("test", 1)
	handleErr(err)
	err = lc.GetAllTXs()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
