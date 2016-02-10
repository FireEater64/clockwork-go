package clockwork

import (
	"encoding/xml"
	"testing"
)

func TestBalance_BasicBalanceMessageDeserializesCorrectly(t *testing.T) {
	testXML := `<?xml version="1.0" encoding="utf-8"?>
<Balance_Resp>
  <AccountType>PAYG</AccountType>
  <Balance>49.95</Balance>
  <Currency>
    <Code>GBP</Code>
    <Symbol>£</Symbol>
  </Currency>
</Balance_Resp>`

	result := Balance{}
	err := xml.Unmarshal([]byte(testXML), &result)

	if err != nil {
		t.Errorf("Expected no error - received: %s", err.Error())
		t.Fail()
	}

	if result.AccountType != "PAYG" {
		t.Errorf("Expected AccountType - PAYG. Received - %s", result.AccountType)
		t.Fail()
	}

	// TODO: Lazy Decimal comparison
	if result.Balance.String() != "49.95" {
		t.Errorf("Expected Balance - 49.95. Received - %s", result.Balance.String())
		t.Fail()
	}

	if result.CurrencyCode != "GBP" {
		t.Errorf("Expected CurrencyCode - GBP. Received - %s", result.CurrencyCode)
		t.Fail()
	}

	if result.CurrencySymbol != "£" {
		t.Errorf("Expected CurrencySymbol - GBP. Received - %s", result.CurrencySymbol)
		t.Fail()
	}
}
