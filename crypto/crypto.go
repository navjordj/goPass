// Package crypto osv
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	_ "crypto/sha256" // trengs ikke?
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/ssh/terminal"
)

// https://bruinsslot.jp/post/golang-crypto, https://itnext.io/encrypt-data-with-a-password-in-go-b5366384e291

// Encrypt osv
func Encrypt(key, data []byte) ([]byte, error) {
	key, salt, err := DeriveKey(key, nil)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = append(ciphertext, salt...)
	return ciphertext, nil
}

// Decrypt osv
func Decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]
	key, _, err := DeriveKey(key, salt)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

//DeriveKey osv
func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}

func main() {
	// password := []byte("mysecretpassword")
	fmt.Print("Enter Password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Println()

	data := []byte("our super secret text")

	ciphertext, err := Encrypt(password, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ciphertext: %s\n", hex.EncodeToString(ciphertext))

	plaintext, err := Decrypt(password, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("plaintext: %s\n", plaintext)
}
