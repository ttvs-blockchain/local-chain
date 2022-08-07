package proof

import (
	"sync"

	mt "github.com/tommytim0515/go-merkletree"
	"github.com/ttvs-blockchain/local-chain/internal/model"
)

var (
	tree *mt.MerkleTree
	once sync.Once
)

// Generate generates a Merkle Tree membership proof for a list of bindings (person info || cert into)
func Generate(bindings []*model.Binding) (root []byte, proves []*mt.Proof, err error) {
	once.Do(func() {
		tree = mt.NewMerkleTree(&mt.Config{
			AllowDuplicates: true,
		})
	})
	defer tree.Reset()
	blocks := make([]mt.DataBlock, len(bindings))
	for i, binding := range bindings {
		blocks[i] = binding
	}
	err = tree.Build(blocks)
	if err != nil {
		return nil, nil, err
	}
	root = tree.Root
	proves = tree.Proves
	return
}

// Verify verifies a binding with the Merkle Tree proof and root
func Verify(binding *model.Binding, proof *mt.Proof, root []byte) (bool, error) {
	return mt.Verify(binding, proof, root, nil)
}
