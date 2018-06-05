package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: false,
	DisableKeepAlives:  true,
}
var httpClient = &http.Client{Transport: tr}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Err struct {
	Err    string `json:"error"`
	Reason string
}

func (e Err) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%s: %s", e.Err, e.Reason)
}

type Client struct {
	Username, Password, Instance string
}

func (c *Client) FilterRequests(table, opts string) ([]Request, error) {

	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	req.Close = true
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)

	fmt.Printf("requ")
	CheckErr(err)

	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	CheckErr(err)

	var v struct {
		Records []Request
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil

}

func (c *Client) FilterAssets(table, opts string) []Asset {

	fmt.Printf("asse")
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	req.Close = true
	defer req.Body.Close()
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	res.Body.Close()
	CheckErr(err)

	var v struct {
		Records []Asset
	}
	json.NewDecoder(buf).Decode(&v)

	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	return v.Records
}

func (c *Client) FilterComputers(table, opts string) []Computer {

	fmt.Printf("comp")
	buf := &bytes.Buffer{}
	if table == "" {
		table = "cmdb_ci_computer"
	}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)

	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)

	var v struct {
		Records []Computer
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

func (c *Client) FilterUsers(table, opts string) []User {

	fmt.Printf("user")
	buf := &bytes.Buffer{}
	if table == "" {
		table = "cmdb_ci_computer"
	}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	CheckErr(err)

	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)

	var v struct {
		Records []User
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

func (c *Client) Update(table, opts string, body interface{}) {
	fmt.Printf(".")

	//err := json.NewEncoder(buf).Encode(body)
	//CheckErr(err)
	url := "https://" + c.Instance + table + ".do?JSON&sysparm_query=" + opts + "&sysparm_action=update"
	//fmt.Printf(testurl + "\n")

	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(body)
	CheckErr(err)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	defer func() {
		req.Body.Close()
		io.Copy(ioutil.Discard, req.Body)
	}()
	CheckErr(err)

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	CheckErr(err)
	fmt.Println(res.Body)
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	return
}
