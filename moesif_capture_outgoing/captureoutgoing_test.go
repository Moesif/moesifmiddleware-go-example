package capture_outgoing_test

import (
	"log"
	"net/http"
	"time"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	"github.com/moesif/moesifapi-go/models"
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

	log.Printf("Waiting for the queue to flush")

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)

}

func TestUpdateUser(t *testing.T) {
	
	// Modified Time
	modifiedTime := time.Now().UTC()

	// User Metadata
	metadata := map[string]interface{}{
		"email": "johndoe1@acmeinc.com",
		"Key1": "metadata",
		"Key2": 42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	// Prepare user model
	user := models.UserModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		UserId:			  "golangapiuser",	
		UserAgentString:  nil,
		Metadata:		  &metadata,
	}

	// Update User
	moesifmiddleware.UpdateUser(&user, moesifOption)

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)
}

func TestUpdateUsersBatch(t *testing.T) {

	// Batch Users
	var users []*models.UserModel

	// Modified Time
	modifiedTime := time.Now().UTC()

	// User Metadata
	metadata := map[string]interface{}{
		"email": "johndoe1@acmeinc.com",
		"Key1": "metadata",
		"Key2": 42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	// Prepare user model
	userA := models.UserModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		UserId:			  "golangapiuser",	
		UserAgentString:  nil,
		Metadata:		  &metadata,
	}

	userB := models.UserModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		UserId:			  "golangapiuser1",	
		UserAgentString:  nil,
		Metadata:		  &metadata,
	}

	users = append(users, &userA)
	users = append(users, &userB)

	// Update User
	moesifmiddleware.UpdateUsersBatch(users, moesifOption)

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)
}

func TestUpdateCompany(t *testing.T) {
	
	// Modified Time
	modifiedTime := time.Now().UTC()

	// User Metadata
	metadata := map[string]interface{}{
		"email": "johndoe1@acmeinc.com",
		"Key1": "metadata",
		"Key2": 42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	// Prepare company model
	company := models.CompanyModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		CompanyId:		  "1",	
		CompanyDomain:    nil,
		Metadata:		  &metadata,
	}

	// Update company
	moesifmiddleware.UpdateCompany(&company, moesifOption)

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)
}

func TestUpdateCompaniesBatch(t *testing.T) {

	// Batch Companies
	var companies []*models.CompanyModel

	// Modified Time
	modifiedTime := time.Now().UTC()

	// Company Metadata
	metadata := map[string]interface{}{
		"email": "johndoe1@acmeinc.com",
		"Key1": "metadata",
		"Key2": 42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	// Prepare company model
	companyA := models.CompanyModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		CompanyId:		  "1",	
		CompanyDomain:    nil,
		Metadata:		  &metadata,
	}

	companyB := models.CompanyModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		CompanyId:		  "2",	
		CompanyDomain:    nil,
		Metadata:		  &metadata,
	}

	companies = append(companies, &companyA)
	companies = append(companies, &companyB)

	// Update Companies
	moesifmiddleware.UpdateCompaniesBatch(companies, moesifOption)

	// Sleep to allow queue to flush for testing purpose
	time.Sleep(20*time.Second)
}
