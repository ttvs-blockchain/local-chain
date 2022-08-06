package main

import (
	"github.com/ttvs-blockchain/local-chain/internal/model"
	"github.com/ttvs-blockchain/local-chain/internal/proof"
	"math/rand"
)

func main() {
	bindings := make([]*model.Binding, 10)
	for i := 0; i < 10; i++ {
		bindings[i] = &model.Binding{
			PersonInfoHash: make([]byte, 32),
			CertInfoHash:   make([]byte, 32),
		}
		_, err := rand.Read(bindings[i].PersonInfoHash)
		handleErr(err)
		_, err = rand.Read(bindings[i].CertInfoHash)
		handleErr(err)
	}
	root, proves, err := proof.Generate(bindings)
	handleErr(err)
	for i := 0; i < 10; i++ {
		ok, err := proof.Verify(bindings[i], proves[i], root)
		handleErr(err)
		if !ok {
			panic("verification failed")
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
