package utils

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Country string `json:"country"`
}

type Customer struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobby   []string `json:"hobby"`
	Address Address
}

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data) //encode
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
