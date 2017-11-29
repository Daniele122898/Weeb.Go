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
	"strings"
)

const (
	BASE_URL = "https://api.weeb.sh/"
	DEF_CODE = 200
	wrapperVersion = "1.0.0"
)

type FileType int
type Nsfw int

const(
	JPG FileType = iota
	PNG
	GIF
	ANY
)

const(
	FALSE Nsfw = iota
	TRUE
	ONLY
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
	fmt.Println("Successfully connected to weeb.sh v", d.Version, "using the Weeb.Go v", wrapperVersion, "wrapper")
	return nil
}

func getHidden(hidden bool) string{
	switch hidden {
		case true:
			return "true"
		case false:
			return "false"
	}
	return "false"
}

func getFiletype(ft FileType) string {
	switch ft {
	case GIF:
		return "gif"
	case JPG:
		return "jpg"
	case PNG:
		return "png"
	default:
		return ""

	}
}

func getNsfw(nswf Nsfw) string{
	switch nswf {
	case TRUE:
		return "true"
	case ONLY:
		return "only"
	default:
		//false
		return "false"
	}
}

func GetTypes(hidden bool) (*data.TypesData, error){
	res, err := Request(endpoints.TYPES, "?hidden="+getHidden(hidden), DEF_CODE)
	if err != nil{
		return nil, err
	}

	td:= data.TypesData{}
	err = json.Unmarshal(res, &td)
	if err !=nil{
		return nil, err
	}
	return &td, nil
}

func GetTags(hidden bool) (*data.TagsData, error){
	res, err := Request(endpoints.TAGS,"?hidden="+getHidden(hidden), DEF_CODE)
	if err != nil {
		return nil, err
	}

	td := data.TagsData{}
	err = json.Unmarshal(res, &td)
	if err!= nil{
		return nil, err
	}
	return &td, nil
}

func GetRandom(typ string, tags []string,filetype FileType,nsfw Nsfw, hidden bool) (*data.RandomData, error) {



	query :=""
	if typ != ""{
		query+="&type="+typ
	}
	if tags != nil && len(tags) != 0{
		t:= strings.Join(tags, ",")
		if t != "" {
			query+= "&tags="+t
		}
	}
	query +="&hidden="+getHidden(hidden)+"&nsfw="+getNsfw(nsfw)
	if filetype != ANY {
		query += "&filetype="+getFiletype(filetype)
	}
	query = strings.TrimPrefix(query, "&")
	query = "?"+query
	res, err := Request(endpoints.RANDOM, query, DEF_CODE)
	if err != nil{
		return nil, err
	}
	d := data.RandomData{}
	err = json.Unmarshal(res, &d)
	if err != nil{
		return nil, err
	}
	return &d, nil
}

func GetWelcome() (*data.WelcomeData, error) {
	res, err := Request(endpoints.IMAGES, "", DEF_CODE)
	if err != nil{
		return nil, err
	}
	d := data.WelcomeData{}
	err = json.Unmarshal(res, &d)
	if err != nil{
		return nil, err
	}
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
		return nil, &helpers.UnexpectedStatus{Msg: fmt.Sprintf("Expected status %d; Got %d; MSG: %s", expectedStatus, resp.StatusCode, resp.Status)}
	}

	return buf.Bytes(), nil
}