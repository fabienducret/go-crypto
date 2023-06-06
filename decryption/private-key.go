package decryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func privateKeyFrom(bytes []byte) (*rsa.PrivateKey, error) {
	block, rest := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid private key")
	}

	if len(rest) > 0 {
		return nil, errors.New("extra data included in key")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
