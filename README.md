# Moesif Go Example

[Moesif](https://www.moesif.com) is an API analyatics and monitoring platform.
This is an example API built on Go with Moesif integrated. 

[Source Code on GitHub](https://github.com/moesif/moesifmiddleware-go-example)

## How to add middleware to your application

Add middleware to your application.

```go
http.Handle(pattern string, moesifmiddleware.MoesifMiddleware(http.HandlerFunc(handle), moesifOption))
```

#### handler func(ResponseWriter, *Request)
(__required__), HandlerFunc registers the handler function for the given pattern.

#### moesifOption
(__required__), _map[string]interface{}_, are the configuration options for your application. Please find more details on how to [configure options](https://github.com/Moesif/moesifmiddleware-go#configuration-options).

## How to run this example

1. Install Moesif Middleware if you have not done so. `go get github.com/moesif/moesifmiddleware-go`

2. Be sure to edit the moesif_options/moesif_options.go to add your Moesif application id.

```go
func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{} {
		"Application_Id": "Moesif Application Id",
		"Log_Body": true,
	}
	return moesifOptions
}
```
Your Moesif Application Id can be found in the [_Moesif Portal_](https://www.moesif.com/).
After signing up for a Moesif account, your Moesif Application Id will be displayed during the onboarding steps. 

You can always find your Moesif Application Id at any time by logging 
into the [_Moesif Portal_](https://www.moesif.com/), click on the top right menu,
and then clicking _Installation_.

3. Start the server:

```bash
go run main.go
```

4. See main.go for some urls that you can hit the server with (e.g. http://localhost:3000/api/employee/42), and the data should be captured in the corresponding Moesif account of the application id.

## How to test capture outgoing request

1. Install Moesif Middleware if you have not done so. `go get github.com/moesif/moesifmiddleware-go`

2. Be sure to edit the moesif_options/moesif_options.go to change the application id to your `application_id` obtained from Moesif and set `capture_outgoing_request` to true to capture outgoing request.

```go
func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{} {
        "Application_Id": "Moesif Application Id",
		"Capture_Outoing_Requests": true,
		"Log_Body": true,
	}
	return moesifOptions
}
```

3. Switch to `moesif_capture_outgoing` directory and run the test `go test -run TestCaptureOutgoing` - and the data (incoming and outgoing api call) should be captured in the corresponding Moesif account of the application id.

## Other integrations

To view more documentation on integration options, please visit __[the Integration Options Documentation](https://www.moesif.com/docs/getting-started/integration-options/).__
