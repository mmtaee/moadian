package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
)

// encrypt function return cipher text and iv(nonce)
func encrypt(aesKey, xorText, iv []byte) (string, []byte) {
	block, _ := aes.NewCipher(aesKey)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	ciphertext := gcm.Seal(nil, nonce, xorText, nil)
	tag := gcm.Seal(nil, nonce, nil, nil)
	ciphertextWithTag := append(ciphertext, tag...)
	return base64.StdEncoding.EncodeToString(ciphertextWithTag), nonce
}

// xor function make xor from json data and aesKey
func xor(a, b []byte) []byte {
	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i%len(b)]
	}
	return result
}

// encryptSymmetricKey function encrypt symmetricKey with services org public key
func encryptSymmetricKey(symmetricKey []byte, publicKey *rsa.PublicKey) (string, error) {
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, symmetricKey, nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
