package decryption

import (
	"crypto"
	"crypto/rsa"
)

type decryption struct {
	keyRepository KeyRepository
}

func New(keyRepository KeyRepository) *decryption {
	return &decryption{
		keyRepository: keyRepository,
	}
}

func (d *decryption) From(ciphertext []byte) ([]byte, error) {
	bytes, err := d.keyRepository.Get()
	if err != nil {
		return nil, err
	}

	key, err := privateKeyFrom(bytes)
	if err != nil {
		return nil, err
	}

	plaintext, err := key.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
