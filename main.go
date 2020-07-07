package main

import (
	"encoding/json"
	models "github.com/moesif/moesifapi-go/models"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	options "github.com/moesif/moesifmiddleware-go-example/moesif_options"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

var moesifOption map[string]interface{}

func init() {
	moesifOption = options.MoesifOptions()
}

type Employee struct {
	DateOfBirth *time.Time `json:"date_of_birth" form:"date_of_birth"` //Time when request was made
	Id          int        `json:"id" form:"id"`                       //HTTP Status code such as 200
	FirstName   string     `json:"first_name" form:"first_name"`       //verb of the API request such as GET or POST
	LastName    string     `json:"last_name" form:"last_name"`         //verb of the API request such as GET or POST
}

func ParseID(s string) (id int, err error) {
	p := strings.LastIndex(s, "/")
	if p < 0 {
		return 0, nil
	}

	first := s[:p+1]
	if first != "/api/employee/" {
		return 0, nil
	}

	id, err = strconv.Atoi(s[p+1:])
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func literalFieldValue(value string) *string {
	return &value
}

func main() {
	http.Handle("/api/employee/", moesifmiddleware.MoesifMiddleware(http.HandlerFunc(handle), moesifOption))
	http.Handle("/api/users/", moesifmiddleware.MoesifMiddleware(http.HandlerFunc(Usershandle), moesifOption))
	http.Handle("/api/companies/", moesifmiddleware.MoesifMiddleware(http.HandlerFunc(Companieshandle), moesifOption))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	time := time.Now().UTC().AddDate(-30, 0, 0)
	id, _ := ParseID(r.URL.Path)
	var employee = Employee{
		DateOfBirth: &time,
		Id:          id,
		FirstName:   "firstName",
		LastName:    "lastName",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func Usershandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	// Campaign object is optional, but useful if you want to track ROI of acquisition channels
	// See https://www.moesif.com/docs/api#users for campaign schema
	campaign := models.CampaignModel{
		UtmSource:   literalFieldValue("google"),
		UtmMedium:   literalFieldValue("cpc"),
		UtmCampaign: literalFieldValue("adwords"),
		UtmTerm:     literalFieldValue("api+tooling"),
		UtmContent:  literalFieldValue("landing"),
	}

	// metadata can be any custom dictionary
	metadata := map[string]interface{}{
		"email":      "john@acmeinc.com",
		"first_name": "John",
		"last_name":  "Doe",
		"title":      "Software Engineer",
		"sales_info": map[string]interface{}{
			"stage":          "Customer",
			"lifetime_value": 24000,
			"account_owner":  "mary@contoso.com",
		},
	}

	userId := path.Base(r.URL.Path)

	// Only UserId is required
	user := models.UserModel{
		UserId:    userId,
		CompanyId: literalFieldValue("67890"), // If set, associate user with a company object
		Campaign:  &campaign,
		Metadata:  &metadata,
	}

	moesifmiddleware.UpdateUser(&user, moesifOption)
}

func Companieshandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	// Campaign object is optional, but useful if you want to track ROI of acquisition channels
	// See https://www.moesif.com/docs/api#update-a-company for campaign schema
	campaign := models.CampaignModel{
		UtmSource:   literalFieldValue("google"),
		UtmMedium:   literalFieldValue("cpc"),
		UtmCampaign: literalFieldValue("adwords"),
		UtmTerm:     literalFieldValue("api+tooling"),
		UtmContent:  literalFieldValue("landing"),
	}

	// metadata can be any custom dictionary
	metadata := map[string]interface{}{
		"org_name":   "Acme, Inc",
		"plan_name":  "Free",
		"deal_stage": "Lead",
		"mrr":        24000,
		"demographics": map[string]interface{}{
			"alexa_ranking":  500000,
			"employee_count": 47,
		},
	}

	companyId := path.Base(r.URL.Path)

	// Prepare company model
	company := models.CompanyModel{
		CompanyId:     companyId,                        // The only required field is your company id
		CompanyDomain: literalFieldValue("acmeinc.com"), // If domain is set, Moesif will enrich your profiles with publicly available info
		Campaign:      &campaign,
		Metadata:      &metadata,
	}

	moesifmiddleware.UpdateCompany(&company, moesifOption)
}
