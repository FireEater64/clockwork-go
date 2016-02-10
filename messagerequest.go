package clockwork

import "encoding/xml"

type MessageRequest struct {
	XMLName xml.Name `xml:"Message"`
	Key     string   `xml:"Key"`
	SMS     []SMS    `xml:"SMS"`
}
