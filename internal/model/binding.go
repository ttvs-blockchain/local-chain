package model

// Binding is the binding of personal information hash and certificate information hash
type Binding struct {
	PersonInfoHash []byte `json:"person_info_hash"`
	CertInfoHash   []byte `json:"certificate"`
}

// Serialize serializes a Binding struct into a byte slice
func (b *Binding) Serialize() ([]byte, error) {
	return append(b.PersonInfoHash, b.CertInfoHash...), nil
}
