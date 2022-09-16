package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	employees := []User{
		{"Seb", "23", "17781234567", "VTS", Address{"Waterloo", "Ontario", "Canada", "234234"}},
		{"Charles", "23", "17781234567", "PW", Address{"Vancouver", "British Columbia", "Canada", "234234"}},
		{"Supreme Leader", "23", "17781234567", "Stay home", Address{"Vancouver", "British Columbia", "Canada", "234234"}},
		{"Buc", "23", "1778621234567", "Watching anime", Address{"Vancouver", "British Columbia", "Canada", "234234"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}

	fmt.Println((allusers))

}
