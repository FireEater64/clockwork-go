package clockwork

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	baseURL       string = "api.clockworksms.com"
	smsSuffix     string = "/xml/send.aspx"
	creditSuffix  string = "/xml/credit.aspx"
	balanceSuffix string = "/xml/balance.aspx"
	xmlTypeHeader string = "text/xml"
)

// Clockwork is the main class used to communicate with the Clockwork API
type Clockwork struct {
	Key string
	SSL bool
}

// GetBalance returns number of SMS credits available to the
// account associated with this Clockwork API key.
// Deprecated: Use the GetBalance function to return a cash balance
func (clockwork *Clockwork) GetBalance() *Balance {
	request := balanceRequest{Key: clockwork.Key}
	balanceToReturn := Balance{}
	getXMLResponse(buildURL(balanceSuffix, clockwork.SSL), request, &balanceToReturn)
	return &balanceToReturn
}

// GetCredit returns the current cash balance and currency information for
// the account associated with this Clockwork API key.
func (clockwork *Clockwork) GetCredit() *Credit {
	request := creditRequest{Key: clockwork.Key}
	creditToReturn := Credit{}
	getXMLResponse(buildURL(creditSuffix, clockwork.SSL), request, &creditToReturn)
	return &creditToReturn
}

// SendSMS sends a single SMS message
func (clockwork *Clockwork) SendSMS(givenSMS SMS) *MessageResponse {
	smsArray := []SMS{givenSMS}
	return clockwork.SendMultipleSMS(smsArray)
}

// SendMultipleSMS sends an array of SMS messages
func (clockwork *Clockwork) SendMultipleSMS(givenSMS []SMS) *MessageResponse {
	request := MessageRequest{Key: clockwork.Key, SMS: givenSMS}
	messageResponseToReturn := MessageResponse{}
	getXMLResponse(buildURL(smsSuffix, clockwork.SSL), request, &messageResponseToReturn)
	return &messageResponseToReturn
}

func getXMLResponse(url string, request interface{}, resultToFill interface{}) {
	requestXML, marshalErr := xml.Marshal(request)
	if marshalErr != nil {
		panic(marshalErr)
	}

	resp, postErr := http.Post(url, xmlTypeHeader, bytes.NewBuffer(requestXML))
	if postErr != nil {
		panic(postErr)
	}

	defer resp.Body.Close()
	xmlResponse, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}

	unmarshalErr := xml.Unmarshal(xmlResponse, &resultToFill)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
}

func buildURL(suffix string, SSL bool) string {
	if SSL {
		return "https://" + baseURL + suffix
	}

	return "http://" + baseURL + suffix
}

// NewClockwork returns a new instance of the Clockwork client, using the given
// API key
func NewClockwork(Key string) *Clockwork {
	return &Clockwork{Key: Key, SSL: true}
}
