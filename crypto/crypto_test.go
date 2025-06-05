package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	passphrase := "mysecret"
	plaintext := []byte("test message")

	encrypted, err := Encrypt(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, passphrase)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Decrypted text does not match original.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

func TestWrongPassphrase(t *testing.T) {
	passphrase := "correct-pass"
	wrongPass := "wrong-pass"
	plaintext := []byte("this should fail")

	encrypted, err := Encrypt(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	_, err = Decrypt(encrypted, wrongPass)
	if err == nil {
		t.Errorf("Decryption should have failed with wrong passphrase")
	}
}

func TestCorruptedCiphertext(t *testing.T) {
	passphrase := "another-pass"
	plaintext := []byte("don't mess with this")

	encrypted, err := Encrypt(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	encrypted[len(encrypted)-1] ^= 0xFF

	_, err = Decrypt(encrypted, passphrase)
	if err == nil {
		t.Errorf("Decryption should have failed due to corrupted ciphertext")
	}
}
