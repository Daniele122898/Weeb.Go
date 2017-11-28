package net

import (
	"net/http"
	"bytes"
	"io"
	"github.com/Daniele122898/weeb.go/src/helpers"
	"fmt"
	"github.com/Daniele122898/weeb.go/src/data"
	"github.com/Daniele122898/weeb.go/src/endpoints"
	"encoding/json"
)

const (
	BASE_URL = "https://api.weeb.sh/"
	DEF_CODE = 200
)

var (
	token string
)

func Authenticate(t string) error{
	token = t
	d, err := GetWelcome()
	if err != nil{
		return err
	}
	fmt.Println("Connected to weeb.sh v", d.Version)
	return nil
}

func GetWelcome() (*data.WelcomeData, error) {
	res, err := Request(endpoints.IMAGES, "", DEF_CODE)
	if err != nil{
		return nil, err
	}
	d := data.WelcomeData{}
	json.Unmarshal(res, &d)
	return &d, nil
}

func Request(endpoint, query string, expectedStatus int)([]byte, error){
	url := BASE_URL+endpoint
	if query != ""{
		url += query
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil{
		return nil, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	//close the body at the end
	defer resp.Body.Close()

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, resp.Body)
	if err != nil{
		return nil, err
	}

	if resp.StatusCode != expectedStatus{
		return nil, &helpers.UnexpectedStatus{Msg: fmt.Sprintf("Expected status %d; Got %d", expectedStatus, resp.StatusCode)}
	}

	return buf.Bytes(), nil
}