package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []Request
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil

}

func (c *Client) FilterAssets(table, opts string) []Asset {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := http.DefaultClient.Do(req)
	CheckErr(err)
	defer res.Body.Close()
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

	res, err := http.DefaultClient.Do(req)
	CheckErr(err)
	defer res.Body.Close()
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

	res, err := http.DefaultClient.Do(req)
	CheckErr(err)
	defer res.Body.Close()
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

	res, err := http.DefaultClient.Do(req)
	CheckErr(err)
	defer res.Body.Close()
	fmt.Println(res.Body)

}
