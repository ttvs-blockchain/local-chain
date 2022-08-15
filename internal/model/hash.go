package model

// Hash is the data type stored on local chain
type Hash struct {
	Hash []byte `json:"hash"`
}

// Serialize serializes a Binding struct into a byte slice
func (h *Hash) Serialize() ([]byte, error) {
	return h.Hash, nil
}
