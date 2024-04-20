package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type machine struct {
	AllCompanyName []string `json:"all_company_name"`
	User           User     `json:"user"`
}

type User struct {
	UserName      string  `json:"user_name"`
	ModleYear     string  `json:"model_year"`
	PurchageDate  string  `json:"purchase_date"`
	IsWarnty      bool    `json:"is_warranty"`
	Address       Address `json:"address"`
	MobileNumbers string  `json:"mobile_numbers"`
}

type Address struct {
	HouseNo  int    `json:"house_no"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
}

var allData = []machine{

	{AllCompanyName: []string{"LG", "Videocone", "Philips", "Godrej"},

		User: User{
			UserName:      "Pawan",
			ModleYear:     "2023",
			PurchageDate:  "22-Dec-2023",
			IsWarnty:      true,
			MobileNumbers: "7988323110",
			Address: Address{
				HouseNo:  56,
				City:     "Meham",
				District: "Rohtak",
				State:    "Haryana",
			},
		},
	},
}

func main() {
	data, err := json.MarshalIndent(allData, "  ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}
