package goserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
var httpClient = &http.Client{
	Transport: tr,
	Timeout:   time.Duration(10 * time.Second),
}

type errStruct struct {
	Err    string `json:"error"`
	Reason string
}

func (e errStruct) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%s: %s", e.Err, e.Reason)
}

//Client used as main client for service-now
type Client struct {
	Username, Password, Instance string
}

//PULL Pulls data from SN and returns an array of maps for each item/ticket
func (c *Client) PULL(table, opts string) ([]map[string]string, int) {

	bufDT := &bytes.Buffer{}
	bufDF := &bytes.Buffer{}
	displayTrue := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts +
		"&displayvariables=true" +
		"&displayvalue=true" +
		"&sysparm_record_count=true"
	displayFalse := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_query=" + opts +
		"&displayvariables=true" +
		"&displayvalue=false" +
		"&sysparm_record_count=true"

	reqDT, err := http.NewRequest(http.MethodGet, displayTrue, bufDT)
	checkErr(err)
	reqDF, err := http.NewRequest(http.MethodGet, displayFalse, bufDF)
	checkErr(err)

	reqDT.SetBasicAuth(c.Username, c.Password)
	reqDF.SetBasicAuth(c.Username, c.Password)

	resDT, err := httpClient.Do(reqDT)
	checkErr(err)
	resDF, err := httpClient.Do(reqDF)
	checkErr(err)

	bufDT.Reset()
	bufDF.Reset()

	var echeck errStruct

	err = json.NewDecoder(io.TeeReader(resDT.Body, bufDT)).Decode(&echeck)
	checkErr(err)
	err = json.NewDecoder(io.TeeReader(resDF.Body, bufDF)).Decode(&echeck)
	checkErr(err)

	//fmt.Printf(buf.String())

	var dT, dF struct {
		Records []map[string]string
	}
	json.NewDecoder(bufDT).Decode(&dT)
	json.NewDecoder(bufDF).Decode(&dF)

	//fmt.Println(reflect.TypeOf(v.Records))
	//for k, vv := range v.Records {
	//	fmt.Printf("\n\n-- %d --\n%s", k, vv)
	//}
	for i, eachDT := range dT.Records {
		for k, v := range eachDT {
			if strings.Contains(k, "-ID") {
				continue
			}
			if dF.Records[i][k] != v {
				dT.Records[i][k+"-ID"] = dF.Records[i][k]
			}
		}
	}

	return dT.Records, len(dT.Records)
}

//PUSH Pushes data to SN
func (c *Client) PUSH(table, opts string, body map[string]string) {
	//for _, body := range bodies {
	url := "https://" + c.Instance + table + ".do?JSON&sysparm_query=" + opts + "&sysparm_action=update"

	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(body)
	checkErr(err)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(c.Username, c.Password)

	_, err = httpClient.Do(req)
	checkErr(err)
	//}
}

/*
body[k] = v

*/

//PrintOptions for printing out options
func PrintOptions(bodies []map[string]string) {
	keys := make([]string, 0, len(bodies[0]))
	for k := range bodies[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, each := range bodies {
		fmt.Printf("\n-------------------------------------- -- --------------------------------------")
		for _, k := range keys {
			fmt.Printf("\n%38s -- %-38s", k, each[k])
		}
		fmt.Printf("\n")
	}
}

//DATAEXTRACTOR for extracting all data from a ticket / table without necessarily knowing the json formatting for the item
func (c *Client) DATAEXTRACTOR(table, opts string) string {
	buf := &bytes.Buffer{}
	testurl := "https://" + c.Instance + table + ".do?JSON&sysparm_action=getRecords&sysparm_record_count=300&sysparm_query=" + opts + "&displayvariables=true&displayvalue=true&sysparm_fields=comments"
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
