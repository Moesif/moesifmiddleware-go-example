package main

import(
	"log"
	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
	options "github.com/moesif/moesifmiddleware-go-example/moesif_options"
	"net/http"
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

var moesifOption map[interface{}]string

func init() {
	moesifOption = options.MoesifOptions()
}

type Employee struct{
	DateOfBirth      *time.Time 		`json:"date_of_birth" form:"date_of_birth"`			//Time when request was made
	Id				 int				`json:"id" form:"id"`                               //HTTP Status code such as 200
	FirstName		 string				`json:"first_name" form:"first_name"`               //verb of the API request such as GET or POST
	LastName		 string				`json:"last_name" form:"last_name"`                 //verb of the API request such as GET or POST
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

func main() {
	http.Handle("/api/employee/", moesifmiddleware.MoesifMiddleware(http.HandlerFunc(handle), moesifOption))
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
		Id: id,
		FirstName: "firstName",
		LastName: "lastName",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
