# placetel

With this small library you should be able to address some nodes of the Placetel API. Please note that we have only implemented a tiny part of it and will continue to develop it as needed.

Feel free to help us with the development!

## Install

```console
go get github.com/jjideenschmiede/goplacetel
```

## How to use it?

### Send SMS

With the SMS sent you are able to send sms through your Placetel account. The whole thing works as follows:

```go
// Define body
body := &goplacetel.SendSmsBody{"0123456789", "Test message!"}

// Send sms
response, err := goplacetel.SendSms(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### Get all contacts

To get a list of all contacts, you must call the following function. This function returns all values in a struct.

```go
// Get all contacts
contacts, err := goplacetel.Contacts("token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contacts)
}
```

### Get a contact by id

Damit man einen speziellen Kontakt auslesen kann, wird die ID benötigt. Diese erhalten Sie mit der Funktion goplacetel.Contacts().

```go
// Get a contact by id
contact, err := Contact("id", "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contact)
}
```

### Add a contact

If a new contact is to be created, then this can be done as follows. The field: id (it is the first field) can be ignored.

```go
// Define body
body := &goplacetel.ContactBody{
    0,
    "first_name",
    "last_name",
    "email",
    "email_work",
    "company",
    "address",
    "address_work",
    "phone_work",
    "mobile_work",
    "phone",
    "mobile",
    "fax",
    "fax_work",
    "facebook_url",
    "linkedin_url",
    "xing_url",
    "twitter_account",
    false,
}

// Add a new contact
contact, err := goplacetel.AddContact(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contact)
}
```

### Update a contact

This function is identical to the AddContact function. With the difference that the Id must be specified.

```go
// Define body
body := &goplacetel.ContactBody{
    id,
    "first_name",
    "last_name",
    "email",
    "email_work",
    "company",
    "address",
    "address_work",
    "phone_work",
    "mobile_work",
    "phone",
    "mobile",
    "fax",
    "fax_work",
    "facebook_url",
    "linkedin_url",
    "xing_url",
    "twitter_account",
    false,
}

// Update a contact
contact, err := goplacetel.UpdateContact(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contact)
}
```

### Remove a contact

To be able to remove a contact, the ID is needed. You get this with the function goplacetel.Contacts().

```go
// Remove a contact
delete, err := goplacetel.DeleteContact("id", "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(delete)
}
```