package main

import "fmt"

type StructExp1 struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type StructExp2 struct {
	FirstName string
	LastName  string
}

type StructExp3 struct {
	FirstName string
	LastName  string
	Colors    []string
}

func main() {

	s := StructExp3{FirstName: "omid", LastName: "esmaeili", Colors: []string{"red", "blue", "green"}}
	f, err := convertStructToMapArray(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

}
