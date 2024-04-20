package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var s1 School

func main() {

	s1.Classes = append(s1.Classes, Class{
		Name: "12th",
		Students: []Student{{
			Name:   "Pawan",
			Age:    "2 months",
			RollNo: 12,
			Parents: Parents{
				MotherName: "Kamlesh",
				FatherName: "Sombir",
				ContactDetails: []ContactDetails{{
					Addresses: []Address{
						{
							HNO:         12,
							AddressType: AddressTypePermanent,
						},
					},
					Mobiles: []MobileNumbers{
						{
							Number: []string{
								"7988323110",
							},
						},
					},
				}},
			},
		}},
	})

	// data, err := json.MarshalIndent(s1, "", "   ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(data))

	http.HandleFunc("/home", ShowSchoolDetails)
	http.ListenAndServe(":8090", nil)
}

func ShowSchoolDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := json.NewEncoder(w).Encode(s1)
		if err != nil { 
			log.Fatalln(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request method not allowed !!"))
	}

}

type School struct {
	Classes []Class `json:"classes"`
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

type Student struct {
	Name    string  `json:"name"`
	Age     string  `json:"age"`
	RollNo  int     `json:"roll_no"`
	Parents Parents `json:"parents"`
}

type Parents struct {
	MotherName     string           `json:"mother_name"`
	FatherName     string           `json:"father_name"`
	ContactDetails []ContactDetails `json:"contact_details"`
}

type ContactDetails struct {
	Addresses []Address       `json:"addresses"`
	Mobiles   []MobileNumbers `json:"mobiles"`
}

type Address struct {
	HNO         int    `json:"hno"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PinCode     string `json:"pin_code,omitempty"`
	AddressType string `json:"address_type,omitempty"`
}
type MobileNumbers struct {
	Number []string `json:"number"`
}

const (
	AddressTypeCurrent   = "Current"
	AddressTypePermanent = "Parmenent"
)
