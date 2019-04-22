# MoesifMiddleware Example for Golang

Send REST API Events to Moesif for error analysis

[Source Code on GitHub](https://github.com/moesif/moesifmiddleware-go-example)

## How to run this example

1. Install Moesif Middleware if you have not done so. `go get github.com/moesif/moesifmiddleware-go`

2. Be sure to edit the moesif_options/moesif_options.go to change the application id to your application id obtained from Moesif.

3. Start the server - `go run main.go`

4. See main.go for some urls that you can hit the server with (e.g. http://localhost:3000/api/employee/42), and the data should be captured in the corresponding Moesif account of the application id.

## How to test capture outgoing request

1. Install Moesif Middleware if you have not done so. `go get github.com/moesif/moesifmiddleware-go`

2. Be sure to edit the moesif_options/moesif_options.go to change the application id to your `application_id` obtained from Moesif and set `capture_outgoing_request` to true to capture outgoing request.

3. Switch to `moesif_capture_outgoing` directory and run the test `go test -v` - and the data (incoming and outgoing api call) should be captured in the corresponding Moesif account of the application id.

## Other integrations

To view more more documentation on integration options, please visit __[the Integration Options Documentation](https://www.moesif.com/docs/getting-started/integration-options/).__
