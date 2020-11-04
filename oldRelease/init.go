package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var queue []string

func postSend(mode string, json string, url string) {

	fmt.Println("URL:>", url)

	req, rqErr := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json)))

	req.Header.Set("PASSWORD", "0nm5b1ju")
	req.Header.Set("MODE", mode)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, doErr := client.Do(req)
	if rqErr != nil {
		panic(rqErr)
	}
	if doErr != nil {
		panic(doErr)
	}
	defer resp.Body.Close()
	fmt.Println("json:", json)
	fmt.Println("response Status:", mode)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
func getTopicsfromRAW(raw string) string {
	c := strings.SplitAfterN(raw, ":", 2)
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "\t", 2)
	}
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "on", 2)
	}
	if len(c) < 2 {
		c = strings.SplitAfterN(raw, "     ", 2)
	}
	e := strings.TrimSpace(c[0])
	e = strings.ReplaceAll(e, " ", ",")
	a := strings.Split(e, ",")
	topics := ""
	for index, element := range a {
		if index < len(a)-1 {
			topics += element + ", "
		} else {
			topics += element
		}

	}

	return topics
}
func getTopicsfromRFC3164(raw string) string {
	e := strings.Split(raw, " ")

	topics := ""
	if len(e) > 6 {
		topics = e[4] + ", " + e[5]
	}
	return topics
}
func getTopicsfromRFC5424(raw string) string {
	e := strings.Split(raw, " ")

	topics := ""
	if len(e) > 4 {
		topics = e[2] + ", " + e[3]
	}
	return topics
}
func HMS2UT(s string) int64 {

	t2, err := time.Parse("Jan 02 15:04:05 2006", "Jan 01 "+s+" 1970")
	if err != nil {
		panic(err)
	}
	nix := t2.Unix()
	//fmt.Printf(">time:  %d\n", nix)
	timestamp := nix

	return timestamp
}
func UTC2UT(s string) int64 {

	var timestamp int64 = 0

	//1985-04-12T19:20:50.52-04:00
	t2, err := time.Parse("2006-01-02T15:04:05.999999-07:00", s)
	if err != nil {
		panic(err)
	}
	nix := t2.Unix()
	//fmt.Printf(">time:  %d\n", nix)
	timestamp = nix

	return timestamp
}
func getMsgfromRAW(raw string) string {
	e := strings.SplitAfterN(raw, ":", 2)
	msg := ""
	if len(e) < 2 {
		e = strings.SplitAfterN(raw, "\t", 2)
	} else {
		if len(e) < 2 {
			e = strings.SplitAfterN(raw, "on", 2)
		} else {
			if len(e) < 2 {
				e = strings.SplitAfterN(raw, "     ", 2)
			}
		}
	}
	if len(e) < 2 {
		msg = e[0]
	} else {
		msg = e[1]
	}
	return msg
}
func getMsgfromRFC5424(raw string) string {
	e := strings.Split(raw, "- ")
	g := e[2]
	export := ""
	if strings.Count(g, "]") == 1 {
		export = strings.Split(g, "]")[1]
	}
	if strings.Count(g, "]") == 0 {
		export = strings.Split(raw, "- ")[2]
	}

	return strings.TrimSpace(export)

}
func getMsgfromRFC3164(raw string) string {
	e := strings.Split(raw, ":")
	if len(e) == 3 {
		return strings.TrimSpace(e[2])
	} else {
		return strings.TrimSpace(e[3])
	}
}
func getCritIndexByString(s string) int8 {
	i := strings.ToLower(s)
	ret := int8(127)
	/*
		var v = map[string]int8{
			"emergency":     0,
			"alert":         1,
			"critical":      2,
			"error":         3,
			"warning":       4,
			"notice":        5,
			"informational": 6,
			"debug":         7,
		}
	*/
	if strings.Contains(i, "debug") {
		ret = 7
	}
	if strings.Contains(i, "info") {
		ret = 6

	}
	if strings.Contains(i, "action") {
		ret = 6

	}
	if strings.Contains(i, "notice") {
		ret = 5
	}
	if strings.Contains(i, "warning") {
		ret = 4
	}
	if strings.Contains(i, "error") {
		ret = 3
	}
	if strings.Contains(i, "critical") {
		ret = 2
	}
	if strings.Contains(i, "alert") {
		ret = 1
	}
	if strings.Contains(i, "emergency") {
		ret = 0
	}

	return ret
}
func getStatusByDict(status int8) string {
	exp := ""

	var m = map[int8]string{
		0: "Emergency",
		1: "Alert",
		2: "Critical",
		3: "Error",
		4: "Warning",
		5: "Notice",
		6: "Informational",
		7: "Debug",
	}

	if status > 0 && status <= 7 {
		//export := (dict[status])
		exp = m[status]
		//exp = export

	} else {
		exp = "Debug"
	}

	return exp
}

var RFCList = map[int8]string{
	2: "RFC3164 (BSD)",
	1: "RFC5424",
	0: "RAW",
}

func parseMessage(data string, pid uint64) {
	fmt.Println(pid, " started (thread)", data)
	var syslogMeta struct {
		Raw          string `json:"raw"`
		Priority     int8   `json:"priority"`
		PriorityName string `json:"priorityName"`
		Subject      int8   `json:"subjectData"`
		Topic        string `json:"topics"`
		Date         string `json:"date"`
		Timestamp    int64  `json:"timestamp"`
		TimeUTC      string `json:"timeutc"`
		TimeNow      int64  `json:"timenow"`
		Msg          string `json:"msg"`
		TypeOf       int8   `json:"typeOf"` //1=RFC5424 2=RFC3164 0=RAW
		TypeOfName   string `json:"typeOfName"`
	}
	syslogMeta.Raw = data

	priority := strings.Split(data, ">")[0]
	priority = priority[1:len(priority)]

	priorityInt, errorOfParsePriorityMeta := strconv.ParseInt(priority, 10, 10)

	if errorOfParsePriorityMeta != nil {
		priorityInt = -1
	}
	syslogMeta.Priority = int8(priorityInt - 8)

	checkrfc := strings.Split(data, ">")

	if len(checkrfc) == 1 {
		syslogMeta.TypeOf = 0 //raw
	} else {
		if len(checkrfc[1]) > 0 {
			subject := checkrfc[1][0:1]
			//

			if _, err := strconv.Atoi(subject); err == nil {
				//RFC5424 by DATE validator

				subjectInt, errorOfParsesubjectMeta := strconv.ParseInt(subject, 10, 2)

				syslogMeta.Subject = int8(subjectInt)
				if errorOfParsesubjectMeta != nil {
					subjectInt = 1
				}

				if subjectInt <= 23 && subjectInt >= 0 {
					syslogMeta.TypeOf = 1
				} else {
					syslogMeta.TypeOf = 2
				}
			} else {
				syslogMeta.TypeOf = 2
			}

		} else {
			syslogMeta.TypeOf = 0
		}
	}
	if syslogMeta.TypeOf == 2 {
		roof := strings.Split(data, " ")[3]
		dayNow := (time.Now().Unix() - time.Now().Unix()%86400)
		hms := HMS2UT(roof)
		syslogMeta.Timestamp = dayNow + hms
		syslogMeta.TimeUTC = roof

		syslogMeta.Topic = getTopicsfromRFC3164(syslogMeta.Raw)
		syslogMeta.Msg = getMsgfromRFC3164(syslogMeta.Raw)
	}
	if syslogMeta.TypeOf == 1 {
		roof := strings.Split(data, " ")[1]
		syslogMeta.Timestamp = UTC2UT(roof)
		//syslogMeta.TimeUTC = roof

		syslogMeta.Topic = getTopicsfromRFC5424(syslogMeta.Raw)
		syslogMeta.Msg = getMsgfromRFC5424(syslogMeta.Raw)
	}
	if syslogMeta.TypeOf == 0 {
		timenow := time.Now()
		syslogMeta.Timestamp = timenow.Unix()
		syslogMeta.TimeUTC = timenow.Format("01-02-2006 15:04:05.000000")

		syslogMeta.Topic = getTopicsfromRFC5424(syslogMeta.Raw)
		syslogMeta.Priority = getCritIndexByString(syslogMeta.Topic)
		syslogMeta.Msg = getMsgfromRAW(syslogMeta.Raw)
	}
	syslogMeta.TimeNow = time.Now().Unix()
	syslogMeta.PriorityName = getStatusByDict(syslogMeta.Priority)
	syslogMeta.TypeOfName = RFCList[syslogMeta.TypeOf]
	//fmt.Printf("typeOf:   %d\n", pexInt)
	/*
		fmt.Printf("->:  %s\n", syslogMeta.Raw)
		fmt.Printf("\\-typeOf:   \t%d\n", syslogMeta.TypeOf)
		fmt.Printf("\\-tpe.OfNme:\t%d\n", syslogMeta.TypeOf)
		fmt.Printf("\\-subject:  \t%d\n", syslogMeta.Subject)
		fmt.Printf("\\-priority: \t%d\n", syslogMeta.Priority)
		fmt.Printf("\\-pri. name:\t%s\n", syslogMeta.PriorityName)
		fmt.Printf("\\-timestamp:\t%d\n", syslogMeta.Timestamp)
		fmt.Printf("\\-timeUTC:  \t%s\n", syslogMeta.TimeUTC)
		fmt.Printf("\\-timeNow:  \t%d\n", syslogMeta.TimeNow)
		fmt.Printf("\\-topics:   \t%s\n", syslogMeta.Topic)

		fmt.Printf("\\-message:  \t%s\n", syslogMeta.Msg)

		//fmt.Println(pid, " wait (thread)")
		//time.Sleep(20 * time.Second)
	*/
	jms, err := json.Marshal(syslogMeta)
	jsonString := string(jms)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jms))
	postSend("PUT_NEW_ROW", jsonString, "http://weblog.rt.khai.pw/get.php")
	//fmt.Println(string(syslogMeta))
	fmt.Println(pid, " end (thread)")
}
func queueChannel(c chan uint64) {
	for {
		newMessage := <-c
		fmt.Println(newMessage, " ended")
	}
}

func listen(messages chan string, done chan bool) {
	for { //в вечном цикле ждём что появилось нового
		select {
		case <-done:
			return
		case msg := <-messages: //получаем сообщение в обработку повешав цикл
			go parseMessage(msg, 55) //запуск аснхронного потока
		}
	}

}

func main() {

	messages := make(chan string)
	done := make(chan bool)
	go listen(messages, done)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 514,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 4096)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		//dataFrom := strings.Count(data, ':')
		messages <- data
		print(len(queue))
		fmt.Printf("from %s\n\n", remote)
	}

}
