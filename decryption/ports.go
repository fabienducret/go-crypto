package decryption

import "crypto/rsa"

type KeyRepository interface {
	Get() ([]byte, error)
}

type KeyBuilder interface {
	From([]byte) (*rsa.PrivateKey, error)
}
