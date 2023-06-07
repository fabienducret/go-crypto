package main

import (
	"cryptography/decryption"
	"cryptography/encryption"
	"cryptography/repositories"
	"fmt"
)

func main() {
	encrypt := encryption.New(
		repositories.NewPublicKeyRepository(),
	)
	decrypt := decryption.New(
		repositories.NewPrivateKeyRepository(),
	)

	toEncrypt := []byte("Hello, world !")
	toDecrypt, err := encrypt.From(toEncrypt)
	if err != nil {
		fmt.Println(err)
	}

	decrypted, err := decrypt.From(toDecrypt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(decrypted))
}
