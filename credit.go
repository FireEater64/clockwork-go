package clockwork

import (
	"encoding/xml"
)

type Credit struct {
	XMLName xml.Name `xml:"Credit_Resp"`
	Credit  int      `xml:"Credit"`
}
