package clockwork

import (
	"encoding/xml"
)

type creditRequest struct {
	XMLName xml.Name `xml:"Balance"`
	Key     string   `xml:"Key"`
}
