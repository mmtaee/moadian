package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func Sign(text string, privateKey *rsa.PrivateKey) ([]byte, string) {
	hash := sha256.Sum256([]byte(text))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	return signature, base64.StdEncoding.EncodeToString(signature)
}

func Verify(data string, signature []byte, privateKey *rsa.PrivateKey) error {
	textBytes := []byte(data)
	hash := sha256.Sum256(textBytes)
	return rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hash[:], signature)
}
