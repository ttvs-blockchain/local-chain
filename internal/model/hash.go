package model

// Hash is the data type stored on local chain
type Hash []byte

// Serialize serializes a Binding struct into a byte slice
func (h Hash) Serialize() ([]byte, error) {
	return h, nil
}
