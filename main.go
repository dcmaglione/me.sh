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

	// Encrypt
	encrypted, err := crypto.Encrypt([]byte(plaintext), passphrase)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	// Print ciphertext as hex for readability
	fmt.Println("Encrypted (hex):", hex.EncodeToString(encrypted))

	// Decrypt
	decrypted, err := crypto.Decrypt(encrypted, passphrase)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted plaintext:", string(decrypted))
}
