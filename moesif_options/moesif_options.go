package moesif_options

import (
	"net/http"
	"strings"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	"github.com/moesif/moesifapi-go/models"
)

// Mask Event Model
func maskEventModel(eventModel models.EventModel) models.EventModel {
	return eventModel
}

// Set User Id
func identifyUser(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) string{
	return "golangapiuser"
}

// Set Session Token
func getSessionToken(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) string{
	return "token is blah blah blah"
}

// Skip Event
func shouldSkip(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) bool{
	return strings.Contains(request.RequestURI, "test")
}

// Set Metadata
func getMetadata(request *http.Request, response moesifmiddleware.MoesifResponseRecorder) map[string]interface{} {
	
	var innerNestedFields = map[string] interface{} {
		"nestedInner": "test",
	}

	var nestedFields = map[string] interface{} {
		"inner":  innerNestedFields,
	}
	
	var metadata = map[string]interface{} {
		"foo" : "bar",
		"user": "golangapiuser",
		"test": nestedFields,
	}
	return metadata
}

func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{} {
		"Application_Id": "Moesif Application Id",
		"Api_Version": "1.0.0",
		"Get_Metadata": getMetadata,
		"Should_Skip": shouldSkip,
		"Identify_User": identifyUser,
		"Get_Session_Token": getSessionToken,
		"Mask_Event_Model": maskEventModel,
		"Debug": true,
	}
	return moesifOptions
}
