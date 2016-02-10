package clockwork

import "encoding/xml"

type MessageResponse struct {
	XMLName   xml.Name    `xml:"Message_Resp"`
	SMSResult []SMSResult `xml:"SMS_Resp"`
	ErrNo     int         `xml:"ErrNo"`
	ErrDesc   int         `xml:"ErrDesc"`
}
