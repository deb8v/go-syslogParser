package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postSend(mode, jsonStr, url string) {

	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))

	req.Header.Set("PASSWORD", "0nm5b1ju")
	req.Header.Set("MODE", mode)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
func main() {
	/*
		var SyslogMeta struct {
			raw          string `json:"raw"`
			priority     int8   `json:"pri"`
			priorityName string `json:"priName"`
			subject      int8   `json:"subj"`
			topic        string `json:"topics"`
			date         string `json:"date"`
			timestamp    int64  `json:"timestamp"`
			timeUTC      string `json:"timeutc"`
			timeNow      int64  `json:"timenow"`
			msg          string `json:"msg"`
			typeOf       int8   `json:"typeof"` //1=RFC5424 2=RFC3164 0=RAW
			typeOfName   string `json:"typeOfName"`
		}
	*/

	json := `{"Raw":"dns query from 192.168.88.238: #132076459 bt.rutor.org. A","Priority":0,"PriorityName":"RAW","Subject":0,"Topic":"from, 192.168.88.238:","Date":"","Timestamp":1604296462,"TimeUTC":"11-02-2020 12:54:22.841775","TimeNow":1604296462,"Msg":" #132076459 bt.rutor.org. A","TypeOf":0,"TypeOfName":""}`
	postSend("PUT_NEW_ROW", json, "http://weblog.rt.khai.pw/get.php")
}
