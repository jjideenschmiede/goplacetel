//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goplacetel.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goplacetel

import (
	"encoding/json"
	"fmt"
)

// ContactBody is to build a new contact
type ContactBody struct {
	Id             int    `json:"id,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	EmailWork      string `json:"email_work"`
	Company        string `json:"company"`
	Address        string `json:"address"`
	AddressWork    string `json:"address_work"`
	PhoneWork      string `json:"phone_work"`
	MobileWork     string `json:"mobile_work"`
	Phone          string `json:"phone"`
	Mobile         string `json:"mobile"`
	Fax            string `json:"fax"`
	FaxWork        string `json:"fax_work"`
	FacebookUrl    string `json:"facebook_url"`
	LinkedinUrl    string `json:"linkedin_url"`
	XingUrl        string `json:"xing_url"`
	TwitterAccount string `json:"twitter_account"`
	Blocked        bool   `json:"blocked"`
}

// ContactReturn is to decode the json return
type ContactReturn struct {
	Id             int         `json:"id"`
	UserId         int         `json:"user_id"`
	Speeddail      interface{} `json:"speeddail"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	EmailWork      string      `json:"email_work"`
	Company        string      `json:"company"`
	Address        string      `json:"address"`
	AddressWork    string      `json:"address_work"`
	PhoneWork      string      `json:"phone_work"`
	MobileWork     string      `json:"mobile_work"`
	Phone          string      `json:"phone"`
	Mobile         string      `json:"mobile"`
	Fax            string      `json:"fax"`
	FaxWork        string      `json:"fax_work"`
	FacebookUrl    string      `json:"facebook_url"`
	LinkedinUrl    string      `json:"linkedin_url"`
	XingUrl        string      `json:"xing_url"`
	TwitterAccount string      `json:"twitter_account"`
	Blocked        bool        `json:"blocked"`
	UpdatedAt      string      `json:"updated_at"`
	CreatedAt      string      `json:"created_at"`
}

// DeleteContactReturn is to decode the json data
type DeleteContactReturn struct {
	Id             int         `json:"id"`
	UserId         int         `json:"user_id"`
	Speeddail      interface{} `json:"speeddail"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Number1        string      `json:"number1"`
	Number2        string      `json:"number2"`
	Number3        string      `json:"number3"`
	Number4        string      `json:"number4"`
	Number5        string      `json:"number5"`
	IsOnBlacklist  bool        `json:"isOnBlacklist"`
	Email          string      `json:"email"`
	MoreNumbers    interface{} `json:"more_numbers"`
	ClientId       int         `json:"client_id"`
	IsGlobal       bool        `json:"is_global"`
	Company        string      `json:"company"`
	Address        string      `json:"address"`
	ContactBookId  string      `json:"contact_book_id"`
	ImportFrom     interface{} `json:"import_from"`
	ImportId       interface{} `json:"import_id"`
	EmailWork      string      `json:"email_work"`
	AddressWork    string      `json:"address_work"`
	Fax            string      `json:"fax"`
	FaxWork        string      `json:"fax_work"`
	Color          interface{} `json:"color"`
	FacebookUrl    string      `json:"facebook_url"`
	LinkedinUrl    string      `json:"linkedin_url"`
	XingUrl        string      `json:"xing_url"`
	GoogleplusUrl  interface{} `json:"googleplus_url"`
	TwitterAccount string      `json:"twitter_account"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
}

// Contacts is to get all contacts from the api
func Contacts(token string) (*[]ContactReturn, error) {

	// To decode the json data
	var contacts []ContactReturn

	// To call the pages
	page := 1

	// Loop over all sites
	for {

		// Set config for new request
		r := Request{fmt.Sprintf("/contacts?page=%d&per_page=25", page), "GET", token, nil}

		// Send new request
		response, err := r.Send()
		if err != nil {
			return nil, err
		}

		// Decode data
		var decode []ContactReturn

		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return nil, err
		}

		// Close response body
		response.Body.Close()

		// Add contacts
		for _, value := range decode {
			contacts = append(contacts, value)
		}

		// Check length & break the loop
		if len(decode) < 25 {
			break
		}

	}

	// Return data
	return &contacts, nil

}

// AddContact is to add a new contact
func AddContact(body *ContactBody, token string) (*ContactReturn, error) {

	// Convert data
	convert, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Set config for new request
	r := Request{"/contacts", "POST", token, convert}

	// Send new request
	response, err := r.Send()
	if err != nil {
		return nil, err
	}

	// Close response body after function ends
	defer response.Body.Close()

	// Decode data
	var decode ContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}

// UpdateContact is to add a new contact
func UpdateContact(body *ContactBody, token string) (*ContactReturn, error) {

	// Convert data
	convert, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Set config for new request
	r := Request{fmt.Sprintf("/contacts/%d", body.Id), "PUT", token, convert}

	// Send new request
	response, err := r.Send()
	if err != nil {
		return nil, err
	}

	// Close response body after function ends
	defer response.Body.Close()

	// Decode data
	var decode ContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}

// DeleteContact is to add a new contact
func DeleteContact(id string, token string) (*DeleteContactReturn, error) {

	// Set config for new request
	r := Request{"/contacts/" + id, "DELETE", token, nil}

	// Send new request
	response, err := r.Send()
	if err != nil {
		return nil, err
	}

	// Close response body after function ends
	defer response.Body.Close()

	// Decode data
	var decode DeleteContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}
