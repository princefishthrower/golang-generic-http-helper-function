package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http_helper/http_helper"
	"net/url"
)

func main() {

	/////////////////
	// GET REQUEST //
	/////////////////

	// the url to call
	fullUrl := "https://reqres.in/api/users/2"

	// the headers to pass - none required for this test api
	headers := map[string]string{}

	// the query parameters to pass - none required but example commented out here
	queryParameters := url.Values{}
	// queryParameters.Add("per_page", "1")

	// the type to unmarshal the response into
	type UserResponse struct {
		Data struct {
			ID        int    `json:"id"`
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Avatar    string `json:"avatar"`
		} `json:"data"`
		Support struct {
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"support"`
	}

	var response UserResponse

	// call the function
	response, err := http_helper.MakeHTTPRequest(fullUrl, "GET", headers, queryParameters, nil, response)
	if err != nil {
		panic(err)
	}

	// do something with the response
	fmt.Printf("response: %+v", response)

	//////////////////
	// POST REQUEST //
	//////////////////

	// the url to call
	fullUrl = "https://reqres.in/api/users"

	// the headers to pass
	headers = map[string]string{}

	// the query parameters to pass
	// queryParameters = url.Values{}
	// queryParameters.Add("per_page", "1")

	// the body to pass - a string like this is possible...
	// body := bytes.NewBufferString(`{"name": "luigi","job": "Goomba stomper"}`)
	// but much better is to use a struct and marshal it to json:
	type User struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}
	user := User{Name: "luigi", Job: "Goomba stomper"}
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// call the function
	response, err = http_helper.MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, bytes.NewBuffer(body), response)
	if err != nil {
		panic(err)
	}

	// do something with the response
	fmt.Printf("response: %+v", response)

	//////////////////
	// PUT REQUEST //
	//////////////////

	// the url to call
	fullUrl = "https://reqres.in/api/users/2"

	// the headers to pass
	headers = map[string]string{}

	// the query parameters to pass
	// queryParameters = url.Values{}
	// queryParameters.Add("per_page", "1")

	// the body to pass - let's update luigi's job
	user = User{Name: "luigi", Job: "Amazing Goomba stomper!!!!"}
	body, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	type UserPutResponse struct {
		Name      string `json:"name"`
		Job       string `json:"job"`
		UpdatedAt string `json:"updatedAt"`
	}

	// call the function
	response, err = http_helper.MakeHTTPRequest(fullUrl, "POST", headers, queryParameters, bytes.NewBuffer(body), response)
	if err != nil {
		panic(err)
	}

	// do something with the response
	fmt.Printf("response: %+v", response)

}
