package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
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
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
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

func (c *Client) JSONBUILDER(table, opts string) string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=300&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)
	item := buf.String()
	//fmt.Printf(temp + "\n\n")

	item = strings.TrimPrefix(item, "{\"records\":[{")
	item = strings.TrimSuffix(item, "}]}")

	item = strings.Split(item, "},{")[0]

	var FirstCut = regexp.MustCompile(`:"([^,]*)("{1})`)
	item = string(FirstCut.ReplaceAll([]byte(item), []byte("")))

	var SecondCut = regexp.MustCompile(`:".*?(")`)
	item = string(SecondCut.ReplaceAll([]byte(item), []byte("")))

	var ThirdCut = regexp.MustCompile(`","`)
	item = string(ThirdCut.ReplaceAll([]byte(item), []byte("\n")))

	var FourthCut = regexp.MustCompile(`"[^\n]*\n`)
	item = string(FourthCut.ReplaceAll([]byte(item), []byte("")))

	var FifthCut = regexp.MustCompile(`"`)
	item = string(FifthCut.ReplaceAll([]byte(item), []byte("")))

	var Variables = regexp.MustCompile(`(variables).*`)
	item = string(Variables.ReplaceAllStringFunc(item, func(m string) string {
		return m[:9] + "\n" + m[9:]
	}))

	jsonList := strings.Split(item, "\n")

	var SingleWord = regexp.MustCompile(`_(.{1})`)
	item = SingleWord.ReplaceAllStringFunc(item, func(m string) string {
		return strings.ToUpper(m)[1:]
	})

	var FrontUp = regexp.MustCompile(`\n.`)
	item = FrontUp.ReplaceAllStringFunc(item, func(m string) string {
		return strings.ToUpper(m)
	})

	regList := strings.Split(item, "\n")
	//fmt.Printf(item)

	for i, each := range jsonList {
		fmt.Printf("%-40s string          `json:\"%s,omitempty\"`\n", regList[i], each)
	}
	return " "

}

func (c *Client) DATAEXTRACTOR(table, opts string) string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=300&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_fields=comments"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)
	item := buf.String()

	item = strings.TrimPrefix(item, "{\"records\":[{")
	item = strings.TrimSuffix(item, "}]}")

	item = strings.Split(item, "},{")[0]

	var ThirdCut = regexp.MustCompile(`","`)
	item = string(ThirdCut.ReplaceAll([]byte(item), []byte("\n")))

	var FourthCut = regexp.MustCompile(`":`)
	item = string(FourthCut.ReplaceAll([]byte(item), []byte("|")))

	var FifthCut = regexp.MustCompile(`"`)
	item = string(FifthCut.ReplaceAll([]byte(item), []byte("")))

	regList := strings.Split(item, "\n")
	//fmt.Printf(item)
	for _, each := range regList {
		temp := strings.Split(each, "|")
		if len(temp) > 1 {
			fmt.Printf("%40s -- %s\n", temp[0], temp[1])
		} else {
			fmt.Printf("%40s --\n", temp[0])
		}
	}
	return ""

}

func (c *Client) FilterAssets(table, opts string) []Asset {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
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
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
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

func (c *Client) FilterIncidents(table, opts string) []Incident {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
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
		Records []Incident
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

func (c *Client) FilterUsers(table, opts string) ([]User, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []User
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

func (c *Client) FilterSCTasks(table, opts string) ([]sctask, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []sctask
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

func (c *Client) FilterHistory(table, opts string) ([]History, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []History
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

func (c *Client) FilterSysAudit(table, opts string) ([]sysAudit, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []sysAudit
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

//
//
// THIS RETURNS A MAP OF EACH ITEM, almost like a struct was built for it. THIS ONLY WORKS WITH ONE TICKET/ITEM return at a time.
// If more than one is returned, the last element is used.
func (c *Client) FilterInterface(table, opts string) map[string]string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=1&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	CheckErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := HTTPClient.Do(req)
	CheckErr(err)
	buf.Reset()
	var echeck Err

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	CheckErr(err)
	item := buf.String()
	//fmt.Printf(temp + "\n\n")

	item = strings.TrimPrefix(item, "{\"records\":[{")
	item = strings.TrimSuffix(item, "}]}")
	item = strings.Replace(item, "[", "\"", -1)
	item = strings.Replace(item, "]", "\"", -1)
	//fmt.Println(item)
	//fmt.Printf("\n\n")
	datas := make(map[string]string)
	itemArray := strings.Split(item, "\",\"")

	for _, elem := range itemArray {
		elem = strings.Replace(elem, "\"", "", -1)
		elemArray := strings.SplitN(elem, ":", 2)
		datas[elemArray[0]] = elemArray[1]
		//fmt.Printf("%20s = %s\n", elemArray[0], elemArray[1])
	}
	return datas
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
