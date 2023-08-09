package moesif_options

import (
	"net/http"
	"strings"

	"github.com/moesif/moesifapi-go/models"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
)

func maskRequestHeader() []string {
	return []string{"reqHeader"}
}

func maskRequestBody() []string {
	return []string{"reqBody", "password"}
}

func maskResponseHeader() []string {
	return []string{"rspHeader"}
}

func maskResponseBody() []string {
	return []string{"id", "rspBody"}
}

// Set User Id
func identifyUser(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) string {
	return request.Header.Get("X-User-Id")
}

// Set Company Id
func identifyCompany(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) string {
	return request.Header.Get("X-Company-Id")
}

// Set Session Token
func getSessionToken(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) string {
	return "token is blah blah blah"
}

// Skip Event
func shouldSkip(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) bool {
	return strings.Contains(request.RequestURI, "test")
}

// Set Metadata
func getMetadata(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) map[string]interface{} {

	var innerNestedFields = map[string]interface{}{
		"nestedInner": "test",
	}

	var nestedFields = map[string]interface{}{
		"inner": innerNestedFields,
	}

	var metadata = map[string]interface{}{
		"foo":  "bar",
		"user": "golangapiuser",
		"test": nestedFields,
	}
	return metadata
}

// Skip Outgoing Event
func shouldSkipOutgoing(request *http.Request, response *http.Response) bool {
	return strings.Contains(request.URL.String(), "test")
}

// Set Outgoing Event User Id
func identifyUserOutgoing(request *http.Request, response *http.Response) string {
	return "golangapiuserOutgoing"
}

// Set Outgoing Event Company Id
func identifyCompanyOutgoing(request *http.Request, response *http.Response) string {
	return "golangapicompanyOutgoing"
}

// Set Outgoing Event Session Token
func getSessionTokenOutgoing(request *http.Request, response *http.Response) string {
	return "token is blah blah blah"
}

// Mask Outgoing Event Model
func maskEventModelOutgoing(eventModel models.EventModel) models.EventModel {
	return eventModel
}

// Set Outoing Event Metadata
func getMetadataOutgoing(request *http.Request, response *http.Response) map[string]interface{} {

	var innerNestedFields = map[string]interface{}{
		"nestedInner": "test",
	}

	var nestedFields = map[string]interface{}{
		"inner": innerNestedFields,
	}

	var metadata = map[string]interface{}{
		"foo":  "bar",
		"user": "golangapiuser",
		"test": nestedFields,
	}
	return metadata
}

func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{}{
		"Application_Id":             "Your Moesif Application Id",
		"Api_Version":                "1.0.0",
		"Get_Metadata":               getMetadata,
		"Should_Skip":                shouldSkip,
		"Identify_User":              identifyUser,
		"Identify_Company":           identifyCompany,
		"Get_Session_Token":          getSessionToken,
		"Request_Header_Masks":       maskRequestHeader,
		"Request_Body_Masks":         maskRequestBody,
		"Response_Header_Masks":      maskResponseHeader,
		"Response_Body_Masks":        maskResponseBody,
		"Debug":                      true,
		"Log_Body":                   true,
		"Log_Body_Outgoing":          true,
		"Should_Skip_Outgoing":       shouldSkipOutgoing,
		"Identify_User_Outgoing":     identifyUserOutgoing,
		"Identify_Company_Outgoing":  identifyCompanyOutgoing,
		"Get_Metadata_Outgoing":      getMetadataOutgoing,
		"Get_Session_Token_Outgoing": getSessionTokenOutgoing,
		"Mask_Event_Model_Outgoing":  maskEventModelOutgoing,
	}
	return moesifOptions
}
