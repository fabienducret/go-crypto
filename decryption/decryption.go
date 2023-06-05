package decryption

import (
	"crypto"
	"crypto/rsa"
)

type decryption struct {
	keyRepository KeyRepository
	keyBuilder    KeyBuilder
}

func New(keyRepository KeyRepository, keyBuilder KeyBuilder) *decryption {
	return &decryption{
		keyRepository: keyRepository,
		keyBuilder:    keyBuilder,
	}
}

func (d *decryption) From(ciphertext []byte) ([]byte, error) {
	bytes, err := d.keyRepository.Get()
	if err != nil {
		return nil, err
	}

	key, err := d.keyBuilder.From(bytes)
	if err != nil {
		return nil, err
	}

	plaintext, err := key.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
