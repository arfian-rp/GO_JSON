package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json/utils"
	"os"
	"testing"
)

func TestMarshal(t *testing.T) {
	utils.LogJson(1)
	utils.LogJson("hello world")
	utils.LogJson(true)
	utils.LogJson([]string{"hello", "world"})

	utils.LogJson(utils.Customer{
		Name:  "Hendrich",
		Age:   33,
		Hobby: []string{"run"},
		Address: utils.Address{
			Country: "Dutch",
		},
	})
}

func TestUnmarshal(t *testing.T) {
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	customer := &utils.Customer{}
	json.Unmarshal(jsonData, customer)
	fmt.Println(customer)
	fmt.Println("name:", customer.Name)
}

func TestMap(t *testing.T) {
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)

	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["age"])
	fmt.Println(result["hobby"])
	fmt.Println(result["address"])
}

func TestJsonDecoder(t *testing.T) {
	reader, _ := os.Open("data.json")
	decoder := json.NewDecoder(reader)
	customer := utils.Customer{}

	decoder.Decode(&customer)
	fmt.Println(customer)
}

func TestJsonEncoder(t *testing.T) {
	writer, _ := os.Create("sample.json")
	encoder := json.NewEncoder(writer)

	customer := utils.Customer{
		Name:  "Richard",
		Age:   17,
		Hobby: []string{"race"},
		Address: utils.Address{
			Country: "USA",
		},
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
