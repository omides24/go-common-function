package main

import (
	"encoding/json"
	"fmt"
)

func convertStructToMapArray(myStruct interface{}) (interface{}, error) {

	// Convert struct to json
	b, err := json.Marshal(myStruct)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(b))

	// Convert json []byte to interface{}
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Error Unmarshal", err)
		return nil, err
	}

	return f, nil
}
