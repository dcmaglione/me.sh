package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"dcmaglione.com/me.sh/crypto"
)

func main() {
	passphrase := "supersecurepassword"
	plaintext := "Hello, encrypted world!"

	fmt.Println("Original plaintext:", plaintext)

	encrypted, err := crypto.Encrypt([]byte(plaintext), passphrase)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	fmt.Println("Encrypted (hex):", hex.EncodeToString(encrypted))

	decrypted, err := crypto.Decrypt(encrypted, passphrase)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted plaintext:", string(decrypted))
}
