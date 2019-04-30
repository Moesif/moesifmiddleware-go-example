package capture_outgoing_test

import (
	"log"
	"net/http"
	"time"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	options "github.com/moesif/moesifmiddleware-go-example/moesif_options"
	"testing"
)

var moesifOption map[string]interface{}

// Function to initialize the options
func init() {
	moesifOption = options.MoesifOptions()
}

func TestCaptureOutgoing(t *testing.T) {

	// Start Capturing Outgoing Request
	moesifmiddleware.StartCaptureOutgoing(moesifOption)

	// Outgoing API call to third parties like Github / Stripe or to your own dependencies
	_, err := http.Get("https://api.github.com")

	// Check for any errors while sending outgoing request
	if err != nil {
		log.Printf("Error while sending request : %s.\n", err.Error())
	}

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)

}
