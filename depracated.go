package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

//FilterRequests used for filtering through a request type table
func (c *Client) FilterRequests(table, opts string) []Request {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)

	var v struct {
		Records []Request
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records

}

//FilterAssets used for filtering through a asset type table
func (c *Client) FilterAssets(table, opts string) []Asset {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)

	var v struct {
		Records []Asset
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

//FilterComputers used for filtering through a computer type table
func (c *Client) FilterComputers(table, opts string) []Computer {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)

	var v struct {
		Records []Computer
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

//FilterIncidents used for filtering through a incident type table
func (c *Client) FilterIncidents(table, opts string) []Incident {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)

	var v struct {
		Records []Incident
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records
}

//FilterUsers used for filtering through a user type table
func (c *Client) FilterUsers(table, opts string) ([]User, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

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

//FilterSCTasks used for filtering through a sctask type table
func (c *Client) FilterSCTasks(table, opts string) ([]SCTask, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []SCTask
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

//FilterHistory used for filtering through a history type table
func (c *Client) FilterHistory(table, opts string) ([]History, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

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

//FilterSysAudit used for filtering through a sys_audit type table
func (c *Client) FilterSysAudit(table, opts string) ([]SysAudit, error) {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	if err != nil {
		return nil, err
	}

	var v struct {
		Records []SysAudit
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records, nil
}

//FilterInterface is:
// THIS RETURNS A MAP OF EACH ITEM, almost like a struct was built for it. THIS ONLY WORKS WITH ONE TICKET/ITEM return at a time.
// If more than one is returned, the last element is used.
func (c *Client) FilterInterface(table, opts string) map[string]string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=1&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)
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

//Update pushes a body item to the filtered string. THIS CAN UPDATE AN INFINITE NUMBER OF ITEMS! PLEASE VERIFY YOUR STRING FIRST!
func (c *Client) Update(table, opts string, body interface{}) {

	//err := json.NewEncoder(buf).Encode(body)
	//CheckErr(err)
	url := "https://" + c.Instance + table + ".do?JSON&sysparm_query=" + opts + "&sysparm_action=update"
	//fmt.Printf(testurl + "\n")

	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(body)
	checkErr(err)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	fmt.Println(res.Body)
}

//FilterRequestsLITERAL returns a table of the ID information, not a username / information
func (c *Client) FilterRequestsLITERAL(table, opts string) []Request {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts + "&displayvariables=true&sysparm_record_count=300"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)

	var v struct {
		Records []Request
	}
	json.NewDecoder(buf).Decode(&v)

	return v.Records

}

//JSONBUILDER for building JSON objects to expand this program
func (c *Client) JSONBUILDER(table, opts string) string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=300&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true"
	req, err := http.NewRequest(http.MethodGet, testurl, buf)
	checkErr(err)

	req.SetBasicAuth(c.Username, c.Password)

	res, err := httpClient.Do(req)
	checkErr(err)
	buf.Reset()
	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck)
	checkErr(err)
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
