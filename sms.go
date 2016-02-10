package clockwork

type InvalidCharacterAction int

const (
	AccountDefault InvalidCharacterAction = iota
	Error
	Remove
	Replace
)

type SMS struct {
	To                     string                 `xml:"To"`
	Message                string                 `xml:"Content"`
	From                   string                 `xml:"From"`
	ClientID               string                 `xml:"ClientID"`
	Truncate               bool                   `xml:"Truncate"`
	InvalidCharacterAction InvalidCharacterAction `xml:"InvalidCharAction"`
}
