package encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(pemData)
	if pemBlock == nil || pemBlock.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid PEM block")
	}
	var keyPKCS1 *rsa.PrivateKey
	keyPKCS1, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err == nil {
		return keyPKCS1, nil
	}
	var keyPKCS8 any
	keyPKCS8, err = x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err == nil {
		privateKey, ok := keyPKCS8.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
		return privateKey, nil
	}
	return nil, err
}
