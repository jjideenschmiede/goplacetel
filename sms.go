//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of placetel.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goplacetel

import (
	"encoding/json"
)

// SendSmsBody is to build the request body
type SendSmsBody struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type SendSmsReturn struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

// SendSms is to send a sms via the placetel api
func SendSms(body SendSmsBody, token string) (SendSmsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return SendSmsReturn{}, err
	}

	// Set config for new request
	r := Request{"/sms", "POST", token, convert}

	// Send new request
	response, err := r.Send()
	if err != nil {
		return SendSmsReturn{}, err
	}

	// Close response body after function ends
	defer response.Body.Close()

	// Decode data
	var decode SendSmsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SendSmsReturn{}, err
	}

	// Return data
	return decode, nil

}
