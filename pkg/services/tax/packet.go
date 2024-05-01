package tax

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/mmtaee/moadian/internal/encryption"
	"github.com/mmtaee/moadian/internal/normalizer"
	"log"
)

type Packet struct {
	Uid             string `json:"uid"`
	PacketType      string `json:"packetType"`
	Retry           string `json:"retry"`
	Data            any    `json:"data"`
	EncryptionKeyId string `json:"encryptionKeyId"`
	SymmetricKey    string `json:"symmetricKey"`
	IV              string `json:"iv"`
	FiscalID        string `json:"fiscalID"`
	DataSignature   string `json:"dataSignature"`
}

type PacketMethods interface {
	ToMap() map[string]interface{}
	JsonData() []byte
	Normalize() string
	Sign(key *rsa.PrivateKey) ([]byte, string)
	Verify(string, []byte, *rsa.PrivateKey) bool
}

// ToMap function create a map from packet
func (p *Packet) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"uid":             p.Uid,
		"packetType":      p.PacketType,
		"retry":           p.Retry,
		"data":            p.Data,
		"encryptionKeyId": p.EncryptionKeyId,
		"symmetricKey":    p.SymmetricKey,
		"iv":              p.IV,
		"fiscalId":        p.FiscalID,
		"dataSignature":   p.DataSignature,
	}
	return result
}

// JsonData function create json loads from packet data
func (p *Packet) JsonData() []byte {
	data, err := json.Marshal(p.Data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

// Normalize function create sorted key joined with #
func (p *Packet) Normalize() string {
	var (
		header     map[string]interface{}
		body       []map[string]interface{}
		payments   []map[string]interface{}
		extensions []map[string]interface{}
	)
	if h, ok := p.Data.(map[string]interface{})["header"]; ok {
		header = h.(map[string]interface{})
	}
	if b, ok := p.Data.(map[string]interface{})["body"]; ok {
		body = b.([]map[string]interface{})
	}
	if py, ok := p.Data.(map[string]interface{})["payments"]; ok {
		payments = py.([]map[string]interface{})
	}
	if e, ok := p.Data.(map[string]interface{})["extensions"]; ok {
		extensions = e.([]map[string]interface{})
	}
	return normalizer.Normalize(header, body, payments, extensions)
}

// Sign function return normalized packet signed cipherBytes, cipher
func (p *Packet) Sign(privateKey *rsa.PrivateKey) ([]byte, string) {
	normalized := p.Normalize()
	return encryption.Sign(normalized, privateKey)
}

// Verify function check if sign is valid
func (p *Packet) Verify(normalized string, cipherBytes []byte, privateKey *rsa.PrivateKey) bool {
	return encryption.Verify(normalized, cipherBytes, privateKey)
}
