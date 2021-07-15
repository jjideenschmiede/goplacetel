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