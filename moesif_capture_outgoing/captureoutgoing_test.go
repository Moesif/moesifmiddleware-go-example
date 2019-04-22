package capture_outgoing_test

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"io"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	options "github.com/moesif/moesifmiddleware-go-example/moesif_options"
	"testing"
	"net/http/httptest"
)

var moesifOption map[string]interface{}

// Function to initialize the options
func init() {
	moesifOption = options.MoesifOptions()
}

// Pretty print function
func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func TestCaptureOutgoing(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
	req := httptest.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)

	// Response recorder to record the response
	rr := httptest.NewRecorder()

	// Moesif Middleware function - func MoesifMiddleware(next http.Handler, configurationOption map[string]interface{})
	handler := moesifmiddleware.MoesifMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write the response
		io.WriteString(w, `{
			"userId": 1,
			"id": 1,
			"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
		  }`)
		// Set the content-type as application/json
		w.Header().Set("Content-Type", "application/json")

		// Outgoing API call to third parties like Github / Stripe or to your own dependencies
		_, err := http.Get("https://api.github.com")

		// Check for any errors while sending outgoing request
		if err != nil {
			log.Printf("Error while sending request : %s.\n", err.Error())
		}

	}), moesifOption)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)

	// Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
