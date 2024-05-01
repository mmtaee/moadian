package moadian

import (
	"crypto/rsa"
	"time"
)

type Service struct {
	FiscalID           string
	EconomicCode       string
	PrivateKey         *rsa.PrivateKey
	TaxToken           string
	TaxTokenExpireTime *time.Time
	TaxPublicKey       string
	TaxServerKeyID     string
	TaxUrl             string
	Version            string
	Priority           string
}

func NewService(privateKey *rsa.PrivateKey, fiscalID string, economicCode string) *Service {
	return &Service{
		FiscalID:     fiscalID,
		EconomicCode: economicCode,
		PrivateKey:   privateKey,
		TaxUrl:       "https://tp.tax.gov.ir/req/api/self-tsp",
		Version:      "01",
		Priority:     "normal-enqueue",
	}
}
