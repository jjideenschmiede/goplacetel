//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of placetel.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package placetel

import (
	"net/http"
	"strings"
)

// Send sms
func SendSMS(recipient, message, token string) error {

	// Define client
	client := &http.Client{}

	// Set body
	body := strings.NewReader(`{"recipient": "` + recipient + `", "message": "` + message + `"}`)

	// Set request
	request, err := http.NewRequest("POST", "https://api.placetel.de/v2/sms", body)
	if err != nil {
		return err
	}

	// Define header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	// Send request to placetel
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// Close response after function ends
	defer response.Body.Close()

	// Return
	return nil

}
