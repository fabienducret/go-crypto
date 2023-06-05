package main

import (
	"cryptography/builders"
	"cryptography/decryption"
	"cryptography/encryption"
	"cryptography/repositories"
	"fmt"
	"log"
)

func main() {
	encrypt := encryption.New(
		repositories.NewPublicKeyRepository(),
		builders.NewPublicKeyBuilder(),
	)
	decrypt := decryption.New(
		repositories.NewPrivateKeyRepository(),
		builders.NewPrivateKeyBuilder(),
	)

	toEncrypt := []byte("Hello, world !")
	toDecrypt, err := encrypt.From(toEncrypt)
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := decrypt.From(toDecrypt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decrypted))
}
