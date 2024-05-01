package moadian

import (
	"crypto/rsa"
	"github.com/mmtaee/moadian/internal/encryption"
	"github.com/mmtaee/moadian/internal/taxid"
	"github.com/mmtaee/moadian/pkg/services/tax"
	"log"
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

type Services interface {
	GenerateTaxId(int, int64) string
	GetServerInfo()
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

func NewFromPath(privateKeyPth string, fiscalID string, economicCode string) *Service {
	privateKey, err := encryption.LoadPrivateKey(privateKeyPth)
	if err != nil {
		log.Fatalln(err)
	}
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
	return taxid.GenerateTaxID(s.FiscalID, serial, t)
}

func (s *Service) GetServerInfo() (interface{}, int, error) {
	result, status, err := tax.ServerInfo(s.TaxUrl)
	if err != nil {
		s.TaxPublicKey = result.Result.Data.PublicKeys[0].Key
		s.TaxServerKeyID = result.Result.Data.PublicKeys[0].ID
	}
	return result, status, err
}
