package main

import "github.com/ttvs-blockchain/local-chain/internal/ledger"

func main() {
	lc := ledger.NewController()
	defer lc.Close()
	id, err := lc.SubmitTX("test", 1)
	handleErr(err)
	err = lc.GetAllTXs()
	handleErr(err)
	err = lc.FindTX(id)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
