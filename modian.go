package moadian

import (
	"crypto/rsa"
	"github.com/mmtaee/moadian/internal/encryption"
	"github.com/mmtaee/moadian/internal/tax_id"
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

func PrivateKeyFromPath(path string) (*rsa.PrivateKey, error) {
	return encryption.LoadPrivateKey(path)
}

func New(privateKey *rsa.PrivateKey, fiscalID string, economicCode string) *Service {
	return &Service{
		FiscalID:     fiscalID,
		EconomicCode: economicCode,
		PrivateKey:   privateKey,
		TaxUrl:       "https://tp.tax.gov.ir/req/api/self-tsp",
		Version:      "01",
		Priority:     "normal-enqueue",
	}
}

func (s *Service) GenerateTaxId(serial int, timestamp ...int64) string {
	var t int64
	if len(timestamp) > 0 {
		t = timestamp[0]
	}
	return tax_id.GenerateTaxID(s.FiscalID, serial, t)
}
