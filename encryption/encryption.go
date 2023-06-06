package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type encryption struct {
	keyRepository KeyRepository
}

func New(keyRepository KeyRepository) *encryption {
	return &encryption{
		keyRepository: keyRepository,
	}
}

func (e *encryption) From(plaintext []byte) ([]byte, error) {
	bytes, err := e.keyRepository.Get()
	if err != nil {
		return nil, err
	}

	key, err := publicKeyFrom(bytes)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, key, plaintext, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
