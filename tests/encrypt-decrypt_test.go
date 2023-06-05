package tests

import (
	"cryptography/builders"
	"cryptography/decryption"
	"cryptography/encryption"
	"cryptography/tests/stubs"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	encrypt := encryption.New(
		stubs.NewPublicKeyRepository(),
		builders.NewPublicKeyBuilder(),
	)
	decrypt := decryption.New(
		stubs.NewPrivateKeyRepository(),
		builders.NewPrivateKeyBuilder(),
	)

	t.Run("encrypt and decrypt a message", func(t *testing.T) {
		//Given
		toEncrypt := []byte("message to encypt")

		// When
		toDecrypt, _ := encrypt.From(toEncrypt)
		decrypted, _ := decrypt.From(toDecrypt)

		// Then
		assertEqual(t, string(decrypted), string(toEncrypt))
	})

	t.Run("get error on invalid public key", func(t *testing.T) {
		//Given
		encrypt := encryption.New(
			stubs.NewInvalidPublicKeyRepository(),
			builders.NewPublicKeyBuilder(),
		)
		toEncrypt := []byte("message to encypt")

		// When
		_, err := encrypt.From(toEncrypt)

		// Then
		assertEqual(t, err.Error(), "invalid public key")
	})
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Error, got %s, want %s", got, want)
	}
}
