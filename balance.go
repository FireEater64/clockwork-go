package clockwork

import (
	"encoding/xml"
	"github.com/shopspring/decimal"
)

type AccountType int

const (
	PayAsYouGo AccountType = iota
	Invoice
)

type Balance struct {
	XMLName        xml.Name        `xml:"Balance_Resp"`
	Balance        decimal.Decimal `xml:"Balance"`
	CurrencySymbol string          `xml:"Currency>Symbol"`
	CurrencyCode   string          `xml:"Currency>Code"`
	AccountType    string          `xml:"AccountType"`
}
