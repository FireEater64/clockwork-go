package clockwork

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	smsURL        string = "https://api.clockworksms.com/xml/send.aspx"
	creditURL     string = "https://api.clockworksms.com/xml/credit.aspx"
	balanceURL    string = "https://api.clockworksms.com/xml/balance.aspx"
	xmlTypeHeader string = "text/xml"
)

type Clockwork struct {
	Key string
	SSL bool
}

func (clockwork *Clockwork) GetBalance() *Balance {
	request := balanceRequest{Key: clockwork.Key}
	balanceToReturn := Balance{}
	getXmlResponse(balanceURL, request, &balanceToReturn)
	return &balanceToReturn
}

func (clockwork *Clockwork) GetCredit() *Credit {
	request := creditRequest{Key: clockwork.Key}
	creditToReturn := Credit{}
	getXmlResponse(creditURL, request, &creditToReturn)
	return &creditToReturn
}

func (clockwork *Clockwork) SendSMS(givenSMS SMS) *MessageResponse {
	smsArray := []SMS{givenSMS}
	return clockwork.SendMultipleSMS(smsArray)
}

func (clockwork *Clockwork) SendMultipleSMS(givenSMS []SMS) *MessageResponse {
	request := MessageRequest{Key: clockwork.Key, SMS: givenSMS}
	messageResponseToReturn := MessageResponse{}
	getXmlResponse(smsURL, request, &messageResponseToReturn)
	return &messageResponseToReturn
}

// TODO: Should be a slightly more restricted interface?
func getXmlResponse(url string, request interface{}, resultToFill interface{}) {
	requestXml, marshalErr := xml.Marshal(request)
	if marshalErr != nil {
		panic(marshalErr)
	}

	resp, postErr := http.Post(url, xmlTypeHeader, bytes.NewBuffer(requestXml))
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

func NewClockwork(Key string) *Clockwork {
	return &Clockwork{Key: Key, SSL: true}
}
