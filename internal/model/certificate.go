package model

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
)

type Certificates struct {
	CerID       string `json:"certificate"`
	PersonSysID string `json:"person_sys_id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	NumOfDose   string `json:"num_of_dose"`
	Time        string `json:"time"`
	Issuer      string `json:"issuer"`
	Remark      string `json:"remark"`
	Payload     []byte `json:"payload"`
}

// Encode encodes a Certificates struct into a byte slice
func (c *Certificates) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}
	return compress(buf.Bytes())
}

// compress a certificate byte slice with gzip
func compress(b []byte) (bt []byte, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	defer func(zw *gzip.Writer) {
		err = zw.Close()
	}(zw)
	if _, err = zw.Write(b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode decodes a byte slice into a Certificates struct
func Decode(data []byte) (*Certificates, error) {
	decompressed, err := decompress(data)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.Write(decompressed)
	dec := gob.NewDecoder(&buf)
	var c *Certificates
	if err = dec.Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}

// decompress a certificate byte slice with gzip
func decompress(b []byte) (bt []byte, err error) {
	var reader *gzip.Reader
	if reader, err = gzip.NewReader(bytes.NewReader(b)); err != nil {
		return nil, err
	}
	defer func(reader *gzip.Reader) {
		err = reader.Close()
	}(reader)
	return io.ReadAll(reader)
}
