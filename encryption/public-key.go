package encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func publicKeyFrom(bytes []byte) (*rsa.PublicKey, error) {
	block, rest := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid public key")
	}

	if len(rest) > 0 {
		return nil, errors.New("extra data included in key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey.(*rsa.PublicKey), nil
}
