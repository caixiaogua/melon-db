package melondb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read error.")
	}
	src := string(body)
	return src
}

func httpPost(url string, data string, ctype string) string {
	if ctype == "" {
		ctype = "application/x-www-form-urlencoded"
	} else if ctype == "json" {
		ctype = "application/json"
	}
	// fmt.Println("data", data)
	res, err := http.Post(url, ctype, strings.NewReader(data))
	if err != nil {
		return err.Error()
	}
	// fmt.Println("res", res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func Init(db string, url string) func(string) string {
	return func(str string) string {
		jsonObj := map[string]string{
			"t": db,
			"s": str,
		}
		data, err := json.Marshal(jsonObj)
		if err != nil {
			return "jsonParseError"
		}
		return httpPost(url, string(data), "")
	}
}
