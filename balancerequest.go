package clockwork

import (
	"encoding/xml"
)

type balanceRequest struct {
	XMLName xml.Name `xml:"Balance"`
	Key     string   `xml:"Key"`
}
