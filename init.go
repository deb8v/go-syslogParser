package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var queue []string

func getTopicsfromRAW(raw string) string {
	e := strings.Split(raw, ":")[0]
	e = strings.ReplaceAll(e, " ", ",")
	a := strings.Split(e, ",")
	topics := ""
	for index, element := range a {
		if index < len(a)-2 {
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
	e := strings.Split(raw, ": ")[1]

	return e
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

	return strings.TrimSpace(e[3])
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

func parseMessage(data string, pid uint64, queueChannel chan uint64) {
	fmt.Println(pid, " started (thread)")
	var syslogMeta struct {
		raw          string
		priority     int8
		priorityName string
		subject      int8
		topic        string
		date         string
		timestamp    int64
		timeUTC      string
		timeNow      int64
		msg          string
		typeOf       int8 //1=RFC5424 2=RFC3164 0=RAW
		typeOfName   string
	}
	syslogMeta.raw = data

	priority := strings.Split(data, ">")[0]
	priority = priority[1:len(priority)]

	priorityInt, errorOfParsePriorityMeta := strconv.ParseInt(priority, 10, 10)

	if errorOfParsePriorityMeta != nil {
		priorityInt = -1
	}
	syslogMeta.priority = int8(priorityInt - 8)

	checkrfc := strings.Split(data, ">")
	if len(checkrfc) == 1 {
		syslogMeta.typeOf = 0 //raw
	} else {

		subject := checkrfc[1][0:1]
		//

		if _, err := strconv.Atoi(subject); err == nil {
			//RFC5424 by DATE validator

			subjectInt, errorOfParsesubjectMeta := strconv.ParseInt(subject, 10, 2)

			syslogMeta.subject = int8(subjectInt)
			if errorOfParsesubjectMeta != nil {
				subjectInt = 1
			}

			if subjectInt <= 23 && subjectInt >= 0 {
				syslogMeta.typeOf = 1
			} else {
				syslogMeta.typeOf = 2
			}
		} else {
			syslogMeta.typeOf = 2
		}

	}
	if syslogMeta.typeOf == 2 {
		roof := strings.Split(data, " ")[3]
		dayNow := (time.Now().Unix() - time.Now().Unix()%86400)
		hms := HMS2UT(roof)
		syslogMeta.timestamp = dayNow + hms
		syslogMeta.timeUTC = roof

		syslogMeta.topic = getTopicsfromRFC3164(syslogMeta.raw)
		syslogMeta.msg = getMsgfromRFC3164(syslogMeta.raw)
	}
	if syslogMeta.typeOf == 1 {
		roof := strings.Split(data, " ")[1]
		syslogMeta.timestamp = UTC2UT(roof)
		//syslogMeta.timeUTC = roof

		syslogMeta.topic = getTopicsfromRFC5424(syslogMeta.raw)
		syslogMeta.msg = getMsgfromRFC5424(syslogMeta.raw)
	}
	if syslogMeta.typeOf == 0 {
		timenow := time.Now()
		syslogMeta.timestamp = timenow.Unix()
		syslogMeta.timeUTC = timenow.Format("01-02-2006 15:04:05.000000")
		syslogMeta.priority = 0

		syslogMeta.topic = getTopicsfromRFC5424(syslogMeta.raw)
		syslogMeta.msg = getMsgfromRAW(syslogMeta.raw)
	}
	syslogMeta.timeNow = time.Now().Unix()
	syslogMeta.priorityName = getStatusByDict(syslogMeta.priority)
	syslogMeta.priorityName = RFCList[syslogMeta.typeOf]
	//fmt.Printf("typeOf:   %d\n", pexInt)
	/*
		fmt.Printf("->:  %s\n", syslogMeta.raw)
		fmt.Printf("\\-typeOf:   \t%d\n", syslogMeta.typeOf)
		fmt.Printf("\\-tpe.OfNme:\t%d\n", syslogMeta.typeOf)
		fmt.Printf("\\-subject:  \t%d\n", syslogMeta.subject)
		fmt.Printf("\\-priority: \t%d\n", syslogMeta.priority)
		fmt.Printf("\\-pri. name:\t%s\n", syslogMeta.priorityName)
		fmt.Printf("\\-timestamp:\t%d\n", syslogMeta.timestamp)
		fmt.Printf("\\-timeUTC:  \t%s\n", syslogMeta.timeUTC)
		fmt.Printf("\\-timeNow:  \t%d\n", syslogMeta.timeNow)
		fmt.Printf("\\-topics:   \t%s\n", syslogMeta.topic)
		fmt.Printf("\\-message:  \t%s\n", syslogMeta.msg)
	*/

	queueChannel <- pid
	fmt.Println(pid, " ended (thread)")
}
func queueChannel(c chan uint64) {
	for {
		newMessage := <-c
		fmt.Println(newMessage, " ended")
	}
}
func queueManagaer(c chan uint64) {

	var pid uint64 = 0
	for len(queue) > 0 {
		fmt.Print(len(queue)) // Первый элемент
		go parseMessage(queue[0], pid, c)
		queue[0] = ""
		queue = queue[1:] // Удаление из очереди
		pid += 1
	}
	fmt.Print("queDeath") // Первый элемент
}
func main() {
	queueChannel := make(chan uint64)
	go queueManagaer(queueChannel)

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
		queue = append(queue, data)
		print(len(queue))
		fmt.Printf("from %s\n\n", remote)
	}

}
