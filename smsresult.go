package clockwork

type SMSResult struct {
	To           string `xml:"To"`
	MessageID    string `xml:"MessageID"`
	ClientID     string `xml:"ClientID"`
	ErrorCode    int    `xml:"ErrNo"`
	ErrorMessage string `xml:"ErrDesc"`
}
