package main

import "fmt"

type Contact struct {
	MobileNo string
	Emails   []string
	Address  Address
}

type Address struct {
	HNo     int
	City    string
	State   string
	PinCode int
}

type detail struct {
	Name    string
	Contact Contact
}

func main() {
	var infos = []detail{
		{
			Name: "Test",
			Contact: Contact{
				MobileNo: "7988323110",
				Emails:   []string{"67567", "sfsdfasdf", "sfsdsf"},
				Address: Address{
					HNo:     12,
					City:    "Meham",
					State:   "Haryana",
					PinCode: 124514,
				},
			},
		},
		{
			Name: "Test 2",
			Contact: Contact{
				MobileNo: "xxxxxxx",
				Address: Address{
					HNo:     12,
					City:    "Rohtak",
					State:   "Haryana",
					PinCode: 124514,
				},
			},
		},

		{
			Name: "Test 3",
			Contact: Contact{
				MobileNo: "234224",
				Address: Address{
					HNo:     12,
					City:    "Gohana",
					State:   "Haryana",
					PinCode: 124514,
				},
			},
		},
	}

	for i, v := range infos {
		fmt.Println("my index position is  : ", i)
		if v.Contact.Address.City == "Rohtak" {
			continue
		} else {
			fmt.Println(v.Contact.Address.City)
		}

	}
}
