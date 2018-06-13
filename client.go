package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    10 * time.Second,
	DisableCompression: true,
}
var HTTPClient = &http.Client{
	Transport: tr,
	Timeout:   time.Duration(10 * time.Second),
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

func (c *Client) FilterRequests(table, opts string) []Request {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)

	var v struct {
		Records []Request
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records

}

func (c *Client) FilterAssets(table, opts string) []Asset {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&displayvalue=true&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)

	var v struct {
		Records []Asset
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

func (c *Client) FilterComputers(table, opts string) []Computer {
	buf := &bytes.Buffer{}
	if table == "" {
		table = "cmdb_ci_computer"
	}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
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
	buf := &bytes.Buffer{}
	if table == "" {
		table = "cmdb_ci_computer"
	}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
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

	//err := json.NewEncoder(buf).Encode(body)
	//CheckErr(err)
	url := "https://" + c.Instance + table + ".do?JSON&sysparm_query=" + opts + "&sysparm_action=update"
	//fmt.Printf(testurl + "\n")

	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(body)
	CheckErr(err)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	CheckErr(err)

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	fmt.Println(res.Body)
}
