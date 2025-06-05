package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"dcmaglione.com/me.sh/crypto"
	"dcmaglione.com/me.sh/storage"
)

func main() {
	passphrase := "supersecurepassword"
	plaintext := "Hello, encrypted world!"

	fmt.Println("Original plaintext:", plaintext)

	// encrypt
	encrypted, err := crypto.Encrypt([]byte(plaintext), passphrase)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Println("Encrypted (hex):", hex.EncodeToString(encrypted))

	// save encrypted post
	filename, err := storage.SaveEncryptedPost(encrypted)
	if err != nil {
		log.Fatalf("Saving encrypted post failed: %v", err)
	}
	fmt.Println("Saved encrypted post as:", filename)

	// load all saved posts
	files, err := storage.LoadEncryptedPosts()
	if err != nil {
		log.Fatalf("Loading post list failed: %v", err)
	}
	fmt.Println("Available encrypted posts:", files)

	// read the saved post
	readData, err := storage.ReadEncryptedPost(filename)
	if err != nil {
		log.Fatalf("Reading encrypted post failed: %v", err)
	}

	// decrypt
	decrypted, err := crypto.Decrypt(readData, passphrase)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Println("Decrypted plaintext:", string(decrypted))

	// delete post afterward
	err = storage.DeletePostFile(filename)
	if err != nil {
		log.Fatalf("Deleting post failed: %v", err)
	}
	fmt.Println("Deleted encrypted post:", filename)
}
